package service

import (
	"backupAgent/domain/config"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/proto/backupAgent/task"
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"
)

type TaskService struct{}

func (t *TaskService) TaskAdd(ctx context.Context, taskInfo *task.TaskAddInput) error {
	hostDB := &dao.HostDatabase{Id: taskInfo.HostID}
	host, err := hostDB.Find(ctx, database.Gorm, hostDB)
	if err != nil {
		return err
	}
	//添加任务进行数据库检测
	if err := TestHost(ctx, taskInfo.HostID); err != nil {
		return err
	}
	ok, err := t.isExists(ctx, taskInfo)
	if err != nil {
		return err
	}
	if ok {
		return errors.New("当前数据库已创建任务")
	}
	taskDb := &dao.TaskInfo{
		HostID:      taskInfo.HostID,
		ServiceName: taskInfo.ServiceName,
		DBName:      taskInfo.DBName,
		BackupCycle: taskInfo.BackupCycle,
		KeepNumber:  taskInfo.KeepNumber,
		IsAllDBBak:  taskInfo.IsAllDBBak,
		IsDelete:    sql.NullInt64{Int64: 0, Valid: true},
		Type:        host.Type,
		Status:      0,
	}
	tx := database.Gorm.Begin()
	if err := taskDb.Save(ctx, tx); err != nil {
		tx.Rollback()
		return err
	}
	//如果主机类型为elastic，不需要创建后面的数据
	if hostDB.Type == pkg.ElasticHost {
		tx.Commit()
		return nil
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

func (t *TaskService) isExists(ctx context.Context, taskInfo *task.TaskAddInput) (bool, error) {
	taskDB := &dao.TaskInfo{DBName: taskInfo.DBName, HostID: taskInfo.HostID}
	ts, err := taskDB.Find(ctx, database.Gorm, taskDB)
	if err != nil {
		return false, err
	}
	return ts.Id != 0, nil
}

func (t *TaskService) TaskDelete(ctx context.Context, taskinfo *task.TaskIDInput) error {
	taskDB := &dao.TaskInfo{Id: taskinfo.ID}
	task, err := taskDB.Find(ctx, database.Gorm, taskDB)
	if err != nil {
		return err
	}
	task.IsDelete = sql.NullInt64{Int64: 1, Valid: true}
	task.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return task.Save(ctx, database.Gorm)
}

func (t *TaskService) TaskDestroy(ctx context.Context, taskinfo *task.TaskIDInput) error {
	taskDB := &dao.TaskInfo{Id: taskinfo.ID}
	task, err := taskDB.Find(ctx, database.Gorm, taskDB)
	if err != nil {
		return err
	}
	return task.Delete(ctx, database.Gorm)
}

func (t *TaskService) TaskRestore(ctx context.Context, taskinfo *task.TaskIDInput) error {
	taskDB := &dao.TaskInfo{Id: taskinfo.ID}
	task, err := taskDB.Find(ctx, database.Gorm, taskDB)
	if err != nil {
		return err
	}
	task.IsDelete = sql.NullInt64{Int64: 0, Valid: true}
	task.DeletedAt = sql.NullTime{Time: time.Now(), Valid: true}
	return task.Save(ctx, database.Gorm)
}

func (t *TaskService) TaskUpdate(ctx context.Context, taskInfo *task.TaskUpdateInput) error {
	//添加任务进行数据库检测
	if err := TestHost(ctx, taskInfo.HostID); err != nil {
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
		finNum, err := t.getTaskFinNum(ctx, databseres.Type, listIterm.Id)
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
			DeletedAt:   listIterm.DeletedAt.Time.Format("2006年01月02日15:04"),
			FinishNum:   finNum,
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

func (t *TaskService) getTaskFinNum(ctx context.Context, HostType, taskId int64) (int64, error) {
	switch HostType {
	case pkg.MysqlHost:
		mysqlHistoryDB := &dao.BakHistory{TaskID: taskId}
		list, err := mysqlHistoryDB.FindList(ctx, database.Gorm, mysqlHistoryDB)
		if err != nil {
			return 0, err
		}
		return int64(len(list)), nil
	case pkg.ElasticHost:
		elasticHistoryDB := &dao.ESHistoryDB{TaskID: taskId}
		list, err := elasticHistoryDB.FindList(ctx, database.Gorm, elasticHistoryDB)
		if err != nil {
			return 0, err
		}
		return int64(len(list)), nil
	}
	return 0, errors.New("类型不匹配")
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
		finNum, err := t.getTaskFinNum(ctx, databseres.Type, listIterm.Id)
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
			DeletedAt:   listIterm.DeletedAt.Time.Format("2006年01月02日15:04"),
			FinishNum:   finNum,
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

func (t *TaskService) GetDateNumInfo(ctx context.Context, info *task.DateNumInfoInput) (*task.DateNumInfoOut, error) {
	taskDB := &dao.TaskInfo{}
	list, err := taskDB.GetTaskByDate(ctx, database.Gorm, info.Date)
	if err != nil {
		return nil, err
	}
	mysqlHistoryDb := &dao.BakHistory{}
	mysqlList, err := mysqlHistoryDb.FindHistoryByDate(ctx, database.Gorm, info.Date)
	if err != nil {
		return nil, err
	}
	elasticDB := &dao.ESHistoryDB{}
	esList, err := elasticDB.FindHistoryByDate(ctx, database.Gorm, info.Date)
	allFinishNums := len(mysqlList) + len(esList)
	log.Logger.Infof("查询日期%s,查询数据任务数量%d,Mysql历史记录:%d,Elastic历史记录%d", info.Date, len(list), len(mysqlList), len(esList))
	return &task.DateNumInfoOut{
		Date:      info.Date,
		TaskNum:   int64(len(list)),
		FinishNum: int64(allFinishNums),
	}, nil
}

func (t *TaskService) TaskAutoCreate(ctx context.Context, info *task.TaskAutoCreateInPut) error {
	hostDB := &dao.HostDatabase{Id: info.HostID}
	host, err := hostDB.Find(ctx, database.Gorm, hostDB)
	if err != nil {
		return err
	}
	temp := strings.Split(host.Host, ":")
	hostAddr := temp[0]
	hostPort := temp[1]
	mysqlConInfo := database.NewMysqlConInfo(
		hostAddr,
		host.User,
		host.Password,
		hostPort,
		"mysql",
	)
	db, err := database.CreateDB(mysqlConInfo)
	if err != nil {
		return err
	}
	var databaseList []string
	var tempDatabaseList []string
	if err := db.Raw("show databases").Scan(&tempDatabaseList).Error; err != nil {
		return err
	}
	//处理数据，避免对默认库创建任务
	for _, dbName := range tempDatabaseList {
		if dbName != "mysql" && dbName != "sys" && dbName != "information_schema" && dbName != "performance_schema" {
			databaseList = append(databaseList, dbName)
		}
	}
	for _, dbName := range databaseList {
		input := &task.TaskAddInput{
			HostID:          info.HostID,
			DBName:          dbName,
			BackupCycle:     info.BackupCycle,
			KeepNumber:      info.KeepNumber,
			IsAllDBBak:      info.IsAllDBBak,
			IsDingSend:      info.IsDingSend,
			DingAccessToken: info.DingAccessToken,
			DingSecret:      info.DingSecret,
			OssType:         info.OssType,
			IsOssSave:       info.IsOssSave,
			Endpoint:        info.Endpoint,
			OssAccess:       info.OssAccess,
			OssSecret:       info.OssSecret,
			BucketName:      info.BucketName,
			Directory:       info.Directory,
			ServiceName:     config.GetStringConf("base", "serviceName"),
		}
		if err := t.TaskAdd(ctx, input); err != nil {
			return err
		}
	}
	return nil
}

//// TaskTest 在添加任务时，进行数据库连接测试，避免添加无用信息导致备份失败
//func TaskTest(ctx context.Context, hid int64, DBName, endpoint string) error {
//	hostDB := &dao.HostDatabase{Id: hid}
//	host, err := hostDB.Find(ctx, database.Gorm, hostDB)
//	if err != nil {
//		return err
//	}
//	log.Logger.Info("开始任务测试")
//	if err := HostPingCheck(host.User, host.Password, host.Host, DBName, 1); err != nil {
//		return err
//	}
//	if err := EndPointTest(endpoint); err != nil {
//		return err
//	}
//	return nil
//}
//
//// EndPointTest 进行endpoint的一个端口检测
//func EndPointTest(endpoint string) error {
//	if endpoint == "" {
//		return nil
//	}
//	if temp := strings.Split(endpoint, ":"); len(temp) != 2 {
//		return errors.New("您输入的Endpoint地址有误  正确地址: 192.168.1.1:9000")
//	}
//	_, err := net.DialTimeout("tcp", endpoint, 2*time.Second)
//	if err != nil {
//		return err
//	}
//	return nil
//}
