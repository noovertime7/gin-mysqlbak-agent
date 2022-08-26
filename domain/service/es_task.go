package service

import (
	"backupAgent/domain/config"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/database"
	"backupAgent/proto/backupAgent/esbak"
	"context"
	"time"
)

type ESTaskService struct{}

func NewEsTaskService() *ESTaskService {
	return &ESTaskService{}
}

func (e *ESTaskService) TaskADD(ctx context.Context, taskInfo *esbak.EsBakTaskADDInput) error {
	esTaskDB := &dao.EsTaskDB{
		ServiceName:   config.GetStringConf("base", "serviceName"),
		Host:          taskInfo.EsHost,
		Password:      taskInfo.EsPassword,
		Username:      taskInfo.EsUser,
		BackupCycle:   taskInfo.BackupCycle,
		KeepNumber:    taskInfo.KeepNumber,
		Index:         taskInfo.Index,
		IsAllIndexBak: taskInfo.IsEsBakAll,
		IsDelete:      0,
		Status:        1,
		UpdatedAt:     time.Now(),
		CreatedAt:     time.Now(),
	}
	return esTaskDB.Save(ctx, database.Gorm)
}

func (e *ESTaskService) TaskDelete(ctx context.Context, id int64) error {
	var esDB *dao.EsTaskDB
	esDB.ID = id
	esTaskDB, err := esDB.Find(ctx, database.Gorm, esDB)
	if err != nil {
		return err
	}
	esTaskDB.IsDelete = 1
	return esTaskDB.Updates(ctx, database.Gorm)
}

func (e *ESTaskService) TaskUpdate(ctx context.Context, taskInfo *esbak.EsBakTaskUpdateInput) error {
	esTaskDB := &dao.EsTaskDB{
		ID:            taskInfo.ID,
		Host:          taskInfo.EsHost,
		Password:      taskInfo.EsPassword,
		Username:      taskInfo.EsUser,
		BackupCycle:   taskInfo.BackupCycle,
		KeepNumber:    taskInfo.KeepNumber,
		Index:         taskInfo.Index,
		IsAllIndexBak: taskInfo.IsEsBakAll,
		UpdatedAt:     time.Now(),
	}
	return esTaskDB.Updates(ctx, database.Gorm)
}

func (e *ESTaskService) GetTaskList(ctx context.Context, taskInfo *esbak.EsTaskListInput) (*esbak.EsTaskListOutPut, error) {
	var esDB *dao.EsTaskDB
	list, total, err := esDB.PageList(ctx, database.Gorm, taskInfo)
	if err != nil {
		return nil, err
	}
	var outList []*esbak.EsTaskListOutPutItem
	for _, listIterm := range list {
		outItem := &esbak.EsTaskListOutPutItem{
			ID:          listIterm.ID,
			EsHost:      listIterm.Host,
			Index:       listIterm.Index,
			BackupCycle: listIterm.BackupCycle,
			KeepNumber:  listIterm.KeepNumber,
		}
		outList = append(outList, outItem)
	}
	out := &esbak.EsTaskListOutPut{
		Total:                total,
		EsTaskListOutPutItem: outList,
	}
	return out, nil
}

func (e *ESTaskService) GetTaskDetail(ctx context.Context, id int64) (*esbak.EsTaskDetailOutPut, error) {
	taskInfo := &dao.EsTaskDB{}
	detail, err := taskInfo.TaskDetail(ctx, database.Gorm, &dao.EsTaskDB{ID: id})
	if err != nil {
		return nil, err
	}
	out := &esbak.EsTaskDetailOutPut{EsTaskInfo: &esbak.EsTaskInfo{
		EsHost:      detail.ESTaskInfo.Host,
		EsUser:      detail.ESTaskInfo.Username,
		EsPassword:  detail.ESTaskInfo.Password,
		Index:       detail.ESTaskInfo.Index,
		BackupCycle: detail.ESTaskInfo.BackupCycle,
		KeepNumber:  detail.ESTaskInfo.KeepNumber,
		IsEsBakAll:  detail.ESTaskInfo.IsAllIndexBak,
	}}
	return out, nil
}
