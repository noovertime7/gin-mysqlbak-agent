package handler

import (
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/bakhistory"
	"context"
)

var HistoryService service.HistoryService

type HistoryHandler struct{}

func (h *HistoryHandler) GetHistoryList(ctx context.Context, in *bakhistory.HistoryListInput, out *bakhistory.HistoryListOutput) error {
	data, err := HistoryService.GetHistoryList(ctx, in)
	if err != nil {
		log.Logger.Error("获取历史记录列表失败", err)
		return err
	}
	out.HistoryListOutItem = data.HistoryListOutItem
	out.Total = data.Total
	return nil
}

func (h *HistoryHandler) DeleteHistory(ctx context.Context, in *bakhistory.HistoryIDInput, out *bakhistory.HistoryOneMessage) error {
	if err := HistoryService.DeleteHistory(ctx, in); err != nil {
		out.Message = "删除失败"
		out.OK = false
		return err
	}
	out.Message = "删除成功"
	out.OK = true
	return nil
}
