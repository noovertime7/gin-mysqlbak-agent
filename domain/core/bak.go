package core

import (
	"backupAgent/domain/config"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/alioss"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/ding"
	"backupAgent/domain/pkg/dingproxy"
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/pkg/minio"
	"context"
	"errors"
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/robfig/cron/v3"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	TaskID      int64
	Cron        *cron.Cron
	Engine      *xorm.Engine
	Host        string
	BackInfoId  string
	PassWord    string
	User        string
	Port        string
	DbName      string
	BackupCycle string
	KeepNumber  int64
	ISAllDBBak  int64
	DingConfig  *dao.DingDatabase
	OssConfig   *dao.OssDatabase
	DingStatus  int64
	OssStatus   int64
	BakStatus   int64
	BakMsg      string
	FileName    string
	FileSize    int64
}

var CronJob = make(map[int64]*cron.Cron)

func NewBakHandler(detail *dao.TaskDetail) (*Handler, error) {
	handler := &Handler{
		TaskID:      detail.Info.Id,
		Host:        detail.Host.Host,
		PassWord:    detail.Host.Password,
		User:        detail.Host.User,
		DbName:      detail.Info.DBName,
		BackupCycle: detail.Info.BackupCycle,
		KeepNumber:  detail.Info.KeepNumber,
		ISAllDBBak:  detail.Info.IsAllDBBak,
		DingConfig:  detail.Ding,
		OssConfig:   detail.Oss,
	}
	en, err := xorm.NewEngine("mysql", handler.User+":"+handler.PassWord+"@tcp("+handler.Host+")/"+handler.DbName+"?charset=utf8&parseTime=true")
	if err != nil {
		return nil, err
	}
	handler.Engine = en
	c := cron.New()
	if _, ok := CronJob[handler.TaskID]; ok {
		return nil, errors.New("当前备份任务已启动，切勿重复启动")
	}
	CronJob[handler.TaskID] = c
	handler.Cron = c
	return handler, nil
}

func (b *Handler) BeforBak() {
	path, _ := os.Getwd()
	pkg.CreateDir(path + "/bakfile/")
	dir := path + "/bakfile/" + strconv.Itoa(int(b.TaskID))
	baktime := time.Now().Format("2006-01-02-15-04")
	pkg.CreateDir(dir)
	host := strings.Split(b.Host, ":")[0]
	file := dir + "/" + host + "_" + b.DbName + "_" + baktime + ".sql"
	b.FileName = file
}

func (b *Handler) StartBak() error {
	_, err := b.Cron.AddJob(b.BackupCycle, b)
	if err != nil {
		return err
	}
	log.Logger.Infof("创建备份任务成功，任务id：%d", b.TaskID)
	// 备份前准备工作
	b.BeforBak()
	//启动数据库备份服务
	b.Cron.Start()
	return nil
}

func (b *Handler) IsStart(tid int64) bool {
	if _, ok := CronJob[tid]; !ok {
		return false
	}
	return true
}

func (b *Handler) StopBak(tid int64) error {
	if ok := b.IsStart(tid); !ok {
		log.Logger.Warningf("任务ID不存在%d,当前运行队列%v", tid, CronJob)
		return errors.New("任务非运行,停止失败")
	}
	log.Logger.Debug("StopBak", CronJob)
	for id, cronjob := range CronJob {
		if id == tid {
			cronjob.Stop()
			delete(CronJob, tid)
			log.Logger.Infof("任务ID:%d,备份库名:%s 停止成功", id, b.DbName)
			return nil
		}
	}
	return nil
}

func (b *Handler) Run() {
	log.Logger.Info("BakHandler 开始备份数据库")
	b.BeforBak()
	err := b.Engine.DumpAllToFile(b.FileName)
	if err != nil {
		if err := b.RunMysqlDump(); err != nil {
			b.BakStatus = 0
			b.OssStatus = 2
			b.DingStatus = 2
			b.BakMsg = fmt.Sprintf("%s", err)
			b.FileName = "unknown"
			AfterBak(b)
			log.Logger.Error("备份失败,保存备份历史到数据库,停止备份任务,发送消息", err)
			if err := b.StoreDatabase(); err != nil {
				log.Logger.Error("数据库存储失败", err)
				return
			}
			log.Logger.Debug("数据库存储成功")
			return
		}
		b.BakMsg = "success"
		b.BakStatus = 1
		b.FileSize = int64(pkg.GetFileSize(b.FileName))
		//首先判定钉钉 oss都操作成功，状态改为1
		b.DingStatus = 2
		b.OssStatus = 2
		//判断是否启动钉钉提醒
		AfterBak(b)
		//发送对象到channel
		log.Logger.Info("备份数据库成功,保存备份历史到数据库,发送消息")
		if err := b.StoreDatabase(); err != nil {
			log.Logger.Error("数据库存储失败", err)
			return
		}
		log.Logger.Debug("数据库存储成功")
		return
	}
	b.BakMsg = "success"
	b.BakStatus = 1
	b.FileSize = int64(pkg.GetFileSize(b.FileName))
	//首先判定钉钉 oss都操作成功，状态改为1
	b.DingStatus = 2
	b.OssStatus = 2
	//判断是否启动钉钉提醒
	AfterBak(b)
	//发送对象到channel
	log.Logger.Info("备份数据库成功,保存备份历史到数据库,发送消息")
	if err := b.StoreDatabase(); err != nil {
		log.Logger.Error("数据库存储失败", err)
		return
	}
	log.Logger.Debug("数据库存储成功")
}

