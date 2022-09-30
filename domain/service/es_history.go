package service

import (
	"backupAgent/domain/dao"
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
		esTaskDB := &dao.EsTaskDB{ID: listItem.TaskID}
		taskinfo, err := esTaskDB.Find(ctx, database.Gorm, esTaskDB)
		if err != nil {
			return nil, err
		}
		outItem := &esbak.ESHistoryListOutItem{
			ID:               listItem.Id,
			TaskID:           listItem.TaskID,
			Host:             taskinfo.Host,
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

func (e *esHistoryService) DeleteEsHistory(ctx context.Context, esHistoryInfo *esbak.ESHistoryIDInput) error {
	esHistoryDB := &dao.ESHistoryDB{Id: esHistoryInfo.ID}
	es, err := esHistoryDB.Find(ctx, database.Gorm, esHistoryDB)
	if err != nil {
		return err
	}
	es.IsDeleted = 1
	return es.Updates(ctx, database.Gorm)
}
