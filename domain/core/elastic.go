package core

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/staging/src/elasticbak"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/olivere/elastic"
	"github.com/robfig/cron/v3"
	"strings"
	"sync"
	"time"
)

type esBakHandler struct {
	c            *cron.Cron
	esBaker      elasticbak.EsBaker
	lock         sync.RWMutex
	info         *dao.TaskDetail
	snapShotName string
}

var RunningCronJob = make(map[int64]*cron.Cron)

func NewEsBakHandler(detail *dao.TaskDetail) (*esBakHandler, error) {
	baker, err := elasticbak.NewEsBaker(&elasticbak.EsHostInfo{
		Host:     detail.Host.Host,
		UserName: detail.Host.User,
		Password: detail.Host.Password,
	})
	if err != nil {
		return nil, err
	}
	return &esBakHandler{
		c:       cron.New(),
		esBaker: baker,
		info:    detail,
	}, nil
}

func (e *esBakHandler) Start() error {
	e.lock.RLock()
	defer e.lock.RUnlock()
	id, err := e.c.AddJob(e.info.Info.BackupCycle, e)
	if err != nil {
		return err
	}
	RunningCronJob[e.info.Info.Id] = e.c
	e.c.Start()
	log.Logger.Infof("创建Elastic备份任务%v,备份任务ID:%d,备份周期:%s", id, e.info.Info.Id, e.info.Info.BackupCycle)
	return nil
}

func (e *esBakHandler) Run() {
	e.snapShotName = time.Now().Format("2006-01-02-15-04-01")
	if err := e.esBaker.CreateSnapshot(context.TODO(), e.snapShotName); err != nil {
		log.Logger.Error("创建快照失败", err)
		if err := e.Store(false, err.Error()); err != nil {
			log.Logger.Error("快照失败,保存数据库失败", err)
			return
		}
		log.Logger.Infof("快照失败,保存数据库成功:%s", e.snapShotName)
		return
	}
	log.Logger.Infof("创建快照成功,快照名:%v", e.snapShotName)
	log.Logger.Info("等待快照完成,休眠15秒")
	time.Sleep(15 * time.Second)
	if err := e.Store(true, ""); err != nil {
		log.Logger.Error("快照成功,保存数据库失败", err)
		return
	}
	log.Logger.Infof("快照成功,保存数据库成功:%s", e.snapShotName)
}

func (e *esBakHandler) Stop() error {
	log.Logger.Debugf("当前备份任务列表%v,传入ID:%v", RunningCronJob, e.info.Info.Id)
	if err := e.isStart(); err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	for index, corn := range RunningCronJob {
		if index == e.info.Info.Id {
			corn.Stop()
		}
	}
	delete(RunningCronJob, e.info.Info.Id)
	log.Logger.Info("停止任务成功", RunningCronJob)
	return nil
}

func (e *esBakHandler) GetSnapshotDetail(ctx context.Context, snapName string) (*elastic.Snapshot, error) {
	data, err := e.esBaker.GetSnapshot(ctx)
	if err != nil {
		return nil, err
	}
	for _, snap := range data.Snapshots {
		if snap.Snapshot == snapName {
			return snap, nil
		}
	}
	return nil, err
}

func (e *esBakHandler) Store(success bool, message string) error {
	//如果创建失败
	if !success {
		esHistoryDb := &dao.ESHistoryDB{
			TaskID:            e.info.Info.Id,
			Snapshot:          "快照失败",
			Repository:        e.esBaker.GetRepositoryName(),
			UUID:              "快照失败",
			Version:           "快照失败",
			Indices:           "快照失败",
			State:             "failed",
			StartTime:         time.Now(),
			StartTimeInMillis: 0,
			EndTime:           time.Now(),
			BakTime:           time.Now(),
			EndTimeInMillis:   0,
			DurationInMillis:  0,
			Message:           message,
			IsDeleted:         0,
			Status: sql.NullInt64{
				Int64: 0,
				Valid: true,
			},
		}
		return esHistoryDb.Save(context.TODO(), database.Gorm)
	}
	detail, err := e.GetSnapshotDetail(context.TODO(), e.snapShotName)
	if err != nil {
		return err
	}
	esHistoryDb := &dao.ESHistoryDB{
		TaskID:            e.info.Info.Id,
		Snapshot:          detail.Snapshot,
		Repository:        e.esBaker.GetRepositoryName(),
		UUID:              detail.UUID,
		Version:           detail.Version,
		Indices:           fmt.Sprintf(strings.Join(detail.Indices, ";")),
		State:             detail.State,
		StartTime:         detail.StartTime,
		StartTimeInMillis: detail.StartTimeInMillis,
		EndTime:           detail.EndTime,
		BakTime:           time.Now(),
		EndTimeInMillis:   detail.EndTimeInMillis,
		DurationInMillis:  detail.DurationInMillis,
		Message:           detail.State,
		IsDeleted:         0,
		Status: sql.NullInt64{
			Int64: 1,
			Valid: true,
		},
	}
	return esHistoryDb.Save(context.TODO(), database.Gorm)
}

func (e *esBakHandler) isStart() error {
	e.lock.RLock()
	defer e.lock.RUnlock()
	if _, ok := RunningCronJob[e.info.Info.Id]; !ok {
		return errors.New("当前任务未启动")
	}
	return nil
}
