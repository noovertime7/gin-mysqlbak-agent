package main

import (
	"backupAgent/domain/config"
	"backupAgent/domain/pkg/database"
	tracer "backupAgent/domain/pkg/trace"
	"backupAgent/domain/server"
	"backupAgent/handler"
	"backupAgent/proto/backupAgent/bak"
	"backupAgent/proto/backupAgent/bakhistory"
	"backupAgent/proto/backupAgent/host"
	"backupAgent/proto/backupAgent/task"
	"github.com/micro/go-micro/v2"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/wrapper/trace/opentracing/v2"
	"time"
)

var (
	ServiceName = config.GetStringConf("base", "serviceName")
	Address     = config.GetStringConf("base", "addr")
	JaegerAddr  = config.GetStringConf("jaeger", "addr")
)

func main() {
	if err := server.InitResourceAndStart(); err != nil {
		log.Fatal("初始化失败", err)
	}
	jaegerTracer, closer, err := tracer.NewJaegerTracer(ServiceName, JaegerAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	// New microService
	if config.GetBoolConf("jaeger", "enable") {
		microService := micro.NewService(
			micro.Name(config.GetStringConf("base", "serviceName")),
			micro.Version(config.GetStringConf("base", "version")),
			//micro.Address(config.GetStringConf("base", "addr")),
			micro.Address(Address),
			micro.RegisterTTL(30*time.Second),
			micro.RegisterInterval(15*time.Second),
			//micro.Registry(reg),
			micro.WrapHandler(opentracing.NewHandlerWrapper(jaegerTracer)),
			micro.AfterStop(func() error {
				data, err := server.Reg.DeRegister()
				if err != nil {
					log.Fatal("注销失败", err)
					return err
				}
				log.Info(data)
				return nil
			}),
		)
		// Initialise microService
		microService.Init()
		database.InitDB()
		// Register Handler
		_ = host.RegisterHostHandler(microService.Server(), new(handler.HostHandler))
		_ = task.RegisterTaskHandler(microService.Server(), new(handler.TaskHandler))
		_ = bakhistory.RegisterHistoryHandler(microService.Server(), new(handler.HistoryHandler))
		_ = bak.RegisterBakServiceHandler(microService.Server(), new(handler.BakHandler))
		// Run microService
		if err := microService.Run(); err != nil {
			log.Fatal(err)
		}
	}
	//reg.Init()
	microService := micro.NewService(
		micro.Name(config.GetStringConf("base", "serviceName")),
		micro.Version(config.GetStringConf("base", "version")),
		micro.Address(Address),
		micro.RegisterTTL(30*time.Second),
		micro.RegisterInterval(15*time.Second),
		micro.AfterStop(func() error {
			data, err := server.Reg.DeRegister()
			if err != nil {
				log.Fatal("注销失败", err)
				return err
			}
			log.Info(data)
			return nil
		}),
		//micro.Registry(reg),
	)
	microService.Options()
	// Initialise microService
	microService.Init()
	// 启动所有备份任务，避免程序中止后，任务停止
	// Register Handler
	_ = host.RegisterHostHandler(microService.Server(), new(handler.HostHandler))
	_ = task.RegisterTaskHandler(microService.Server(), new(handler.TaskHandler))
	_ = bakhistory.RegisterHistoryHandler(microService.Server(), new(handler.HistoryHandler))
	_ = bak.RegisterBakServiceHandler(microService.Server(), new(handler.BakHandler))
	// Run microService
	if err := microService.Run(); err != nil {
		log.Fatal(err)
	}
}
