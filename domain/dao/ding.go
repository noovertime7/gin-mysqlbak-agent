package dao

import (
	"context"
	"gorm.io/gorm"
)

type DingDatabase struct {
	Id              int64  `gorm:"primary_key" description:"自增主键"`
	TaskID          int64  `json:"task_id" gorm:"column:task_id" description:"任务id"`
	IsDingSend      int64  `json:"is_ding_send"  gorm:"column:is_ding_send" description:"是否发送钉钉消息"`
	DingAccessToken string `json:"ding_access_token"  gorm:"column:ding_access_token" description:"accessToken"`
	DingSecret      string `json:"ding_secret" gorm:"column:ding_secret" description:"secret"`
}

func (d *DingDatabase) TableName() string {
	return "t_ding"
}

func (d *DingDatabase) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(d).Error
}

func (d *DingDatabase) UpdatesByMap(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(d.TableName()).Where("task_id = ?", d.TaskID).Updates(map[string]interface{}{
		"is_ding_send":      d.IsDingSend,
		"ding_access_token": d.DingAccessToken,
		"ding_secret":       d.DingSecret,
	}).Error
}

func (d *DingDatabase) Find(ctx context.Context, tx *gorm.DB, search *DingDatabase) (*DingDatabase, error) {
	out := &DingDatabase{}
	err := tx.WithContext(ctx).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}
