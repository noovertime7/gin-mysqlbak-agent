package clean

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/alioss"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/pkg/minio"
	"context"
	"database/sql"
	"errors"
	"os"
	"path/filepath"
	"time"
)

type MysqlCleanHandler struct {
	HistoryID int64
	FilePath  string
	FileName  string
	OssInfo   *OssInfo
}

func NewMysqlCleanHandler(filePath string, HistoryID int64, oss *OssInfo) *MysqlCleanHandler {
	filename, _ := filepath.Split(filePath)
	return &MysqlCleanHandler{FilePath: filePath, FileName: filename, OssInfo: oss, HistoryID: HistoryID}
}

type OssInfo struct {
	IsOssSave       int64
	OssType         int64
	Endpoint        string
	AccessKey       string
	SecretAccessKey string
	BucketName      string
	Dir             string
	Filepath        string
}

func (m *MysqlCleanHandler) CleanFile() error {
	//删除本地文件
	if err := m.CleanLocalFile(); err != nil {
		return err
	}
	//3、判断任务是否启用oss，如果启用，调用oss删除方法
	//删除oss文件
	if m.OssInfo.IsOssSave == 1 {
		if err := m.CleanOssFile(); err != nil {
			return err
		}
	}
	//更改状态
	return m.Store()
}

func (m *MysqlCleanHandler) CleanLocalFile() error {
	log.Logger.Infof("mysqlCleanHandler删除本地文件:%s", m.FilePath)
	if m.isLocalFileExists(m.FilePath) {
		return os.Remove(m.FilePath)
	}
	log.Logger.Warningf("%s文件未找到", m.FilePath)
	return nil
}

func (m *MysqlCleanHandler) isLocalFileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return false
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func (m *MysqlCleanHandler) CleanOssFile() error {
	switch m.OssInfo.OssType {
	//0代表阿里oss
	case 0:
		return m.cleanAliOssData()
	case 1:
		return m.cleanMinioData()
	default:
		return errors.New("存储类型不正确")
	}
}

func (m *MysqlCleanHandler) cleanMinioData() error {
	mc, err := minio.NewClient(m.OssInfo.Endpoint, m.OssInfo.AccessKey, m.OssInfo.SecretAccessKey, m.OssInfo.BucketName, m.OssInfo.Dir, m.OssInfo.Filepath)
	if err != nil {
		return err
	}
	log.Logger.Info("开始清理minio过期数据")
	return mc.Remove()
}

func (m *MysqlCleanHandler) cleanAliOssData() error {
	oc, err := alioss.NewClient(m.OssInfo.Filepath, m.OssInfo.Endpoint, m.OssInfo.AccessKey, m.OssInfo.SecretAccessKey, m.OssInfo.BucketName, m.OssInfo.Dir)
	if err != nil {
		return err
	}
	return oc.Remove()
}

func (m *MysqlCleanHandler) Store() error {
	historyDB := &dao.BakHistory{Id: m.HistoryID}
	history, err := historyDB.Find(context.TODO(), database.Gorm, historyDB)
	if err != nil {
		return err
	}
	history.IsCleaned = 1
	history.CleanedAt = sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}
	return history.Updates(context.TODO(), database.Gorm)
}

//4、判断是否启用钉钉，如果启用钉钉发送删除成功消息
