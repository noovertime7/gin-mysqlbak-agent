package core

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/log"
	"backupAgent/staging/src/elasticbak"
	"context"
	"errors"
	"github.com/robfig/cron/v3"
	"sync"
	"time"
)

type esBakHandler struct {
	c       *cron.Cron
	esBaker elasticbak.EsBaker
	cronJob map[string]*cron.Cron
	lock    sync.RWMutex
	info    *dao.EsTaskDetail
}

func NewEsBakHandler(detail *dao.EsTaskDetail) (*esBakHandler, error) {
	baker, err := elasticbak.NewEsBaker(&elasticbak.EsHostInfo{
		Host:     detail.ESTaskInfo.Host,
		UserName: detail.ESTaskInfo.Username,
		Password: detail.ESTaskInfo.Password,
	})
	if err != nil {
		return nil, err
	}
	return &esBakHandler{
		c:       cron.New(),
		esBaker: baker,
		cronJob: make(map[string]*cron.Cron),
		info:    detail,
	}, nil
}

func (e *esBakHandler) Start() error {
	e.lock.RLock()
	defer e.lock.RUnlock()
	id, err := e.c.AddJob(e.info.ESTaskInfo.BackupCycle, e)
	if err != nil {
		return err
	}
	e.cronJob[e.info.ESTaskInfo.Index] = e.c
	e.c.Start()
	log.Logger.Infof("创建Elastic备份任务%v,备份索引:%s,备份周期:%s", id, e.info.ESTaskInfo.Index, e.info.ESTaskInfo.BackupCycle)
	return nil
}

func (e *esBakHandler) Run() {
	curTime := time.Now().Format("2006-01-02-15-04")
	snapName := e.info.ESTaskInfo.Index + "_" + curTime
	if err := e.esBaker.CreateSnapshot(context.TODO(), snapName); err != nil {
		log.Logger.Error("创建快照失败", err)
		return
	}
	log.Logger.Infof("创建快照成功,快照名:%v", snapName)
}

func (e *esBakHandler) Stop() error {
	if err := e.isStart(); err != nil {
		return err
	}
	e.lock.Lock()
	defer e.lock.Unlock()
	log.Logger.Debug("当前备份任务列表", e.cronJob)
	for index, corn := range e.cronJob {
		if index == e.info.ESTaskInfo.Index {
			corn.Stop()
		}
	}
	delete(e.cronJob, e.info.ESTaskInfo.Index)
	log.Logger.Info("停止任务成功", e.cronJob)
	return nil
}

func (e *esBakHandler) isStart() error {
	if _, ok := e.cronJob[e.info.ESTaskInfo.Index]; !ok {
		return errors.New("当前任务未启动")
	}
	return nil
}
