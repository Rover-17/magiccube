// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: dashboard.proto

/*
Package dashboard is a generated protocol buffer package.

It is generated from these files:
	dashboard.proto

It has these top-level messages:
	GetNodeInfosRequest
	GetNodeInfosResponse
	NodeInfoData
	Position
	GetTxListRequest
	GetTxListResponse
	TxListData
	Tx
	GetBlockListRequest
	GetBlockListResponse
	BlockData
	Block
	GetBlockInfoRequest
	GetBlockInfoResponse
	BlockInfoData
	TxList
	GetRequirementNumByDayRequest
	GetRequirementNumByDayResponse
	RequirementNumByDayData
	GetTxNumRequest
	GetTxNumResponse
	GetAssetNumByDayRequest
	GetAssetNumByDayResponse
	AssetNumByDayData
	GetAccountNumByDayRequest
	GetAccountNumByDayResponse
	AccountNumByDayData
	GetTxAmountRequest
	GetTxAmountResponse
	GetTxNumByDayRequest
	GetTxNumByDayResponse
	TxNumByDayData
	Num
	GetTxAmountByDayRequest
	GetTxAmountByDayResponse
	TxAmountByDay
	GetAllTypeTotalRequest
	GetAllTypeTotalResponse
	AllTypeData
*/
package dashboard

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Dashboard service

type DashboardClient interface {
	GetBlockList(ctx context.Context, in *GetBlockListRequest, opts ...client.CallOption) (*GetBlockListResponse, error)
	GetBlockInfo(ctx context.Context, in *GetBlockInfoRequest, opts ...client.CallOption) (*GetBlockInfoResponse, error)
	GetNodeInfos(ctx context.Context, in *GetNodeInfosRequest, opts ...client.CallOption) (*GetNodeInfosResponse, error)
	GetTxList(ctx context.Context, in *GetTxListRequest, opts ...client.CallOption) (*GetTxListResponse, error)
	GetTxNum(ctx context.Context, in *GetTxNumRequest, opts ...client.CallOption) (*GetTxNumResponse, error)
	GetRequirementNumByDay(ctx context.Context, in *GetRequirementNumByDayRequest, opts ...client.CallOption) (*GetRequirementNumByDayResponse, error)
	GetAssetNumByDay(ctx context.Context, in *GetAssetNumByDayRequest, opts ...client.CallOption) (*GetAssetNumByDayResponse, error)
	GetAccountNumByDay(ctx context.Context, in *GetAccountNumByDayRequest, opts ...client.CallOption) (*GetAccountNumByDayResponse, error)
	GetTxAmount(ctx context.Context, in *GetTxAmountRequest, opts ...client.CallOption) (*GetTxAmountResponse, error)
	GetTxNumByDay(ctx context.Context, in *GetTxNumByDayRequest, opts ...client.CallOption) (*GetTxNumByDayResponse, error)
	GetTxAmountByDay(ctx context.Context, in *GetTxAmountByDayRequest, opts ...client.CallOption) (*GetTxAmountByDayResponse, error)
	GetAllTypeTotal(ctx context.Context, in *GetAllTypeTotalRequest, opts ...client.CallOption) (*GetAllTypeTotalResponse, error)
}

type dashboardClient struct {
	c           client.Client
	serviceName string
}

