// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/backupAgent/host/host.proto

package host

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Host service

func NewHostEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Host service

type HostService interface {
	AddHost(ctx context.Context, in *HostAddInput, opts ...client.CallOption) (*HostOneMessage, error)
	DeleteHost(ctx context.Context, in *HostDeleteInput, opts ...client.CallOption) (*HostOneMessage, error)
	UpdateHost(ctx context.Context, in *HostUpdateInput, opts ...client.CallOption) (*HostOneMessage, error)
	GetHostList(ctx context.Context, in *HostListInput, opts ...client.CallOption) (*HostListOutPut, error)
}

type hostService struct {
	c    client.Client
	name string
}

func NewHostService(name string, c client.Client) HostService {
	return &hostService{
		c:    c,
		name: name,
	}
}

func (c *hostService) AddHost(ctx context.Context, in *HostAddInput, opts ...client.CallOption) (*HostOneMessage, error) {
	req := c.c.NewRequest(c.name, "Host.AddHost", in)
	out := new(HostOneMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostService) DeleteHost(ctx context.Context, in *HostDeleteInput, opts ...client.CallOption) (*HostOneMessage, error) {
	req := c.c.NewRequest(c.name, "Host.DeleteHost", in)
	out := new(HostOneMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostService) UpdateHost(ctx context.Context, in *HostUpdateInput, opts ...client.CallOption) (*HostOneMessage, error) {
	req := c.c.NewRequest(c.name, "Host.UpdateHost", in)
	out := new(HostOneMessage)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *hostService) GetHostList(ctx context.Context, in *HostListInput, opts ...client.CallOption) (*HostListOutPut, error) {
	req := c.c.NewRequest(c.name, "Host.GetHostList", in)
	out := new(HostListOutPut)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Host service

type HostHandler interface {
	AddHost(context.Context, *HostAddInput, *HostOneMessage) error
	DeleteHost(context.Context, *HostDeleteInput, *HostOneMessage) error
	UpdateHost(context.Context, *HostUpdateInput, *HostOneMessage) error
	GetHostList(context.Context, *HostListInput, *HostListOutPut) error
}

func RegisterHostHandler(s server.Server, hdlr HostHandler, opts ...server.HandlerOption) error {
	type host interface {
		AddHost(ctx context.Context, in *HostAddInput, out *HostOneMessage) error
		DeleteHost(ctx context.Context, in *HostDeleteInput, out *HostOneMessage) error
		UpdateHost(ctx context.Context, in *HostUpdateInput, out *HostOneMessage) error
		GetHostList(ctx context.Context, in *HostListInput, out *HostListOutPut) error
	}
	type Host struct {
		host
	}
	h := &hostHandler{hdlr}
	return s.Handle(s.NewHandler(&Host{h}, opts...))
}

type hostHandler struct {
	HostHandler
}

func (h *hostHandler) AddHost(ctx context.Context, in *HostAddInput, out *HostOneMessage) error {
	return h.HostHandler.AddHost(ctx, in, out)
}

func (h *hostHandler) DeleteHost(ctx context.Context, in *HostDeleteInput, out *HostOneMessage) error {
	return h.HostHandler.DeleteHost(ctx, in, out)
}

func (h *hostHandler) UpdateHost(ctx context.Context, in *HostUpdateInput, out *HostOneMessage) error {
	return h.HostHandler.UpdateHost(ctx, in, out)
}

func (h *hostHandler) GetHostList(ctx context.Context, in *HostListInput, out *HostListOutPut) error {
	return h.HostHandler.GetHostList(ctx, in, out)
}
