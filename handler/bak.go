package handler

import (
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/bak"
	"context"
)

type BakHandler struct{}

var bakService *service.BakService

func (b *BakHandler) StartBak(ctx context.Context, in *bak.StartBakInput, out *bak.BakOneMessage) error {
	if err := bakService.StartBak(ctx, in, false); err != nil {
		out.Message = "启动失败"
		out.OK = false
		return err
	}
	out.Message = "启动成功"
	out.OK = true
	return nil
}

func (b *BakHandler) TestBak(ctx context.Context, in *bak.StartBakInput, out *bak.BakOneMessage) error {
	if err := bakService.TestBak(ctx, in); err != nil {
		out.Message = "测试启动失败"
		out.OK = false
		return err
	}
	out.Message = "测试任务启动成功，请一分钟后查看历史记录"
	out.OK = true
	return nil
}

func (b *BakHandler) StopBak(ctx context.Context, in *bak.StopBakInput, out *bak.BakOneMessage) error {
	if err := bakService.StopBak(ctx, in); err != nil {
		out.Message = "停止失败"
		out.OK = false
		return err
	}
	out.Message = "停止成功"
	out.OK = true
	return nil
}
func (b *BakHandler) StartBakByHost(ctx context.Context, in *bak.StartBakByHostInput, out *bak.BakOneMessage) error {
	if err := bakService.StartBakByHost(ctx, in); err != nil {
		out.Message = "启动失败"
		out.OK = false
		return err
	}
	out.Message = "启动成功"
	out.OK = true
	return nil
}
func (b *BakHandler) StopBakByHost(ctx context.Context, in *bak.StopBakByHostInput, out *bak.BakOneMessage) error {
	if err := bakService.StopBakByHost(ctx, in); err != nil {
		out.Message = "停止失败"
		out.OK = false
		return err
	}
	out.Message = "停止成功"
	out.OK = true
	return nil
}
