package handler

import (
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/esbak"
	"context"
)

type EsTaskHandler struct{}

var EsTaskService = service.NewEsTaskService()

func (e *EsTaskHandler) TaskAdd(ctx context.Context, in *esbak.EsBakTaskADDInput, out *esbak.EsOneMessage) error {
	if err := EsTaskService.TaskADD(ctx, in); err != nil {
		out.Message = "添加失败"
		out.OK = false
		return err
	}
	out.Message = "添加成功"
	out.OK = true
	return nil
}
func (e *EsTaskHandler) TaskDelete(ctx context.Context, in *esbak.EsTaskIDInput, out *esbak.EsOneMessage) error {
	if err := EsTaskService.TaskDelete(ctx, in.ID); err != nil {
		out.Message = "删除失败"
		out.OK = false
		return err
	}
	out.Message = "删除成功"
	out.OK = true
	return nil
}
func (e *EsTaskHandler) TaskUpdate(ctx context.Context, in *esbak.EsBakTaskUpdateInput, out *esbak.EsOneMessage) error {
	if err := EsTaskService.TaskUpdate(ctx, in); err != nil {
		out.Message = "更新失败"
		out.OK = false
		return err
	}
	out.Message = "更新成功"
	out.OK = true
	return nil
}
func (e *EsTaskHandler) GetTaskList(ctx context.Context, in *esbak.EsTaskListInput, out *esbak.EsTaskListOutPut) error {
	data, err := EsTaskService.GetTaskList(ctx, in)
	if err != nil {
		return err
	}
	data.EsTaskListOutPutItem = out.EsTaskListOutPutItem
	data.Total = out.Total
	return nil
}
func (e *EsTaskHandler) GetTaskDetail(ctx context.Context, in *esbak.EsTaskIDInput, out *esbak.EsTaskDetailOutPut) error {
	data, err := EsTaskService.GetTaskDetail(ctx, in.ID)
	if err != nil {
		return err
	}
	data.EsTaskInfo = out.EsTaskInfo
	return nil
}
