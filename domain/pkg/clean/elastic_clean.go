package clean

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/database"
	"backupAgent/staging/src/elasticbak"
	"context"
	"database/sql"
	"time"
)

type ElasticCleanHandler struct {
	host      string
	userName  string
	password  string
	snap      string
	historyID int64
}

func NewElasticCleanHandler(host, username, password, snap string, HistoryID int64) *ElasticCleanHandler {
	return &ElasticCleanHandler{
		host:      host,
		userName:  username,
		password:  password,
		snap:      snap,
		historyID: HistoryID,
	}
}

func (e *ElasticCleanHandler) CleanFile() error {
	baker, err := elasticbak.NewEsBaker(&elasticbak.EsHostInfo{
		Host:     e.host,
		UserName: e.userName,
		Password: e.password,
	})
	if err != nil {
		return err
	}
	if err := baker.DeleteSnapshot(context.TODO(), e.snap); err != nil {
		return err
	}
	return e.changeHistoryStatus()
}

func (e *ElasticCleanHandler) changeHistoryStatus() error {
	historyDB := &dao.ESHistoryDB{Id: e.historyID}
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
