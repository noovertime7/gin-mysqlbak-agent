package service

import (
	"backupAgent/domain/config"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/proto/backupAgent/esbak"
	"backupAgent/proto/backupAgent/task"
	"context"
	"database/sql"
	"time"
)

type ESTaskService struct{}

func NewEsTaskService() *ESTaskService {
	return &ESTaskService{}
}

func (e *ESTaskService) TaskADD(ctx context.Context, taskInfo *esbak.EsBakTaskADDInput) error {
	esTaskDB := &dao.TaskInfo{
		ServiceName: config.GetStringConf("base", "serviceName"),
		HostID:      taskInfo.HostID,
		BackupCycle: taskInfo.BackupCycle,
		KeepNumber:  taskInfo.KeepNumber,
		IsAllDBBak:  1,
		IsDelete:    sql.NullInt64{Int64: 0, Valid: true},
		Status:      0,
		UpdatedAt:   time.Now(),
		CreatedAt:   time.Now(),
	}
	return esTaskDB.Save(ctx, database.Gorm)
}

func (e *ESTaskService) TaskDelete(ctx context.Context, id int64) error {
	esDB := &dao.TaskInfo{Id: id}
	esTaskDB, err := esDB.Find(ctx, database.Gorm, esDB)
	if err != nil {
		return err
	}
	esTaskDB.IsDelete = sql.NullInt64{Int64: 1, Valid: true}
	return esTaskDB.Updates(ctx, database.Gorm)
}

func (e *ESTaskService) TaskRestore(ctx context.Context, id int64) error {
	esDB := &dao.TaskInfo{Id: id}
	esTaskDB, err := esDB.Find(ctx, database.Gorm, esDB)
	if err != nil {
		return err
	}
	esTaskDB.IsDelete = sql.NullInt64{Int64: 0, Valid: true}
	return esTaskDB.Updates(ctx, database.Gorm)
}

func (e *ESTaskService) TaskUpdate(ctx context.Context, taskInfo *esbak.EsBakTaskUpdateInput) error {
	esTaskDB := &dao.TaskInfo{
		Id:          taskInfo.ID,
		HostID:      taskInfo.HostID,
		BackupCycle: taskInfo.BackupCycle,
		KeepNumber:  taskInfo.KeepNumber,
		UpdatedAt:   time.Now(),
	}
	return esTaskDB.Updates(ctx, database.Gorm)
}

func (e *ESTaskService) GetTaskList(ctx context.Context, taskInfo *esbak.EsTaskListInput) (*esbak.EsTaskListOutPut, error) {
	var esDB *dao.TaskInfo
	list, total, err := esDB.PageList(ctx, database.Gorm, &task.TaskListInput{
		HostID:   taskInfo.HostID,
		Info:     taskInfo.Info,
		PageNo:   taskInfo.PageNo,
		PageSize: taskInfo.PageSize,
	})
	if err != nil {
		return nil, err
	}
	var outList []*esbak.EsTaskListOutPutItem
	for _, listIterm := range list {
		hostDB := &dao.HostDatabase{Id: listIterm.HostID}
		host, err := hostDB.Find(ctx, database.Gorm, hostDB)
		if err != nil {
			return nil, err
		}
		outItem := &esbak.EsTaskListOutPutItem{
			ID:          listIterm.Id,
			EsHost:      host.Host,
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
		PageSize:             taskInfo.PageSize,
		PageNo:               taskInfo.PageNo,
	}
	return out, nil
}

func (e *ESTaskService) GetTaskDetail(ctx context.Context, id int64) (*esbak.EsTaskDetailOutPut, error) {
	taskInfo := &dao.TaskInfo{}
	detail, err := taskInfo.TaskDetail(ctx, database.Gorm, &dao.TaskInfo{Id: id})
	if err != nil {
		return nil, err
	}
	hostDB := &dao.HostDatabase{Id: taskInfo.HostID}
	host, err := hostDB.Find(ctx, database.Gorm, hostDB)
	if err != nil {
		return nil, err
	}
	out := &esbak.EsTaskDetailOutPut{EsTaskInfo: &esbak.EsTaskInfo{
		EsHost:      host.Host,
		EsUser:      host.User,
		EsPassword:  host.Password,
		BackupCycle: detail.Info.BackupCycle,
		KeepNumber:  detail.Info.KeepNumber,
		Status:      pkg.IntToBool(detail.Info.Status),
		CreateAt:    detail.Info.CreatedAt.Format("2006年01月02日15:04:01"),
	}}
	return out, nil
}