func (b *Handler) RunMysqlDump() error {
	log.Logger.Warning("使用xorm备份失败，尝试使用mysqldump进行备份")
	iphost, port := strings.Split(b.Host, ":")[0], strings.Split(b.Host, ":")[1]
	command := fmt.Sprintf("mysqldump -u%v -p%v -P%v -h%v  %v >  %v", b.User, b.PassWord, port, iphost, b.DbName, b.FileName)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		log.Logger.Error("mysqldump执行失败:", command, " with error: ", err.Error())
		return err
	}
	log.Logger.Infof("mysqldump执行成功:%v:%v", b.Host, b.DbName)
	return nil
}

func (b *Handler) StoreDatabase() error {
	historyDB := &dao.BakHistory{
		TaskID:     b.TaskID,
		Host:       b.Host,
		DBName:     b.DbName,
		OssStatus:  b.OssStatus,
		DingStatus: b.DingStatus,
		BakStatus:  b.BakStatus,
		Msg:        b.BakMsg,
		FileSize:   b.FileSize,
		FileName:   b.FileName,
		BakTime:    time.Now(),
		IsDeleted:  0,
	}
	return historyDB.Save(context.Background(), database.Gorm)
}

func AfterBak(b *Handler) {
	//判断是否启动OSS保存
	if b.OssConfig.IsOssSave == 1 {
		FileName := b.FileName
		Endpoint := b.OssConfig.Endpoint
		Accesskey := b.OssConfig.OssAccess
		Secretkey := b.OssConfig.OssSecret
		BucketName := b.OssConfig.BucketName
		Directory := b.OssConfig.Directory
		switch b.OssConfig.OssType {
		case 0:
			log.Logger.Debug("OSSType为0保存至minio中")
			log.Logger.Infof("%s:%s开始保存至阿里云对象存储OSS", b.Host, b.DbName)
			if err := alioss.AliOssUploadFile(FileName, Endpoint, Accesskey, Secretkey, BucketName, Directory); err != nil {
				log.Logger.Errorf("%s:%s保存阿里云对象存储OSS失败:%v", b.Host, b.DbName, err.Error())
				b.OssStatus = 0
				return
			}
			log.Logger.Infof("%s:%s阿里云对象存储OSS上传成功", b.Host, b.DbName)
			b.OssStatus = 1
		case 1:
			log.Logger.Debug("OSSType为1保存至minio中")
			client := minio.NewClient(Endpoint, Accesskey, Secretkey, BucketName, Directory, b.FileName)
			if err := client.UploadFile(); err != nil {
				log.Logger.Errorf("%s:%s保存Minio存储失败:%v", b.Host, b.DbName, err.Error())
				b.OssStatus = 0
				return
			}
			log.Logger.Infof("%s:%s保存Minio对象存储上传成功", b.Host, b.DbName)
			b.OssStatus = 1
		}
	}
	//判断是否启动钉钉提醒
	if b.DingConfig.IsDingSend == 1 {
		baktime := time.Now().Format("2006年01月02日15:04:01")
		if config.GetBoolConf("dingProxyAgent", "enable") {
			log.Logger.Infof("%s:%s调用钉钉代理发送钉钉消息", b.Host, b.DbName)
			msg := fmt.Sprintf("\n- 备份时间:%v\n- 备份状态:%s\n- OSS上传状态:%s\n- 备份文件目录:%s\n![screenshot](http://web.ts-it.cn/img/index/section_4_img.png)\n", baktime, b.BakMsg, pkg.StatusConversion(b.OssStatus), b.FileName)
			dingSender := dingproxy.NewDingSender(b.DingConfig.DingAccessToken, b.DingConfig.DingSecret, msg)
			data, err := dingSender.SendMarkdown()
			if err != nil {
				b.DingStatus = 0
				log.Logger.Error("钉钉消息发送失败", err)
				return
			}
			log.Logger.Debug("钉钉发送响应结果:", data)
		} else {
			log.Logger.Infof("%s:%s使用自身能力发送钉钉消息", b.Host, b.DbName)
			markContent := map[string]string{
				"title": b.Host + b.DbName + "备份状态",
				"text": fmt.Sprintf(
					"\n- 备份时间:%v\n- 备份状态:%s\n- OSS上传状态:%s\n- 备份文件目录:%s\n![screenshot](%v)\n-备注:%v",
					baktime, b.BakMsg, pkg.StatusConversion(b.OssStatus), b.FileName, config.GetStringConf("base", "photoUrl"), config.GetStringConf("base", "content")),
			}
			webhook := ding.Webhook{AtAll: true, Secret: b.DingConfig.DingSecret, AccessToken: b.DingConfig.DingAccessToken}
			if err := webhook.SendMarkDown(markContent); err != nil {
				b.DingStatus = 0
				log.Logger.Error("钉钉消息发送失败", err.Error())
				return
			}
		}
		// 钉钉消息发送成功，更新状态
		b.DingStatus = 1
		log.Logger.Infof("%s:%s发送钉钉消息成功", b.Host, b.DbName)
	}
}
