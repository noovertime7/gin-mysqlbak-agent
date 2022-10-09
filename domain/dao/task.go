package dao

import (
	"backupAgent/proto/backupAgent/task"
	"context"
	"database/sql"
	"errors"
	"gorm.io/gorm"
	"time"
)

type TaskInfo struct {
	Id          int64         `gorm:"primary_key" description:"自增主键"`
	HostID      int64         `json:"id" gorm:"column:host_id" description:"主机关系id"`
	ServiceName string        `json:"service_name" gorm:"column:service_name" description:"服务名"`
	DBName      string        `json:"db_name" gorm:"column:db_name" description:"备份库名"`
	BackupCycle string        `json:"backup_cycle" gorm:"column:backup_cycle" description:"备份周期"`
	KeepNumber  int64         `json:"keep_number"  gorm:"column:keep_number" description:"数据保留周期"`
	IsAllDBBak  int64         `json:"is_all_dbbak" gorm:"column:is_all_dbbak" description:"是否全库备份"`
	IsDelete    sql.NullInt64 `json:"is_deleted" gorm:"column:is_deleted" description:"是否删除"`
	Status      int64         `json:"status" gorm:"column:status" description:"开关"`
	UpdatedAt   time.Time     `json:"updated_at" gorm:"column:updated_at" description:"更新时间"`
	CreatedAt   time.Time     `json:"created_at" gorm:"column:created_at" description:"添加时间"`
	DeletedAt   time.Time     `json:"deleted_at" gorm:"column:deleted_at" description:"删除时间"`
}

func (t *TaskInfo) TableName() string {
	return "t_taskinfo"
}

func (t *TaskInfo) Save(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Save(t).Error
}

func (t *TaskInfo) Delete(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Delete(t).Error
}

func (t *TaskInfo) Updates(ctx context.Context, tx *gorm.DB) error {
	return tx.WithContext(ctx).Table(t.TableName()).Updates(t).Error
}

func (t *TaskInfo) UpdateStatus(ctx context.Context, tx *gorm.DB, taskDB *TaskInfo) error {
	if taskDB.Id == 0 {
		return errors.New("TaskID为空")
	}
	return tx.WithContext(ctx).Table(t.TableName()).Where("id = ?", taskDB.Id).Updates(map[string]interface{}{
		"status": taskDB.Status,
	}).Error
}

func (t *TaskInfo) Find(ctx context.Context, tx *gorm.DB, search *TaskInfo) (*TaskInfo, error) {
	out := &TaskInfo{}
	err := tx.WithContext(ctx).Where(search).Find(out).Error
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (t *TaskInfo) PageList(ctx context.Context, tx *gorm.DB, params *task.TaskListInput) ([]*TaskInfo, int64, error) {
	var total int64 = 0
	var list []*TaskInfo
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(ctx)
	query = query.Table(t.TableName()).Where("is_deleted=0 and host_id = ?", params.HostID)
	query.Find(&list).Count(&total)
	if params.Info != "" {
		query = query.Where("(db_name like ? or service_name like ? )", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	if err := query.Limit(int(params.PageSize)).Offset(int(offset)).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, total, nil
}

// UnscopedPageList 分页查询所有task，主要用于同步删除状态
func (t *TaskInfo) UnscopedPageList(ctx context.Context, tx *gorm.DB, params *task.TaskListInput) ([]*TaskInfo, int64, error) {
	var total int64 = 0
	var list []*TaskInfo
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(ctx)
	query = query.Table(t.TableName()).Where(" host_id = ?", params.HostID)
	query.Find(&list).Count(&total)
	if params.Info != "" {
		query = query.Where("(db_name like ? or service_name like ? )", "%"+params.Info+"%", "%"+params.Info+"%")
	}
	if err := query.Limit(int(params.PageSize)).Offset(int(offset)).Order("id desc").Find(&list).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0, err
	}
	return list, total, nil
}

func (t *TaskInfo) FindAllTaskByHost(ctx context.Context, tx *gorm.DB, search *TaskInfo) ([]*TaskInfo, error) {
	var result []*TaskInfo
	err := tx.WithContext(ctx).Where("is_deleted = 0 and host_id = ?", search.HostID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *TaskInfo) FindStatusUPTaskByHost(ctx context.Context, tx *gorm.DB, search *TaskInfo) ([]*TaskInfo, error) {
	var result []*TaskInfo
	err := tx.WithContext(ctx).Where("is_deleted = 0 and host_id = ? and status = 1", search.HostID).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (t *TaskInfo) TaskDetail(ctx context.Context, tx *gorm.DB, search *TaskInfo) (*TaskDetail, error) {
	info := &TaskInfo{Id: search.Id}
	infores, err := info.Find(ctx, tx, info)
	if err != nil {
		return nil, err
	}
	hostinfo := &HostDatabase{Id: infores.HostID, Type: 1}
	//查询mysql数据，type =1
	hostinfores, err := hostinfo.Find(ctx, tx, hostinfo)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	ding := &DingDatabase{TaskID: search.Id}
	dingres, err := ding.Find(ctx, tx, ding)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	oss := &OssDatabase{TaskID: search.Id}
	ossres, err := oss.Find(ctx, tx, oss)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return &TaskDetail{
		Host: hostinfores,
		Info: infores,
		Oss:  ossres,
		Ding: dingres,
	}, nil
}
