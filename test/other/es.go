package main

import (
	"backupAgent/domain/core"
	"backupAgent/domain/dao"
	"log"
	"time"
)

func main() {

	handler, err := core.NewEsBakHandler(&dao.EsTaskDetail{ESTaskInfo: &dao.EsTaskDB{
		Host:          "http://10.20.110.51:39200",
		Username:      "elastic",
		Password:      "Tsit@123",
		BackupCycle:   "* * * * *",
		KeepNumber:    7,
		Index:         "info",
		IsAllIndexBak: 0,
		IsDelete:      0,
		Status:        0,
	}})
	if err != nil {
		log.Fatalln(err)
	}
	if err := handler.Start(); err != nil {
		log.Fatalln(err)
	}
	time.Sleep(80 * time.Second)
	if err := handler.Stop(); err != nil {
		log.Fatalln(err)
	}
	select {}
}
