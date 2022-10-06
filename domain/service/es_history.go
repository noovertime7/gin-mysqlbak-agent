package service

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/proto/backupAgent/esbak"
	"context"
)

type esHistoryService struct{}

func NewESHistoryService() *esHistoryService {
	return &esHistoryService{}
}

func (e *esHistoryService) GetESHistoryList(ctx context.Context, esHistoryInfo *esbak.GetEsHistoryListInput) (*esbak.ESHistoryListOutput, error) {
	esHistoryDB := &dao.ESHistoryDB{}
	list, total, err := esHistoryDB.PageList(ctx, database.Gorm, esHistoryInfo)
	if err != nil {
		return nil, err
	}
	var OutList []*esbak.ESHistoryListOutItem
	for _, listItem := range list {
		esTaskDB := &dao.TaskInfo{Id: listItem.TaskID}
		taskinfo, err := esTaskDB.Find(ctx, database.Gorm, esTaskDB)
		if err != nil {
			return nil, err
		}
		hostDB := &dao.HostDatabase{Id: taskinfo.HostID}
		host, err := hostDB.Find(ctx, database.Gorm, hostDB)
		if err != nil {
			return nil, err
		}
		outItem := &esbak.ESHistoryListOutItem{
			ID:               listItem.Id,
			TaskID:           listItem.TaskID,
			Host:             host.Host,
			UUID:             listItem.UUID,
			DurationInMillis: listItem.DurationInMillis,
			Snapshot:         listItem.Snapshot,
			Repository:       listItem.Repository,
			Indices:          listItem.Indices,
			State:            listItem.State,
			StartTime:        listItem.StartTime.Format("2006年01月02日15:04:01:01"),
			EndTime:          listItem.EndTime.Format("2006年01月02日15:04:01:01"),
			Message:          listItem.Message,
			Status:           listItem.Status.Int64,
		}
		OutList = append(OutList, outItem)
	}
	return &esbak.ESHistoryListOutput{
		Total:                total,
		EsHistoryListOutItem: OutList,
		PageSize:             esHistoryInfo.PageSize,
		PageNo:               esHistoryInfo.PageNo,
	}, nil
}

func (e *esHistoryService) GetEsHistoryDetail(ctx context.Context, esHistoryInfo *esbak.ESHistoryIDInput) (*esbak.EsHistoryDetailOut, error) {
	//1、先查出history
	historyDB := &dao.ESHistoryDB{Id: esHistoryInfo.ID}
	history, err := historyDB.Find(ctx, database.Gorm, historyDB)
	if err != nil {
		return nil, err
	}
	//2、根据task_id查出task
	taskDB := &dao.TaskInfo{Id: history.TaskID}
	task, err := taskDB.Find(ctx, database.Gorm, taskDB)
	if err != nil {
		return nil, err
	}
	//3、根据task的host_id查出host
	hostDB := &dao.HostDatabase{Id: task.HostID}
	host, err := hostDB.Find(ctx, database.Gorm, hostDB)
	if err != nil {
		return nil, err
	}
	return &esbak.EsHistoryDetailOut{
		EsHostDetail: &esbak.EsHostDetail{
			HostID:   host.Id,
			Host:     host.Host,
			Status:   host.HostStatus,
			CreateAt: host.CreatedAt.Format("2006年01月02日15:04:01"),
			UpdateAt: host.UpdatedAt.Format("2006年01月02日15:04:01"),
		},
		ESTaskDetail: &esbak.ESTaskDetail{
			BackupCycle: task.BackupCycle,
			KeepNumber:  task.KeepNumber,
			Status:      task.Status,
			CreateAt:    task.CreatedAt.Format("2006年01月02日15:04:01"),
		},
		EsHistoryDetail: &esbak.ESHistoryListOutItem{
			ID:               history.Id,
			TaskID:           history.TaskID,
			UUID:             history.UUID,
			DurationInMillis: history.DurationInMillis,
			Host:             host.Host,
			Snapshot:         history.Snapshot,
			Repository:       history.Repository,
			Indices:          history.Indices,
			State:            history.State,
			StartTime:        history.StartTime.Format("2006年01月02日15:04:01"),
			EndTime:          history.EndTime.Format("2006年01月02日15:04:01"),
			Message:          history.Message,
			Status:           history.Status.Int64,
		},
	}, nil
}

func (e *esHistoryService) DeleteEsHistory(ctx context.Context, esHistoryInfo *esbak.ESHistoryIDInput) error {
	esHistoryDB := &dao.ESHistoryDB{Id: esHistoryInfo.ID}
	es, err := esHistoryDB.Find(ctx, database.Gorm, esHistoryDB)
	if err != nil {
		return err
	}
	es.IsDeleted = 1
	return es.Updates(ctx, database.Gorm)
}

func (e *esHistoryService) GetEsHistoryNumInfo(ctx context.Context) (*esbak.EsHistoryNumInfoOut, error) {
	var (
		weekNums int64
		failNums int64
	)
	info := &esbak.GetEsHistoryListInput{
		Info:      "",
		PageNo:    1,
		PageSize:  pkg.LargePageSize,
		SortField: "",
		SortOrder: "",
		Status:    pkg.HistoryStatusAll,
	}
	list, err := e.GetESHistoryList(ctx, info)
	if err != nil {
		return nil, err
	}
	//查询7天内任务数量
	esHistoryDB := &dao.ESHistoryDB{}
	data, err := esHistoryDB.FindByDate(ctx, database.Gorm, 7)
	if err != nil {
		return nil, err
	}
	weekNums = int64(len(data))
	for _, h := range list.EsHistoryListOutItem {
		//统计快照失败的数量
		if h.Status != 1 {
			failNums++
		}
	}
	return &esbak.EsHistoryNumInfoOut{
		AllNums:  list.Total,
		WeekNums: weekNums,
		FailNums: failNums,
	}, nil
}
