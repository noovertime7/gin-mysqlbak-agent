package server

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/clean"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/task"
	"context"
	"database/sql"
)

func Clean() error {
	log.Logger.Infof("开启数据清理定时任务")
	ctx := context.Background()
	tx := database.Gorm
	tsvc := service.TaskService{}
	taskDB := &dao.TaskInfo{IsDelete: sql.NullInt64{Int64: 0, Valid: true}}
	tasks, err := taskDB.FindList(ctx, tx, taskDB)
	if err != nil {
		return err
	}
	//1、查询当前任务，获取当前任务的keep_number
	for _, t := range tasks {
		keepTime := pkg.GetDateByKeepNumber(int(t.KeepNumber))
		log.Logger.Infof("当前任务ID%v,需要清理的日期为%s之前的数据", t.Id, keepTime)
		switch t.Type {
		case pkg.MysqlHost:
			//2、查询历史记录，查询大于keep_number的数据，得到文件名
			historyDB := &dao.BakHistory{TaskID: t.Id}
			historys, err := historyDB.FindListBeforDateTask(ctx, tx, keepTime)
			if err != nil {
				return err
			}
			if len(historys) == 0 {
				log.Logger.Infof("当前任务ID %v,需要清理的数量为0", t.Id)
				continue
			}
			log.Logger.Infof("任务ID%d当前需要清理的文件 %d", t.Id, len(historys))
			for _, history := range historys {
				log.Logger.Debugf("[Mysql]当前任务ID%d，超过保留周期%d，开始清理%s之前的sql文件，清理文件%s", t.Id, t.KeepNumber, keepTime, history.FileName)
				detail, err := tsvc.TaskDetail(ctx, &task.TaskIDInput{ID: t.Id})
				if err != nil {
					return err
				}
				OssInfo := &clean.OssInfo{
					IsOssSave:       detail.IsOssSave,
					OssType:         detail.OssType,
					Endpoint:        detail.Endpoint,
					AccessKey:       detail.OssAccess,
					SecretAccessKey: detail.OssSecret,
					BucketName:      detail.BucketName,
					Dir:             detail.Directory,
					Filepath:        history.FileName,
				}
				handler := clean.NewMysqlCleanHandler(history.FileName, history.Id, OssInfo)
				factory := clean.NewCleanerFactory(handler)
				if err := factory.Run(); err != nil {
					log.Logger.Errorf("数据清理失败,清理文件:%s，失败原因%v", history.FileName, err)
					return err
				}
				log.Logger.Infof("数据清理成功%v", history.FileName)
			}
		case pkg.ElasticHost:
			esHistoryDB := &dao.ESHistoryDB{TaskID: t.Id}
			historys, err := esHistoryDB.FindListBeforDateTask(ctx, tx, keepTime)
			if err != nil {
				return err
			}
			if len(historys) == 0 {
				log.Logger.Infof("当前任务ID %v,需要清理的数量为0", t.Id)
				continue
			}
			log.Logger.Infof("任务ID%d当前需要清理的文件 %d", t.Id, len(historys))
			for _, history := range historys {
				log.Logger.Debugf("[Elastic]当前任务ID%d，超过保留周期%d，开始清理%s之前的快照，清理快照%s", t.Id, t.KeepNumber, keepTime, history.Snapshot)
				hostDB := &dao.HostDatabase{Id: t.HostID}
				host, err := hostDB.Find(ctx, tx, hostDB)
				if err != nil {
					return err
				}
				handler := clean.NewElasticCleanHandler(host.Host, host.User, host.Password, history.Snapshot, history.Id)
				factory := clean.NewCleanerFactory(handler)
				if err := factory.Run(); err != nil {
					log.Logger.Errorf("数据清理失败,清理文件:%s，失败原因%v", history.Snapshot, err)
					return err
				}
				log.Logger.Infof("数据清理成功%v", history.Snapshot)
			}
		}
	}
	return nil
}
