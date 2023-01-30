package core

import (
	"backupAgent/domain/config"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/alioss"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/pkg/minio"
	"backupAgent/domain/template"
	"context"
	"errors"
	"fmt"
	"github.com/robfig/cron/v3"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

type Handler struct {
	TaskID           int64
	Cron             *cron.Cron
	Host             string
	BackInfoId       string
	PassWord         string
	User             string
	Port             string
	DbName           string
	BackupCycle      string
	KeepNumber       int64
	ISAllDBBak       int64
	DingConfig       *dao.DingDatabase
	OssConfig        *dao.OssDatabase
	DingStatus       int64
	OssStatus        int64
	BakStatus        int64
	BakMsg           string
	FileName         string
	EncryptionStatus int64
	FileSize         int64
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
	c := cron.New()
	if _, ok := CronJob[handler.TaskID]; ok {
		return nil, errors.New("当前备份任务已启动，切勿重复启动")
	}
	CronJob[handler.TaskID] = c
	handler.Cron = c
	return handler, nil
}

func (h *Handler) BeforBak() {
	path, _ := os.Getwd()
	pkg.CreateDir(path + "/bakfile/")
	dir := path + "/bakfile/" + strconv.Itoa(int(h.TaskID))
	baktime := time.Now().Format("2006-01-02-15-04")
	pkg.CreateDir(dir)
	host := strings.Split(h.Host, ":")[0]
	file := dir + "/" + host + "_" + h.DbName + "_" + baktime + ".sql"
	h.FileName = file
}

func (h *Handler) StartBak() error {
	_, err := h.Cron.AddJob(h.BackupCycle, h)
	if err != nil {
		return err
	}
	log.Logger.Infof("创建备份任务成功，任务id：%d", h.TaskID)
	// 备份前准备工作
	h.BeforBak()
	//启动数据库备份服务
	h.Cron.Start()
	return nil
}

func (h *Handler) IsStart(tid int64) bool {
	if _, ok := CronJob[tid]; !ok {
		return false
	}
	return true
}

func (h *Handler) StopBak(tid int64) error {
	if ok := h.IsStart(tid); !ok {
		log.Logger.Warningf("任务ID不存在%d,当前运行队列%v", tid, CronJob)
		return errors.New("任务非运行,停止失败")
	}
	log.Logger.Debug("StopBak", CronJob)
	for id, cronjob := range CronJob {
		if id == tid {
			cronjob.Stop()
			delete(CronJob, tid)
			log.Logger.Infof("任务ID:%d,备份库名:%s 停止成功", id, h.DbName)
			return nil
		}
	}
	return nil
}

func (h *Handler) Run() {
	log.Logger.Info("BakHandler 开始备份数据库")
	h.BeforBak()
	if err := h.RunMysqlBak(); err != nil {
		h.BakStatus = 0
		h.OssStatus = 2
		h.DingStatus = 2
		h.BakMsg = fmt.Sprintf("%s", err)
		h.FileName = "unknown"
		AfterBak(h)
		log.Logger.Error("备份失败,保存备份历史到数据库,停止备份任务,发送消息", err)
		if err := h.StoreDatabase(); err != nil {
			log.Logger.Error("数据库存储失败", err)
			return
		}
		log.Logger.Debug("备份失败,数据库存储成功")
		return
	}

	h.BakMsg = "success"
	h.BakStatus = 1
	h.FileSize = int64(pkg.GetFileSize(h.FileName))
	//首先判定钉钉 oss都操作成功，状态改为1
	h.DingStatus = 2
	h.OssStatus = 2
	//判断是否启动钉钉提醒
	AfterBak(h)
	// 判断是否加密成功，加密成功后直接删除本地文件
	if h.EncryptionStatus == 1 {
		OldFileName := pkg.GetFilePath(h.FileName) + ".sql"
		if err := pkg.CleanLocalFile(OldFileName); err != nil {
			log.Logger.Errorf("备份完成后清理本地文件失败%s", OldFileName)
			h.BakMsg = "备份完成后清理本地文件失败"
		}
	} else {
		h.BakMsg = "加密备份文件失败，上传未加密文件"
	}
	log.Logger.Infof("备份数据库成功,保存备份历史到数据库,备份文件:%s", h.FileName)
	if err := h.StoreDatabase(); err != nil {
		log.Logger.Error("数据库存储失败", err)
		return
	}
	log.Logger.Debug("数据库存储成功")
}

func (h *Handler) MysqlDump() error {
	iphost, port := strings.Split(h.Host, ":")[0], strings.Split(h.Host, ":")[1]
	command := fmt.Sprintf("mysqldump -u%v -p%v -P%v -h%v  %v >  %v", h.User, h.PassWord, port, iphost, h.DbName, h.FileName)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		log.Logger.Error("mysqldump执行失败:", command, " with error: ", err.Error())
		return err
	}
	log.Logger.Infof("mysqldump执行成功:%v:%v", h.Host, h.DbName)
	return nil
}

func (h *Handler) FlushLogs() error {
	iphost, port := strings.Split(h.Host, ":")[0], strings.Split(h.Host, ":")[1]
	command := fmt.Sprintf("mysqladmin -u%v -p%v -P%v -h%v flush-logs", h.User, h.PassWord, port, iphost)
	cmd := exec.Command("sh", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		log.Logger.Error("mysqlFlush执行失败:", command, " with error: ", err.Error())
		return err
	}
	log.Logger.Infof("mysqlFlush执行成功:%v:%v", h.Host, h.DbName)
	return nil
}

