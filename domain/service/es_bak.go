package service

import (
	"backupAgent/domain/core"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/database"
	"context"
)

type esBakService struct{}

func NewEsBakService() *esBakService {
	return &esBakService{}
}

func (e *esBakService) Start(ctx context.Context, taskID int64) error {
	taskInfo := &dao.EsTaskDB{ID: taskID}
	detail, err := taskInfo.TaskDetail(ctx, database.Gorm, taskInfo)
	if err != nil {
		return err
	}
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
	taskInfo := &dao.EsTaskDB{ID: taskID}
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
