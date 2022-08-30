package handler

import (
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/esbak"
	"context"
)

func NewEsHistoryHandler() *esHistoryHandler {
	return &esHistoryHandler{}
}

var ESHistoryService = service.NewESHistoryService()

type esHistoryHandler struct {
}

func (e *esHistoryHandler) GetEsHistoryList(ctx context.Context, in *esbak.GetEsHistoryListInput, out *esbak.ESHistoryListOutput) error {
	data, err := ESHistoryService.GetESHistoryList(ctx, in)
	if err != nil {
		log.Logger.Error("获取历史记录列表失败", err)
		return err
	}
	out.EsHistoryListOutItem = data.GetEsHistoryListOutItem()
	out.Total = data.GetTotal()
	return nil
}
func (e *esHistoryHandler) DeleteESHistory(ctx context.Context, in *esbak.ESHistoryIDInput, out *esbak.ESHistoryOneMessage) error {
	if err := ESHistoryService.DeleteEsHistory(ctx, in); err != nil {
		out.Message = "删除失败"
		out.OK = false
		return err
	}
	out.Message = "删除成功"
	out.OK = true
	return nil
}
