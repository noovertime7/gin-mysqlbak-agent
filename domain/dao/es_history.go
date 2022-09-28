package dao

import (
	"backupAgent/domain/pkg"
	"backupAgent/proto/backupAgent/esbak"
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type ESHistoryDB struct {
	Id                int64     `gorm:"primary_key"  comment:"自增主键"`
	TaskID            int64     `json:"task_id" gorm:"column:task_id;index:index_task_id"  comment:"任务id"`
	Snapshot          string    `json:"snapshot"  gorm:"column:snapshot;index:index_snapshot"  comment:"快照名字"`
	Repository        string    `json:"repository" gorm:"column:repository"  comment:"仓库名"`
	UUID              string    `json:"uuid"  gorm:"column:uuid"  comment:"UUID"`
	Version           string    `json:"version"  gorm:"column:version"  comment:"版本"`
	Indices           string    `json:"indices"  gorm:"column:indices;type:text"  comment:"包含索引"`
	State             string    `json:"state"  gorm:"column:state"  comment:"状态"`
	StartTime         time.Time `json:"start_time"  gorm:"column:start_time"  comment:"开始时间"`
	StartTimeInMillis int64     `json:"start_time_in_millis"  gorm:"column:start_time_in_millis"  comment:"start_time_in_millis"`
	EndTime           time.Time `json:"end_time"  gorm:"column:end_time"  comment:"结束时间"`
	EndTimeInMillis   int64     `json:"end_time_in_millis"  gorm:"column:end_time_in_millis"  comment:"end_time_in_millis"`
	DurationInMillis  int64     `json:"duration_in_millis"  gorm:"column:duration_in_millis"  comment:"消耗时间"`
	Message           string    `json:"message"  gorm:"column:message"  comment:"备注"`
	IsDeleted         int64     `json:"is_deleted" gorm:"column:is_deleted;type:int(12);default:0"  comment:"软删除标记"`
}

func (e *ESHistoryDB) TableName() string {
	return "es_bak_history"
}

func (e *ESHistoryDB) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(e).Error
}

func (e *ESHistoryDB) Find(c context.Context, tx *gorm.DB, search *ESHistoryDB) (*ESHistoryDB, error) {
	out := &ESHistoryDB{}
	err := tx.WithContext(c).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (e *ESHistoryDB) Updates(ctx context.Context, tx *gorm.DB) error {
	if e.Id == 0 {
		return errors.New("ID 为空")
	}
	return tx.WithContext(ctx).Table(e.TableName()).Where("id = ?", e.Id).Updates(e).Error
}

func (e *ESHistoryDB) PageList(c context.Context, tx *gorm.DB, params *esbak.GetEsHistoryListInput) ([]ESHistoryDB, int64, error) {
	var total int64 = 0
	var list []ESHistoryDB
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(e.TableName()).Where("is_deleted=0")
	query.Find(&list).Count(&total)
	if params.Status != "" {
		switch params.Status {
		case pkg.HistoryStatusAll:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where(fmt.Sprintf(" snapshot like '%%%s' or indices like'%%%s'  ", searchInfo, searchInfo))
			} else {
				query = query.Table(e.TableName()).Where("is_deleted = 0")
			}
		case pkg.HistoryStatusSuccess:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where("(snapshot like ? or indices like ?)", searchInfo, searchInfo)
			} else {
				query = query.Where("message = 'success' ")
			}
		case pkg.HistoryStatusFail:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where("snapshot like ? or indices like ?", searchInfo, searchInfo)
			} else {
				query = query.Where("message != 'success' ")
			}
		}
	}
	var sortRules string
	switch params.SortOrder {
	case "descend":
		sortRules = "desc"
	case "ascend":
		sortRules = "asc"
	}
	if params.SortField == "" {
		params.SortField = "id"
		sortRules = "desc"
	}
	if err := query.Limit(int(params.PageSize)).Offset(int(offset)).Order(fmt.Sprintf("%s %s", params.SortField, sortRules)).Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}

	return list, total, nil
}
