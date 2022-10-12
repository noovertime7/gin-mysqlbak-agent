package server

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"context"
	"database/sql"
)

func Clean() error {
	log.Logger.Infof("开启数据清理定时任务")
	ctx := context.Background()
	tx := database.Gorm
	taskDB := &dao.TaskInfo{IsDelete: sql.NullInt64{Int64: 0, Valid: true}}
	tasks, err := taskDB.FindList(ctx, tx, taskDB)
	if err != nil {
		return err
	}
	//1、查询当前任务，获取当前任务的keep_number
	for _, t := range tasks {
		keepTime := pkg.GetDateByKeepNumber(int(t.KeepNumber))
		log.Logger.Infof("当前任务ID%v,需要清理的日期为%s之前的数据", t.Id, keepTime)
		//2、查询历史记录，查询大于keep_number的数据，得到文件名
	}
	return nil
}

type Cleaner struct {
}

//3、判断任务是否启用oss，如果启用，调用oss删除方法

//4、判断是否启用钉钉，如果启用钉钉发送删除成功消息
