package dao

import (
	"context"
	"gorm.io/gorm"
	"time"
)

type ESHistoryDB struct {
	Id                int64     `gorm:"primary_key" description:"自增主键"`
	TaskID            int64     `json:"task_id" gorm:"column:task_id;index:task_id" description:"任务id"`
	Snapshot          string    `json:"snapshot"  gorm:"column:snapshot" description:"快照名字"`
	Repository        string    `json:"repository" gorm:"column:repository" description:"仓库名"`
	UUID              string    `json:"uuid"  gorm:"column:uuid" description:"UUID"`
	Version           string    `json:"version"  gorm:"column:version" description:"版本"`
	Indices           string    `json:"indices"  gorm:"column:indices;type:text" description:"包含索引"`
	State             string    `json:"state"  gorm:"column:state" description:"状态"`
	StartTime         time.Time `json:"start_time"  gorm:"column:start_time" description:"开始时间"`
	StartTimeInMillis int64     `json:"start_time_in_millis"  gorm:"column:start_time_in_millis" description:"start_time_in_millis"`
	EndTime           time.Time `json:"end_time"  gorm:"column:end_time" description:"结束时间"`
	EndTimeInMillis   int64     `json:"end_time_in_millis"  gorm:"column:end_time_in_millis" description:"end_time_in_millis"`
	DurationInMillis  int64     `json:"duration_in_millis"  gorm:"column:duration_in_millis" description:"消耗时间"`
	Message           string    `json:"message"  gorm:"column:message" description:"备注"`
}

func (e *ESHistoryDB) TableName() string {
	return "es_bak_history"
}

func (e *ESHistoryDB) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(e).Error
}
