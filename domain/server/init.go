package server

import (
	"backupAgent/domain/config"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/bakhistory"
	"backupAgent/proto/backupAgent/host"
	"backupAgent/proto/backupAgent/task"
	"context"
	"time"
)

var (
	ServiceName     = config.GetStringConf("base", "serviceName")
	Address         = config.GetStringConf("base", "addr")
	RegisterAddress = config.GetStringConf("register", "registerAddr")
	Content         = config.GetStringConf("base", "content")
)

var Reg pkg.AgentRegister

func InitResourceAndStart() error {
	//打印logo信息
	pkg.PrintLogo()
	//初始化数据库
	database.InitDB()
	Clean()
	//注册服务
	go LoopRegister()
	//启动状态为1的任务
	if err := service.StartAllBakTask(context.Background()); err != nil {
		log.Logger.Warning("程序启动，开启备份任务失败")
	}
	log.Logger.Info("程序启动开启所有备份任务成功")
	return nil
}

func GetTaskNum(ctx context.Context) (int, error) {
	s := &service.HostService{}
	taskService := &service.TaskService{}
	var taskNum int
	hostList, err := s.GetHostList(context.Background(), &host.HostListInput{
		Info:     "",
		PageNo:   1,
		PageSize: 999,
	})
	if err != nil {
		return 0, err
	}
	for _, hostInfo := range hostList.ListItem {
		taskListOut, err := taskService.TaskList(ctx, &task.TaskListInput{
			HostID:   hostInfo.ID,
			Info:     "",
			PageNo:   1,
			PageSize: pkg.LargePageSize,
		})
		if err != nil {
			return 0, err
		}
		taskNum += int(taskListOut.Total)
	}
	return taskNum, nil
}

func GetFinishNum(ctx context.Context) (int, error) {
	h := &service.HistoryService{}
	historyInfo, err := h.GetHistoryList(ctx, &bakhistory.HistoryListInput{
		Info:      "",
		PageNo:    1,
		PageSize:  pkg.LargePageSize,
		SortOrder: "",
	})
	if err != nil {
		return 0, err
	}
	return int(historyInfo.Total), nil
}

func LoopRegister() {
	var registrationCycle = config.GetIntConf("register", "registrationCycle")
	for {
		log.Logger.Info("开启定时注册任务")
		taskNum, err := GetTaskNum(context.Background())
		if err != nil {
			return
		}
		finishNum, err := GetFinishNum(context.Background())
		if err != nil {
			return
		}
		Reg = pkg.NewAgentRegister(ServiceName, RegisterAddress, Content, taskNum, finishNum)
		data, err := Reg.Register()
		if err != nil {
			log.Logger.Warning("注册失败", err.Error())
			time.Sleep(time.Duration(registrationCycle) * time.Minute)
			continue
		}
		time.Sleep(time.Duration(registrationCycle) * time.Minute)
		log.Logger.Info(data)
	}
}
