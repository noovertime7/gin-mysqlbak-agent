package handler

import (
	"backupAgent/domain/pkg/log"
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/task"
	"context"
)

var TaskSvc service.TaskService

type TaskHandler struct{}

func (t *TaskHandler) TaskAdd(ctx context.Context, in *task.TaskAddInput, out *task.TaskOneMessage) error {
	if err := TaskSvc.TaskAdd(ctx, in); err != nil {
		log.Logger.Error("新增task失败")
		out.Message = "新增失败"
		out.OK = false
		return err
	}
	out.Message = "新增成功"
	out.OK = true
	return nil
}
func (t *TaskHandler) TaskAutoCreate(ctx context.Context, in *task.TaskAutoCreateInPut, out *task.TaskOneMessage) error {
	if err := TaskSvc.TaskAutoCreate(ctx, in); err != nil {
		log.Logger.Error("自动创建任务成功task失败")
		out.Message = "自动创建任务成功失败"
		out.OK = false
		return err
	}
	out.Message = "自动创建任务成功"
	out.OK = true
	return nil
}
func (t *TaskHandler) TaskDelete(ctx context.Context, in *task.TaskIDInput, out *task.TaskOneMessage) error {
	if err := TaskSvc.TaskDelete(ctx, in); err != nil {
		log.Logger.Error("删除task失败")
		out.Message = "删除失败"
		out.OK = false
		return err
	}
	out.Message = "删除成功"
	out.OK = true
	return nil
}
func (t *TaskHandler) TaskDestroy(ctx context.Context, in *task.TaskIDInput, out *task.TaskOneMessage) error {
	if err := TaskSvc.TaskDestroy(ctx, in); err != nil {
		log.Logger.Error("清理task失败")
		out.Message = "清理失败"
		out.OK = false
		return err
	}
	out.Message = "清理成功"
	out.OK = true
	return nil
}
func (t *TaskHandler) RestoreTask(ctx context.Context, in *task.TaskIDInput, out *task.TaskOneMessage) error {
	if err := TaskSvc.TaskRestore(ctx, in); err != nil {
		log.Logger.Error("还原task失败")
		out.Message = "还原失败"
		out.OK = false
		return err
	}
	out.Message = "还原成功"
	out.OK = true
	return nil
}
func (t *TaskHandler) TaskUpdate(ctx context.Context, in *task.TaskUpdateInput, out *task.TaskOneMessage) error {
	log.Logger.Debug("更新task传入数据:", in)
	if err := TaskSvc.TaskUpdate(ctx, in); err != nil {
		log.Logger.Error("更新task失败")
		out.Message = "更新失败"
		out.OK = false
		return err
	}
	out.Message = "更新成功"
	out.OK = true
	return nil
}
func (t *TaskHandler) TaskList(ctx context.Context, in *task.TaskListInput, out *task.TaskListOutPut) error {
	data, err := TaskSvc.TaskList(ctx, in)
	out.TaskListItem = data.TaskListItem
	out.Total = data.Total
	if err != nil {
		log.Logger.Error("查询Task列表失败")
		return err
	}
	out.PageNo = in.PageNo
	out.PageSize = in.PageSize
	return nil
}

func (t *TaskHandler) GetTaskUnscopedList(ctx context.Context, in *task.TaskListInput, out *task.TaskListOutPut) error {
	data, err := TaskSvc.UnscopedTaskList(ctx, in)
	out.TaskListItem = data.TaskListItem
	out.Total = data.Total
	if err != nil {
		log.Logger.Error("查询Task列表失败")
		return err
	}
	return nil
}

func (t *TaskHandler) TaskDetail(ctx context.Context, in *task.TaskIDInput, out *task.TaskDetailOutPut) error {
	data, err := TaskSvc.TaskDetail(ctx, in)
	if err != nil {
		log.Logger.Error("查询Task详情失败")
		return err
	}
	*out = *data
	return nil
}

func (t *TaskHandler) GetDateNumInfo(ctx context.Context, in *task.DateNumInfoInput, out *task.DateNumInfoOut) error {
	data, err := TaskSvc.GetDateNumInfo(ctx, in)
	if err != nil {
		log.Logger.Error("查询日期数量信息失败")
		return err
	}
	*out = *data
	return nil
}
