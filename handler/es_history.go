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

type esHistoryHandler struct{}

func (e *esHistoryHandler) GetEsHistoryList(ctx context.Context, in *esbak.GetEsHistoryListInput, out *esbak.ESHistoryListOutput) error {
	data, err := ESHistoryService.GetESHistoryList(ctx, in)
	if err != nil {
		log.Logger.Error("获取历史记录列表失败", err)
		return err
	}
	out.EsHistoryListOutItem = data.EsHistoryListOutItem
	out.Total = data.Total
	out.PageNo = data.PageNo
	out.PageSize = data.PageSize
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

func (e *esHistoryHandler) GetEsHistoryDetail(ctx context.Context, in *esbak.ESHistoryIDInput, out *esbak.EsHistoryDetailOut) error {
	data, err := ESHistoryService.GetEsHistoryDetail(ctx, in)
	if err != nil {
		log.Logger.Error("获取历史记录详情失败", err)
		return err
	}
	out.ESTaskDetail = data.GetESTaskDetail()
	out.EsHistoryDetail = data.EsHistoryDetail
	out.EsHostDetail = data.GetEsHostDetail()
	return nil
}

func (e *esHistoryHandler) GetEsHistoryNumInfo(ctx context.Context, em *esbak.EsHistoryEmpty, out *esbak.EsHistoryNumInfoOut) error {
	data, err := ESHistoryService.GetEsHistoryNumInfo(ctx)
	if err != nil {
		log.Logger.Error("获取历史记录数量信息失败", err)
		return err
	}
	out.AllNums = data.AllNums
	out.WeekNums = data.WeekNums
	out.FailNums = data.FailNums
	return nil
}