func (h *Handler) RunMysqlBak() error {
	if err := h.FlushLogs(); err != nil {
		return err
	}
	return h.MysqlDump()
}

func (h *Handler) StoreDatabase() error {
	historyDB := &dao.BakHistory{
		TaskID:           h.TaskID,
		Host:             h.Host,
		DBName:           h.DbName,
		OssStatus:        h.OssStatus,
		DingStatus:       h.DingStatus,
		BakStatus:        h.BakStatus,
		Msg:              h.BakMsg,
		FileSize:         h.FileSize,
		FileName:         h.FileName,
		EncryptionStatus: h.EncryptionStatus,
		BakTime:          time.Now(),
		IsDeleted:        0,
	}
	return historyDB.Save(context.Background(), database.Gorm)
}

func AfterBak(b *Handler) {
	// 加密备份文件，如果加密失败，上传原有文件，加密成功上传新文件
	enFile, err := pkg.Encryption(b.FileName)
	if err != nil {
		log.Logger.Errorf("%s加密失败%v", b.FileName, err)
		b.EncryptionStatus = 0
	} else {
		b.EncryptionStatus = 1
		b.FileName = enFile
	}
	//判断是否启动OSS保存
	if b.OssConfig.IsOssSave == 1 && b.BakStatus == 1 {
		FileName := b.FileName
		Endpoint := b.OssConfig.Endpoint
		Accesskey := b.OssConfig.OssAccess
		Secretkey := b.OssConfig.OssSecret
		BucketName := b.OssConfig.BucketName
		Directory := b.OssConfig.Directory
		switch b.OssConfig.OssType {
		case 0:
			log.Logger.Debug("OSSType为0保存至阿里云OSS中")
			log.Logger.Infof("%s:%s开始保存至阿里云对象存储OSS", b.Host, b.DbName)
			ossClient, err := alioss.NewClient(FileName, Endpoint, Accesskey, Secretkey, BucketName, Directory)
			if err != nil {
				b.OssStatus = 0
				log.Logger.Errorf("%s:%s创建阿里云对象存储OSS客户端失败:%v", b.Host, b.DbName, err.Error())
				return
			}
			if err := ossClient.AliOssUploadFile(); err != nil {
				log.Logger.Errorf("%s:%s保存阿里云对象存储OSS失败:%v", b.Host, b.DbName, err.Error())
				b.OssStatus = 0
				return
			}
			log.Logger.Infof("%s:%s阿里云对象存储OSS上传成功", b.Host, b.DbName)
			b.OssStatus = 1
		case 1:
			log.Logger.Debug("OSSType为1保存至minio中")
			client, err := minio.NewClient(Endpoint, Accesskey, Secretkey, BucketName, Directory, FileName)
			if err != nil {
				log.Logger.Errorf("%s:%s创建minio客户端失败:%v", b.Host, b.DbName, err.Error())
				b.OssStatus = 0
				return
			}
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
		log.Logger.Infof("%s数据库备份任务开始发送钉钉消息...", b.Host)
		var onFailSend = config.GetBoolConf("base", "onFailDingSend")
		if onFailSend {
			if b.BakStatus != 1 {
				if err := dingSend(b); err != nil {
					log.Logger.Error("钉钉通知发送失败", err)
					return
				}
				log.Logger.Info("钉钉通知发送成功")
			}
			log.Logger.Info("钉钉onFailSend开关开启,数据备份成功，仅在失败状态下发送钉钉消息")
		} else {
			if err := dingSend(b); err != nil {
				log.Logger.Error("钉钉通知发送失败", err)
				b.DingStatus = 0
				return
			}
			b.DingStatus = 1
			log.Logger.Info("钉钉通知发送成功")
		}
	}
}

func dingSend(b *Handler) error {
	templateFactory := template.NewTemplateFactory(config.GetStringConf("base", "serviceName") + "备份状态")
	sa := b.getDingSa()
	info := b.getInfo()
	dingSender := templateFactory.Ding(sa, info)
	msg, err := dingSender.ParseMessage()
	if err != nil {
		return err
	}
	if config.GetBoolConf("dingProxyAgent", "enable") {
		log.Logger.Infof("%s:%s调用钉钉代理发送钉钉消息", b.Host, b.DbName)
		url := config.GetStringConf("dingProxyAgent", "addr")
		return dingSender.SendByProxy(msg, url)
	} else {
		log.Logger.Infof("%s:%s使用自身能力发送钉钉消息", b.Host, b.DbName)
		return dingSender.SendBySelf(msg)
	}
}

func (h *Handler) getDingSa() *template.DingSA {
	return &template.DingSA{
		DingAccessToken: h.DingConfig.DingAccessToken,
		DingSecret:      h.DingConfig.DingSecret,
	}
}

func (h *Handler) getInfo() *template.SendInfo {
	return &template.SendInfo{
		Host:             h.Host,
		ServiceName:      config.GetStringConf("base", "serviceName"),
		DBName:           h.DbName,
		BakStatus:        pkg.StatusConversion(h.BakStatus),
		OssStatus:        pkg.StatusConversion(h.OssStatus),
		EncryptionStatus: pkg.StatusConversion(h.EncryptionStatus),
		BakTime:          pkg.GetTime(time.Now()),
		FileName:         h.FileName,
	}
}
