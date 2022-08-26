package service

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg/database"
	"backupAgent/proto/backupAgent/bakhistory"
	"context"
	"strconv"
)

type HistoryService struct{}

func (h *HistoryService) GetHistoryList(ctx context.Context, historyInfo *bakhistory.HistoryListInput) (*bakhistory.HistoryListOutput, error) {
	historyDB := &dao.BakHistory{}
	list, total, err := historyDB.PageList(ctx, database.Gorm, historyInfo)
	if err != nil {
		return nil, err
	}
	var outList []*bakhistory.HistoryListOutItem
	for _, listIterm := range list {
		outIterm := &bakhistory.HistoryListOutItem{
			ID:         listIterm.Id,
			Host:       listIterm.Host,
			DBName:     listIterm.DBName,
			DingStatus: listIterm.DingStatus,
			OSSStatus:  listIterm.OssStatus,
			Message:    listIterm.Msg,
			FileSize:   strconv.Itoa(int(listIterm.FileSize)),
			FileName:   listIterm.FileName,
			BakTime:    listIterm.BakTime.Format("2006年01月02日15:04:01"),
		}
		outList = append(outList, outIterm)
	}
	return &bakhistory.HistoryListOutput{
		Total:              total,
		HistoryListOutItem: outList,
	}, nil
}

func (h *HistoryService) DeleteHistory(ctx context.Context, historyInfo *bakhistory.HistoryIDInput) error {
	history := &dao.BakHistory{Id: historyInfo.ID}
	historyDB, err := history.Find(ctx, database.Gorm, history)
	if err != nil {
		return err
	}
	historyDB.IsDeleted = 1
	return historyDB.Save(ctx, database.Gorm)
}
