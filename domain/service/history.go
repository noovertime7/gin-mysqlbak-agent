package service

import (
	"backupAgent/domain/dao"
	"backupAgent/domain/pkg"
	"backupAgent/domain/pkg/database"
	"backupAgent/domain/pkg/log"
	"backupAgent/proto/backupAgent/bakhistory"
	"context"
	"fmt"
	"strconv"
)

type HistoryService struct{}

func (h *HistoryService) GetHistoryList(ctx context.Context, historyInfo *bakhistory.HistoryListInput) (*bakhistory.HistoryListOutput, error) {
	historyDB := &dao.BakHistory{}
	list, total, err := historyDB.PageList(ctx, database.Gorm, historyInfo)
	if err != nil {
		return nil, err
	}
	if len(list) == 0 {
		log.Logger.Warning("当前历史记录列表为空")
		return &bakhistory.HistoryListOutput{
			Total:              total,
			HistoryListOutItem: []*bakhistory.HistoryListOutItem{},
			PageNo:             historyInfo.PageNo,
			PageSize:           historyInfo.PageSize,
		}, nil
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

func (h *HistoryService) GetHistoryNumInfo(ctx context.Context) (*bakhistory.HistoryNumInfoOut, error) {
	data, err := h.GetHistoryList(ctx, &bakhistory.HistoryListInput{
		Info:      "",
		PageNo:    1,
		PageSize:  pkg.LargePageSize,
		SortField: "",
		SortOrder: "",
	})
	if err != nil {
		return nil, err
	}
	//获取文件大小
	var allSize int
	for _, h := range data.HistoryListOutItem {
		iSize, err := strconv.Atoi(h.FileSize)
		if err != nil {
			return nil, err
		}
		allSize += iSize
	}
	mbAllSize := fmt.Sprintf("%.2fMB", float64(allSize)/float64(1024))
	//获取一周内任务数
	his := &dao.BakHistory{}
	dataList, err := his.FindByDate(ctx, database.Gorm, 7)
	if err != nil {
		return nil, err
	}
	return &bakhistory.HistoryNumInfoOut{
		WeekNums:    int64(len(dataList)),
		AllNums:     data.Total,
		AllFileSize: mbAllSize,
	}, err
}
