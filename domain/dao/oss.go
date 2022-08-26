package dao

import (
	"context"
	"gorm.io/gorm"
)

type OssDatabase struct {
	Id         int64  `gorm:"primary_key" description:"自增主键"`
	TaskID     int64  `json:"task_id" gorm:"column:task_id" description:"任务id"`
	IsOssSave  int64  `json:"is_oss_save" gorm:"column:is_oss_save" description:"是否保存到oss中 0关闭1开启"`
	OssType    int64  `json:"oss_type" gorm:"column:oss_type" description:"oss类型"`
	Endpoint   string `json:"endpoint"  gorm:"column:endpoint" description:"endpoint"`
	OssAccess  string `json:"oss_access"  gorm:"column:oss_access" description:"ossaccess"`
	OssSecret  string `json:"oss_secret"  gorm:"column:oss_secret" description:"secret"`
	BucketName string `json:"bucket_name"  gorm:"column:bucket_name" description:"bucket名字"`
	Directory  string `json:"directory" gorm:"column:directory" description:"目录"`
}

func (o *OssDatabase) TableName() string {
	return "t_oss"
}

func (o *OssDatabase) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(o).Error
}

func (o *OssDatabase) Find(ctx context.Context, tx *gorm.DB, search *OssDatabase) (*OssDatabase, error) {
	out := &OssDatabase{}
	err := tx.WithContext(ctx).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (o *OssDatabase) UpdatesByMap(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(o.TableName()).Where("task_id = ?", o.TaskID).Updates(map[string]interface{}{
		"is_oss_save": o.IsOssSave,
		"oss_type":    o.OssType,
		"endpoint":    o.Endpoint,
		"oss_access":  o.OssAccess,
		"oss_secret":  o.OssSecret,
		"bucket_name": o.BucketName,
		"directory":   o.Directory,
	}).Error
}
