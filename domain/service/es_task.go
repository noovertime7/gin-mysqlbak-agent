package service

import (
	"backupAgent/domain/config"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/proto/backupAgent/esbak"
	"context"
	"github.com/olivere/elastic"
	"time"
)

type ESTaskService struct{}

func NewEsTaskService() *ESTaskService {
	return &ESTaskService{}
}

func (e *ESTaskService) TaskADD(ctx context.Context, taskInfo *esbak.EsBakTaskADDInput) error {
	//es主机检查
	if err := e.EsHostCheck(taskInfo.EsHost, taskInfo.EsUser, taskInfo.EsPassword); err != nil {
		return err
	}
	esTaskDB := &dao.EsTaskDB{
		ServiceName: config.GetStringConf("base", "serviceName"),
		Host:        taskInfo.EsHost,
		Password:    taskInfo.EsPassword,
		Username:    taskInfo.EsUser,
		BackupCycle: taskInfo.BackupCycle,
		KeepNumber:  taskInfo.KeepNumber,
		IsDelete:    0,
		Status:      0,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	return esTaskDB.Save(ctx, database.Gorm)
}

func (e *ESTaskService) TaskDelete(ctx context.Context, id int64) error {
	esDB := &dao.EsTaskDB{ID: id}
	esTaskDB, err := esDB.Find(ctx, database.Gorm, esDB)
	if err != nil {
		return err
	}
	esTaskDB.IsDelete = 1
	return esTaskDB.Updates(ctx, database.Gorm)
}

func (e *ESTaskService) TaskUpdate(ctx context.Context, taskInfo *esbak.EsBakTaskUpdateInput) error {
	//es主机检查
	if err := e.EsHostCheck(taskInfo.EsHost, taskInfo.EsUser, taskInfo.EsPassword); err != nil {
		return err
	}
	esTaskDB := &dao.EsTaskDB{
		ID:          taskInfo.ID,
		Host:        taskInfo.EsHost,
		Password:    taskInfo.EsPassword,
		Username:    taskInfo.EsUser,
		BackupCycle: taskInfo.BackupCycle,
		KeepNumber:  taskInfo.KeepNumber,
		UpdatedAt:   time.Now(),
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
			BackupCycle: pkg.CornExprToTime(listIterm.BackupCycle),
			KeepNumber:  listIterm.KeepNumber,
			Status:      pkg.IntToBool(listIterm.Status),
			CreateAt:    listIterm.CreatedAt.Format("2006年01月02日15:04:01"),
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
		BackupCycle: detail.ESTaskInfo.BackupCycle,
		KeepNumber:  detail.ESTaskInfo.KeepNumber,
		Status:      pkg.IntToBool(detail.ESTaskInfo.Status),
		CreateAt:    detail.ESTaskInfo.CreatedAt.Format("2006年01月02日15:04:01"),
	}}
	return out, nil
}

func (e *ESTaskService) EsHostCheck(host, user, password string) error {
	if _, err := elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetBasicAuth(user, password),
		elastic.SetSniff(false)); err != nil {
		return err
	}
	return nil
}
