package handler

import (
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/esbak"
	"context"
)

var EsBakService = service.NewEsBakService()

func NewEsBakHandler() *esBakHandler {
	return &esBakHandler{}
}

type esBakHandler struct{}

func (e *esBakHandler) Start(ctx context.Context, in *esbak.StartEsBakInput, out *esbak.EsBakOneMessage) error {
	if err := EsBakService.Start(ctx, in.TaskID); err != nil {
		out.Message = "启动失败"
		out.OK = false
		return err
	}
	out.Message = "启动成功"
	out.OK = true
	return nil
}
func (e *esBakHandler) Stop(ctx context.Context, in *esbak.StopEsBakInput, out *esbak.EsBakOneMessage) error {
	if err := EsBakService.Stop(ctx, in.TaskID); err != nil {
		out.Message = "停止失败"
		out.OK = false
		return err
	}
	out.Message = "停止成功"
	out.OK = true
	return nil
}
