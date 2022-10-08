package service

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/proto/backupAgent/task"
	"context"
	"database/sql"
	"errors"
	"net"
	"strings"
	"time"
)

type TaskService struct{}

func (t *TaskService) TaskAdd(ctx context.Context, taskInfo *task.TaskAddInput) error {
	//添加任务进行数据库检测
	if err := TaskTest(ctx, taskInfo.HostID, taskInfo.DBName, taskInfo.Endpoint); err != nil {
		return err
	}
	taskDb := &dao.TaskInfo{
		HostID:      taskInfo.HostID,
		ServiceName: taskInfo.ServiceName,
		DBName:      taskInfo.DBName,
		BackupCycle: taskInfo.BackupCycle,
		KeepNumber:  taskInfo.KeepNumber,
		IsAllDBBak:  taskInfo.IsAllDBBak,
		IsDelete:    sql.NullInt64{Int64: 0, Valid: true},
		Status:      0,
	}
	tx := database.Gorm.Begin()
	if err := taskDb.Save(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	ossDB := &dao.OssDatabase{
		TaskID:     taskDb.Id,
		IsOssSave:  taskInfo.IsOssSave,
		OssType:    taskInfo.OssType,
		Endpoint:   taskInfo.Endpoint,
		OssAccess:  taskInfo.OssAccess,
		OssSecret:  taskInfo.OssSecret,
		BucketName: taskInfo.BucketName,
		Directory:  taskInfo.Directory,
	}
	if err := ossDB.Save(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	dingdb := &dao.DingDatabase{
		TaskID:          taskDb.Id,
		IsDingSend:      taskInfo.IsDingSend,
		DingAccessToken: taskInfo.DingAccessToken,
		DingSecret:      taskInfo.DingSecret,
	}
	if err := dingdb.Save(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (t *TaskService) TaskDelete(ctx context.Context, taskinfo *task.TaskIDInput) error {
	taskDB := &dao.TaskInfo{Id: taskinfo.ID}
	task, err := taskDB.Find(ctx, database.Gorm, taskDB)
	if err != nil {
		return err
	}
	task.IsDelete = sql.NullInt64{Int64: 1, Valid: true}
	task.DeletedAt = time.Now()
	return task.Save(ctx, database.Gorm)
}

func (t *TaskService) TaskRestore(ctx context.Context, taskinfo *task.TaskIDInput) error {
	taskDB := &dao.TaskInfo{Id: taskinfo.ID}
	task, err := taskDB.Find(ctx, database.Gorm, taskDB)
	if err != nil {
		return err
	}
	task.IsDelete = sql.NullInt64{Int64: 0, Valid: true}
	task.DeletedAt = time.Now()
	return task.Save(ctx, database.Gorm)
}

func (t *TaskService) TaskUpdate(ctx context.Context, taskInfo *task.TaskUpdateInput) error {
	//添加任务进行数据库检测
	if err := TaskTest(ctx, taskInfo.HostID, taskInfo.DBName, taskInfo.Endpoint); err != nil {
		return err
	}
	taskDb := &dao.TaskInfo{
		Id:          taskInfo.ID,
		DBName:      taskInfo.DBName,
		BackupCycle: taskInfo.BackupCycle,
		KeepNumber:  taskInfo.KeepNumber,
		IsAllDBBak:  taskInfo.IsAllDBBak,
		ServiceName: taskInfo.ServiceName,
	}
	tx := database.Gorm.Begin()
	if err := taskDb.Updates(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	ossDB := &dao.OssDatabase{
		TaskID:     taskDb.Id,
		IsOssSave:  taskInfo.IsOssSave,
		OssType:    taskInfo.OssType,
		Endpoint:   taskInfo.Endpoint,
		OssAccess:  taskInfo.OssAccess,
		OssSecret:  taskInfo.OssSecret,
		BucketName: taskInfo.BucketName,
		Directory:  taskInfo.Directory,
	}
	if err := ossDB.UpdatesByMap(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	dingdb := &dao.DingDatabase{
		TaskID:          taskDb.Id,
		IsDingSend:      taskInfo.IsDingSend,
		DingAccessToken: taskInfo.DingAccessToken,
		DingSecret:      taskInfo.DingSecret,
	}
	if err := dingdb.UpdatesByMap(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}

func (t *TaskService) TaskList(ctx context.Context, taskinfo *task.TaskListInput) (*task.TaskListOutPut, error) {
	log.Logger.Debug("Service层查询Task列表，入参:", taskinfo)
	taskDB := &dao.TaskInfo{}
	list, total, err := taskDB.PageList(ctx, database.Gorm, taskinfo)
	if err != nil {
		return nil, err
	}
	var outList []*task.TaskListItem
	for _, listIterm := range list {
		hostDB := &dao.HostDatabase{Id: listIterm.HostID}
		databseres, err := hostDB.Find(ctx, database.Gorm, hostDB)
		if err != nil {
			return nil, err
		}
		outItem := &task.TaskListItem{
			ID:          listIterm.Id,
			Host:        databseres.Host,
			HostID:      listIterm.HostID,
			DBName:      listIterm.DBName,
			BackupCycle: pkg.CornExprToTime(listIterm.BackupCycle),
			KeepNumber:  listIterm.KeepNumber,
			CreateAt:    listIterm.CreatedAt.Format("2006年01月02日15:04"),
			UpdateAt:    listIterm.UpdatedAt.Format("2006年01月02日15:04"),
			Status:      listIterm.Status,
			IsDeleted:   listIterm.IsDelete.Int64,
			DeletedAt:   listIterm.DeletedAt.Format("2006年01月02日15:04"),
		}
		outList = append(outList, outItem)
	}
	out := &task.TaskListOutPut{
		Total:        total,
		TaskListItem: outList,
	}
	log.Logger.Debug("Service层查询Task列表，出参:", out)
	return out, nil
}

func (t *TaskService) UnscopedTaskList(ctx context.Context, taskinfo *task.TaskListInput) (*task.TaskListOutPut, error) {
	log.Logger.Debug("UnscopedService层查询Task列表，入参:", taskinfo)
	taskDB := &dao.TaskInfo{}
	list, total, err := taskDB.UnscopedPageList(ctx, database.Gorm, taskinfo)
	if err != nil {
		return nil, err
	}
	var outList []*task.TaskListItem
	for _, listIterm := range list {
		hostDB := &dao.HostDatabase{Id: listIterm.HostID}
		databseres, err := hostDB.Find(ctx, database.Gorm, hostDB)
		if err != nil {
			return nil, err
		}
		outItem := &task.TaskListItem{
			ID:          listIterm.Id,
			Host:        databseres.Host,
			HostID:      listIterm.HostID,
			DBName:      listIterm.DBName,
			BackupCycle: pkg.CornExprToTime(listIterm.BackupCycle),
			KeepNumber:  listIterm.KeepNumber,
			CreateAt:    listIterm.CreatedAt.Format("2006年01月02日15:04"),
			UpdateAt:    listIterm.UpdatedAt.Format("2006年01月02日15:04"),
			Status:      listIterm.Status,
			IsDeleted:   listIterm.IsDelete.Int64,
			DeletedAt:   listIterm.DeletedAt.Format("2006年01月02日15:04"),
		}
		outList = append(outList, outItem)
	}
	out := &task.TaskListOutPut{
		Total:        total,
		TaskListItem: outList,
	}
	log.Logger.Debug("UnscopedService层查询Task列表，出参:", out)
	return out, nil
}

func (t *TaskService) TaskDetail(ctx context.Context, taskInfo *task.TaskIDInput) (*task.TaskDetailOutPut, error) {
	taskDb := &dao.TaskInfo{Id: taskInfo.ID}
	taskDetailInfo, err := taskDb.TaskDetail(ctx, database.Gorm, taskDb)
	if err != nil {
		return nil, err
	}
	out := &task.TaskDetailOutPut{
		Host:            taskDetailInfo.Host.Host,
		Content:         taskDetailInfo.Host.Content,
		HostStatus:      taskDetailInfo.Host.HostStatus,
		IsDingSend:      taskDetailInfo.Ding.IsDingSend,
		DingAccessToken: taskDetailInfo.Ding.DingAccessToken,
		DingSecret:      taskDetailInfo.Ding.DingSecret,
		IsOssSave:       taskDetailInfo.Oss.IsOssSave,
		OssType:         taskDetailInfo.Oss.OssType,
		Endpoint:        taskDetailInfo.Oss.Endpoint,
		OssAccess:       taskDetailInfo.Oss.OssAccess,
		OssSecret:       taskDetailInfo.Oss.OssSecret,
		BucketName:      taskDetailInfo.Oss.BucketName,
		Directory:       taskDetailInfo.Oss.Directory,
		TaskID:          taskDetailInfo.Info.Id,
		HostID:          taskDetailInfo.Info.HostID,
		DBName:          taskDetailInfo.Info.DBName,
		ServiceName:     taskDetailInfo.Info.ServiceName,
		BackupCycle:     taskDetailInfo.Info.BackupCycle,
		KeepNumber:      taskDetailInfo.Info.KeepNumber,
		Status:          taskDetailInfo.Info.Status,
		CreateAt:        taskDetailInfo.Info.CreatedAt.Format("2006年01月02日15:04"),
	}
	return out, nil
}

// TaskTest 在添加任务时，进行数据库连接测试，避免添加无用信息导致备份失败
func TaskTest(ctx context.Context, hid int64, DBName, endpoint string) error {
	hostDB := &dao.HostDatabase{Id: hid}
	host, err := hostDB.Find(ctx, database.Gorm, hostDB)
	if err != nil {
		return err
	}
	log.Logger.Info("开始任务测试")
	if err := HostPingCheck(host.User, host.Password, host.Host, DBName, 1); err != nil {
		return err
	}
	if err := EndPointTest(endpoint); err != nil {
		return err
	}
	return nil
}

// EndPointTest 进行endpoint的一个端口检测
func EndPointTest(endpoint string) error {
	if endpoint == "" {
		return nil
	}
	if temp := strings.Split(endpoint, ":"); len(temp) != 2 {
		return errors.New("您输入的Endpoint地址有误  正确地址: 192.168.1.1:9000")
	}
	_, err := net.DialTimeout("tcp", endpoint, 2*time.Second)
	if err != nil {
		return err
	}
	return nil
}