func NewDashboardClient(serviceName string, c client.Client) DashboardClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "dashboard"
	}
	return &dashboardClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *dashboardClient) GetBlockList(ctx context.Context, in *GetBlockListRequest, opts ...client.CallOption) (*GetBlockListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetBlockList", in)
	out := new(GetBlockListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetBlockInfo(ctx context.Context, in *GetBlockInfoRequest, opts ...client.CallOption) (*GetBlockInfoResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetBlockInfo", in)
	out := new(GetBlockInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetNodeInfos(ctx context.Context, in *GetNodeInfosRequest, opts ...client.CallOption) (*GetNodeInfosResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetNodeInfos", in)
	out := new(GetNodeInfosResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetTxList(ctx context.Context, in *GetTxListRequest, opts ...client.CallOption) (*GetTxListResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetTxList", in)
	out := new(GetTxListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetTxNum(ctx context.Context, in *GetTxNumRequest, opts ...client.CallOption) (*GetTxNumResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetTxNum", in)
	out := new(GetTxNumResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetRequirementNumByDay(ctx context.Context, in *GetRequirementNumByDayRequest, opts ...client.CallOption) (*GetRequirementNumByDayResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetRequirementNumByDay", in)
	out := new(GetRequirementNumByDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetAssetNumByDay(ctx context.Context, in *GetAssetNumByDayRequest, opts ...client.CallOption) (*GetAssetNumByDayResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetAssetNumByDay", in)
	out := new(GetAssetNumByDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetAccountNumByDay(ctx context.Context, in *GetAccountNumByDayRequest, opts ...client.CallOption) (*GetAccountNumByDayResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetAccountNumByDay", in)
	out := new(GetAccountNumByDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetTxAmount(ctx context.Context, in *GetTxAmountRequest, opts ...client.CallOption) (*GetTxAmountResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetTxAmount", in)
	out := new(GetTxAmountResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetTxNumByDay(ctx context.Context, in *GetTxNumByDayRequest, opts ...client.CallOption) (*GetTxNumByDayResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetTxNumByDay", in)
	out := new(GetTxNumByDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetTxAmountByDay(ctx context.Context, in *GetTxAmountByDayRequest, opts ...client.CallOption) (*GetTxAmountByDayResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetTxAmountByDay", in)
	out := new(GetTxAmountByDayResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *dashboardClient) GetAllTypeTotal(ctx context.Context, in *GetAllTypeTotalRequest, opts ...client.CallOption) (*GetAllTypeTotalResponse, error) {
	req := c.c.NewRequest(c.serviceName, "Dashboard.GetAllTypeTotal", in)
	out := new(GetAllTypeTotalResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Dashboard service

type DashboardHandler interface {
	GetBlockList(context.Context, *GetBlockListRequest, *GetBlockListResponse) error
	GetBlockInfo(context.Context, *GetBlockInfoRequest, *GetBlockInfoResponse) error
	GetNodeInfos(context.Context, *GetNodeInfosRequest, *GetNodeInfosResponse) error
	GetTxList(context.Context, *GetTxListRequest, *GetTxListResponse) error
	GetTxNum(context.Context, *GetTxNumRequest, *GetTxNumResponse) error
	GetRequirementNumByDay(context.Context, *GetRequirementNumByDayRequest, *GetRequirementNumByDayResponse) error
	GetAssetNumByDay(context.Context, *GetAssetNumByDayRequest, *GetAssetNumByDayResponse) error
	GetAccountNumByDay(context.Context, *GetAccountNumByDayRequest, *GetAccountNumByDayResponse) error
	GetTxAmount(context.Context, *GetTxAmountRequest, *GetTxAmountResponse) error
	GetTxNumByDay(context.Context, *GetTxNumByDayRequest, *GetTxNumByDayResponse) error
	GetTxAmountByDay(context.Context, *GetTxAmountByDayRequest, *GetTxAmountByDayResponse) error
	GetAllTypeTotal(context.Context, *GetAllTypeTotalRequest, *GetAllTypeTotalResponse) error
}

func RegisterDashboardHandler(s server.Server, hdlr DashboardHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Dashboard{hdlr}, opts...))
}

type Dashboard struct {
	DashboardHandler
}

func (h *Dashboard) GetBlockList(ctx context.Context, in *GetBlockListRequest, out *GetBlockListResponse) error {
	return h.DashboardHandler.GetBlockList(ctx, in, out)
}

func (h *Dashboard) GetBlockInfo(ctx context.Context, in *GetBlockInfoRequest, out *GetBlockInfoResponse) error {
	return h.DashboardHandler.GetBlockInfo(ctx, in, out)
}

func (h *Dashboard) GetNodeInfos(ctx context.Context, in *GetNodeInfosRequest, out *GetNodeInfosResponse) error {
	return h.DashboardHandler.GetNodeInfos(ctx, in, out)
}

func (h *Dashboard) GetTxList(ctx context.Context, in *GetTxListRequest, out *GetTxListResponse) error {
	return h.DashboardHandler.GetTxList(ctx, in, out)
}

func (h *Dashboard) GetTxNum(ctx context.Context, in *GetTxNumRequest, out *GetTxNumResponse) error {
	return h.DashboardHandler.GetTxNum(ctx, in, out)
}

func (h *Dashboard) GetRequirementNumByDay(ctx context.Context, in *GetRequirementNumByDayRequest, out *GetRequirementNumByDayResponse) error {
	return h.DashboardHandler.GetRequirementNumByDay(ctx, in, out)
}

func (h *Dashboard) GetAssetNumByDay(ctx context.Context, in *GetAssetNumByDayRequest, out *GetAssetNumByDayResponse) error {
	return h.DashboardHandler.GetAssetNumByDay(ctx, in, out)
}

func (h *Dashboard) GetAccountNumByDay(ctx context.Context, in *GetAccountNumByDayRequest, out *GetAccountNumByDayResponse) error {
	return h.DashboardHandler.GetAccountNumByDay(ctx, in, out)
}

func (h *Dashboard) GetTxAmount(ctx context.Context, in *GetTxAmountRequest, out *GetTxAmountResponse) error {
	return h.DashboardHandler.GetTxAmount(ctx, in, out)
}

func (h *Dashboard) GetTxNumByDay(ctx context.Context, in *GetTxNumByDayRequest, out *GetTxNumByDayResponse) error {
	return h.DashboardHandler.GetTxNumByDay(ctx, in, out)
}

func (h *Dashboard) GetTxAmountByDay(ctx context.Context, in *GetTxAmountByDayRequest, out *GetTxAmountByDayResponse) error {
	return h.DashboardHandler.GetTxAmountByDay(ctx, in, out)
}

func (h *Dashboard) GetAllTypeTotal(ctx context.Context, in *GetAllTypeTotalRequest, out *GetAllTypeTotalResponse) error {
	return h.DashboardHandler.GetAllTypeTotal(ctx, in, out)
}
