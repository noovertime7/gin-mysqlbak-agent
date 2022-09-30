package dao

import (
	"backupAgent/proto/backupAgent/esbak"
	"context"
	"github.com/go-errors/errors"
	"gorm.io/gorm"
	"time"
)

type EsTaskDB struct {
	ID          int64     `json:"id" gorm:"primary_key" description:"自增主键"`
	ServiceName string    `json:"service_name" gorm:"column:service_name" description:"服务名"`
	HostID      int64     `json:"host_id"  gorm:"column:host_id" description:"主机id"`
	BackupCycle string    `json:"backup_cycle" gorm:"column:backup_cycle" description:"备份周期"`
	KeepNumber  int64     `json:"keep_number"  gorm:"column:keep_number" description:"数据保留周期"`
	IsDelete    int64     `json:"is_deleted" gorm:"column:is_deleted" description:"是否删除"`
	Status      int64     `json:"status" gorm:"column:status" description:"开关"`
	UpdatedAt   time.Time `json:"updated_at" gorm:"column:updated_at" description:"更新时间"`
	CreatedAt   time.Time `json:"created_at" gorm:"column:created_at" description:"添加时间"`
}

func (e *EsTaskDB) TableName() string {
	return "es_task"
}

func (e *EsTaskDB) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(e).Error
}

func (e *EsTaskDB) Updates(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(e.TableName()).Updates(e).Error
}

func (e *EsTaskDB) UpdateStatus(ctx context.Context, tx *gorm.DB, esTaskDB *EsTaskDB) error {
	if esTaskDB.ID == 0 {
		return errors.New("TaskID为空")
	}
	return tx.WithContext(ctx).Table(e.TableName()).Where("id = ?", esTaskDB.ID).Updates(map[string]interface{}{
		"status": esTaskDB.Status,
	}).Error
}

func (e *EsTaskDB) Find(ctx context.Context, tx *gorm.DB, search *EsTaskDB) (*EsTaskDB, error) {
	out := &EsTaskDB{}
	err := tx.WithContext(ctx).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (e *EsTaskDB) PageList(ctx context.Context, tx *gorm.DB, params *esbak.EsTaskListInput) ([]*EsTaskDB, int64, error) {
	var total int64 = 0
	var list []*EsTaskDB
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(ctx)
	query = query.Table(e.TableName()).Where("is_deleted=0")
	if params.Info != "" {
		query = query.Where("(host like ? or service_name like ? )", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	if err := query.Limit(int(params.PageSize)).Offset(int(offset)).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	query.Find(&list).Count(&total)
	return list, total, nil
}

func (e *EsTaskDB) TaskDetail(ctx context.Context, tx *gorm.DB, search *EsTaskDB) (*EsTaskDetail, error) {
	esInfo := &EsTaskDB{ID: search.ID}
	esinfoRes, err := esInfo.Find(ctx, tx, esInfo)
	if err != nil {
		return nil, err
	}
	hostDB := &HostDatabase{Id: esinfoRes.HostID}
	host, err := hostDB.Find(ctx, tx, hostDB)
	if err != nil {
		return nil, err
	}
	return &EsTaskDetail{ESTaskInfo: esinfoRes, HostInfo: host}, nil
}
