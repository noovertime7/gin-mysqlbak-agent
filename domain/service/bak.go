package service

import (
	"backupAgent/domain/core"
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/proto/backupAgent/bak"
	"backupAgent/proto/backupAgent/host"
	"context"
	"errors"
	"time"
)

type BakService struct{}

func (b *BakService) StartBak(ctx context.Context, bakInfo *bak.StartBakInput, isTest bool) error {
	log.Logger.Debug("Service层开始启动备份任务，入参:", bakInfo)
	taskDB := &dao.TaskInfo{Id: bakInfo.TaskID}
	taskDetail, err := taskDB.TaskDetail(ctx, database.Gorm, taskDB)
	if err != nil {
		return err
	}
	if taskDetail.Info.IsDelete.Int64 == 1 {
		return errors.New("任务已被删除,启动失败")
	}
	if isTest {
		log.Logger.Info("当前启动为测试任务，将备份周期修改为* * * * *")
		taskDetail.Info.BackupCycle = "* * * * *"
	}
	bakHandler, err := core.NewBakHandler(taskDetail)
	if err != nil {
		return err
	}
	if err := bakHandler.StartBak(); err != nil {
		return err
	}
	taskDB.Status = 1
	log.Logger.Debug("Service层修改bak任务状态", taskDB)
	return taskDB.UpdateStatus(ctx, database.Gorm, taskDB)
}

func (b *BakService) TestBak(ctx context.Context, bakInfo *bak.StartBakInput) error {
	if err := b.StartBak(ctx, bakInfo, true); err != nil {
		return err
	}
	go func() {
		time.Sleep(1 * time.Minute)
		newCtx := context.TODO()
		if err := b.StopBak(newCtx, &bak.StopBakInput{
			TaskID:      bakInfo.TaskID,
			ServiceName: bakInfo.ServiceName,
		}); err != nil {
			log.Logger.Error("测试任务停止失败", err)
			return
		}
		return
	}()
	return nil
}

func (b *BakService) StopBak(ctx context.Context, bakInfo *bak.StopBakInput) error {
	log.Logger.Debug("Service层开始停止备份任务，入参:", bakInfo)
	bakHandler := &core.Handler{}
	if err := bakHandler.StopBak(bakInfo.TaskID); err != nil {
		return err
	}
	log.Logger.Debug("Service层停止备份任务成功，入参:", bakInfo)
	taskDB := &dao.TaskInfo{Id: bakInfo.TaskID}
	taskDB.Status = 0
	log.Logger.Debug("Service层修改bak任务状态", taskDB)
	return taskDB.UpdateStatus(ctx, database.Gorm, taskDB)
}

func (b *BakService) StartBakByHost(ctx context.Context, bakInfo *bak.StartBakByHostInput) error {
	log.Logger.Debug("Service层开始启动主机所有备份任务，入参:", bakInfo)
	taskDB := &dao.TaskInfo{}
	tasks, err := taskDB.FindAllTaskByHost(ctx, database.Gorm, &dao.TaskInfo{HostID: bakInfo.HostID, Status: 0})
	if err != nil {
		return err
	}
	for _, task := range tasks {
		// 因为存在程序停止，状态为1的情况，所以这里不做判断
		log.Logger.Infof("Service启动主机所有备份任务，任务%d-%s:", task.Id, task.DBName)
		taskDB := &dao.TaskInfo{Id: task.Id}
		taskDetail, err := taskDB.TaskDetail(ctx, database.Gorm, taskDB)
		if err != nil {
			return err
		}
		bakHandler, err := core.NewBakHandler(taskDetail)
		if err != nil {
			return err
		}
		if err := bakHandler.StartBak(); err != nil {
			return err
		}
		taskDB.Status = 1
		log.Logger.Debug("Service层修改bak任务状态", taskDB)
		if err := taskDB.UpdateStatus(ctx, database.Gorm, taskDB); err != nil {
			return err
		}
	}
	return err
}

func (b *BakService) StopBakByHost(ctx context.Context, bakInfo *bak.StopBakByHostInput) error {
	log.Logger.Debug("Service层开始停止主机所有备份任务，入参:", bakInfo)
	taskDB := &dao.TaskInfo{HostID: bakInfo.HostID, Status: 1}
	tasks, err := taskDB.FindAllTaskByHost(ctx, database.Gorm, taskDB)
	if err != nil {
		return err
	}
	bakHandler := &core.Handler{}
	for _, task := range tasks {
		//只对状态为启动的任务关闭
		if task.Status == 1 {
			bakDB := &dao.TaskInfo{Id: task.Id, Status: 0}
			log.Logger.Infof("Service层停止主机所有备份任务，任务%d-%s:", task.Id, task.DBName)
			if err := bakHandler.StopBak(task.Id); err != nil {
				return err
			}
			log.Logger.Debug("Service层修改bak任务状态", bakDB)
			if err := bakDB.UpdateStatus(ctx, database.Gorm, bakDB); err != nil {
				return err
			}
			log.Logger.Debug("Service层修改bak任务状态成功", err)
		}
	}
	return nil
}

func StartAllBakTask(ctx context.Context) error {
	log.Logger.Info("Agent启动，开启已启动备份任务")
	s := &HostService{}
	hostList, err := s.GetHostList(context.Background(), &host.HostListInput{
		Info:     "",
		PageNo:   1,
		PageSize: 999,
	})
	if err != nil {
		return err
	}
	for _, hostInfo := range hostList.ListItem {
		taskDB := &dao.TaskInfo{}
		tasks, err := taskDB.FindStatusUPTaskByHost(ctx, database.Gorm, &dao.TaskInfo{HostID: hostInfo.ID})
		if err != nil {
			return err
		}
		for _, task := range tasks {
			// 因为存在程序停止，状态为1的情况，所以这里不做判断
			log.Logger.Infof("Service启动主机所有备份任务，任务%d-%s:", task.Id, task.DBName)
			taskDB := &dao.TaskInfo{Id: task.Id}
			taskDetail, err := taskDB.TaskDetail(ctx, database.Gorm, taskDB)
			if err != nil {
				return err
			}
			bakHandler, err := core.NewBakHandler(taskDetail)
			if err != nil {
				return err
			}
			if err := bakHandler.StartBak(); err != nil {
				return err
			}
			taskDB.Status = 1
			log.Logger.Debug("Service层修改bak任务状态", taskDB)
			if err := taskDB.UpdateStatus(ctx, database.Gorm, taskDB); err != nil {
				return err
			}
		}
	}
	return nil
}
