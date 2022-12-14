package handler

import (
	"backupAgent/domain/service"
	"backupAgent/proto/backupAgent/host"
	"context"
)

type HostHandler struct{}

var s service.HostService

func (h *HostHandler) AddHost(ctx context.Context, in *host.HostAddInput, out *host.HostOneMessage) error {
	if err := s.HostAdd(ctx, in); err != nil {
		out.Message = err.Error()
		out.OK = false
		return err
	}
	out.Message = "新增成功"
	out.OK = true
	return nil
}
func (h *HostHandler) DeleteHost(ctx context.Context, in *host.HostIDInput, out *host.HostOneMessage) error {
	if err := s.HostDelete(ctx, in.ID); err != nil {
		out.Message = err.Error()
		out.OK = false
		return err
	}
	out.Message = "删除成功"
	out.OK = true
	return nil
}
func (h *HostHandler) UpdateHost(ctx context.Context, in *host.HostUpdateInput, out *host.HostOneMessage) error {
	if err := s.HostUpdate(ctx, in); err != nil {
		out.Message = err.Error()
		out.OK = false
		return err
	}
	out.Message = "修改成功"
	out.OK = true
	return nil
}

func (h *HostHandler) GetHostList(ctx context.Context, in *host.HostListInput, out *host.HostListOutPut) error {
	var err error
	data, err := s.GetHostList(ctx, in)
	if err != nil {
		return err
	}
	out.Total = data.Total
	out.ListItem = data.ListItem
	out.PageNo = in.PageNo
	out.PageSize = in.PageSize
	return nil
}

func (h *HostHandler) TestHost(ctx context.Context, in *host.HostIDInput, out *host.HostOneMessage) error {
	if err := service.TestHost(ctx, in.ID); err != nil {
		return err
	}
	out.Message = "主机在线，测试成功"
	out.OK = true
	return nil
}

func (h *HostHandler) GetHostNames(ctx context.Context, in *host.HostNamesInput, out *host.HostNames) error {
	data, err := s.GetHostNames(ctx, in.Type)
	if err != nil {
		return err
	}
	out.HostNameInfo = data.GetHostNameInfo()
	return nil
}
