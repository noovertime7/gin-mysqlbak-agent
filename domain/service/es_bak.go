package service

import (
	"backupAgent/domain/core"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"context"
	"errors"
)

type esBakService struct{}

func NewEsBakService() *esBakService {
	return &esBakService{}
}

func (e *esBakService) Start(ctx context.Context, taskID int64) error {
	taskInfo := &dao.TaskInfo{Id: taskID, Type: pkg.ElasticHost}
	log.Logger.Debug("elastic启动备份,入参:", taskInfo)
	detail, err := taskInfo.TaskDetail(ctx, database.Gorm, taskInfo)
	if err != nil {
		return err
	}
	if detail.Info.IsDelete.Int64 == 1 {
		return errors.New("任务已被删除,启动失败")
	}
	log.Logger.Debug("elastic host detail ", detail.Host)
	bakHandler, err := core.NewEsBakHandler(detail)
	if err != nil {
		return err
	}
	if err := bakHandler.Start(); err != nil {
		return err
	}
	taskInfo.Status = 1
	return taskInfo.UpdateStatus(ctx, database.Gorm, taskInfo)
}

func (e *esBakService) Stop(ctx context.Context, taskID int64) error {
	taskInfo := &dao.TaskInfo{Id: taskID, Type: pkg.ElasticHost}
	log.Logger.Debug("elastic停止备份,入参:", taskInfo)
	detail, err := taskInfo.TaskDetail(ctx, database.Gorm, taskInfo)
	if err != nil {
		return err
	}
	bakHandler, err := core.NewEsBakHandler(detail)
	if err != nil {
		return err
	}
	if err := bakHandler.Stop(); err != nil {
		return err
	}
	taskInfo.Status = 0
	return taskInfo.UpdateStatus(ctx, database.Gorm, taskInfo)
}
