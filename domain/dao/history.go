package dao

import (
	"backupAgent/domain/pkg"
	"backupAgent/proto/backupAgent/bakhistory"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type BakHistory struct {
	Id         int64     `gorm:"primary_key" description:"自增主键"`
	TaskID     int64     `gorm:"column:task_id" description:"任务id"`
	Host       string    `gorm:"column:host" description:"主机"`
	DBName     string    `gorm:"column:db_name" description:"库名"`
	OssStatus  int64     `gorm:"column:oss_status"  description:"钉钉发送状态"`
	DingStatus int64     `gorm:"column:ding_status"  description:"OSS保存状态"`
	BakStatus  int64     `gorm:"column:bak_status" description:"备份状态"`
	Msg        string    `gorm:"column:message" description:"消息"`
	FileSize   int64     `gorm:"column:file_size" description:"文件大小"`
	FileName   string    `gorm:"column:filename" description:"文件名"`
	BakTime    time.Time `gorm:"column:bak_time" description:"备份时间"`
	IsDeleted  int64     `json:"is_deleted" gorm:"column:is_deleted"`
}

func (b *BakHistory) TableName() string {
	return "bak_history"
}

func (b *BakHistory) Save(c context.Context, tx *gorm.DB) error {
	return tx.WithContext(c).Save(b).Error
}

func (b *BakHistory) Find(c context.Context, tx *gorm.DB, search *BakHistory) (*BakHistory, error) {
	out := &BakHistory{}
	return out, tx.WithContext(c).Where(search).Find(out).Error
}

func (b *BakHistory) FindList(ctx context.Context, tx *gorm.DB, search *BakHistory) ([]*BakHistory, error) {
	var out []*BakHistory
	return out, tx.WithContext(ctx).Where(&search).Find(&out).Error
}

func (b *BakHistory) PageList(c context.Context, tx *gorm.DB, params *bakhistory.HistoryListInput) ([]BakHistory, int64, error) {
	var total int64 = 0
	var list []BakHistory
	offset := (params.PageNo - 1) * params.PageSize
	query := tx.WithContext(c)
	query = query.Table(b.TableName()).Where("is_deleted=0")
	query.Find(&list).Count(&total)
	if params.Status != "" {
		switch params.Status {
		case pkg.HistoryStatusAll:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where(fmt.Sprintf(" host like '%%%s' or db_name like'%%%s'  ", searchInfo, searchInfo))
			} else {
				query = query.Table(b.TableName()).Where("is_deleted = 0")
			}
		case pkg.HistoryStatusSuccess:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where("(host like ? or db_name like ?)", searchInfo, searchInfo)
			} else {
				query = query.Where("message = 'success' ")
			}
		case pkg.HistoryStatusFail:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where("host like ? or db_name like ?", searchInfo, searchInfo)
			} else {
				query = query.Where("message != 'success' ")
			}
		default:
			if params.Info != "" {
				searchInfo := "%" + params.Info
				query = query.Where(fmt.Sprintf(" host like '%%%s' or db_name like'%%%s'  ", searchInfo, searchInfo))
			} else {
				query = query.Table(b.TableName()).Where("is_deleted = 0")
			}
		}
	}
	var sortRules string
	switch params.SortOrder {
	case "descend":
		sortRules = "desc"
	case "ascend":
		sortRules = "asc"
	default:
		sortRules = "desc"
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

// FindByDate 查询7天内数据
func (b *BakHistory) FindByDate(ctx context.Context, tx *gorm.DB, num int) ([]BakHistory, error) {
	var out []BakHistory
	return out, tx.WithContext(ctx).Raw("SELECT * FROM bak_history WHERE is_deleted !=1 and date_sub(curdate(), interval ? day) <= date(bak_time);", num).Scan(&out).Error
}
