// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/redis.proto

package proto

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"math"
)

import (
	"context"
	"github.com/micro/go-micro/v2/api"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/server"
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

// Api Endpoints for RedisOperation service

func NewRedisOperationEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for RedisOperation service

type RedisOperationService interface {
	SetString(ctx context.Context, in *SetStringRequest, opts ...client.CallOption) (*SetStringResponse, error)
	GetString(ctx context.Context, in *GetStringRequest, opts ...client.CallOption) (*GetStringResponse, error)
}

type redisOperationService struct {
	c    client.Client
	name string
}

func NewRedisOperationService(name string, c client.Client) RedisOperationService {
	return &redisOperationService{
		c:    c,
		name: name,
	}
}

func (c *redisOperationService) SetString(ctx context.Context, in *SetStringRequest, opts ...client.CallOption) (*SetStringResponse, error) {
	req := c.c.NewRequest(c.name, "RedisOperation.SetString", in)
	out := new(SetStringResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redisOperationService) GetString(ctx context.Context, in *GetStringRequest, opts ...client.CallOption) (*GetStringResponse, error) {
	req := c.c.NewRequest(c.name, "RedisOperation.GetString", in)
	out := new(GetStringResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for RedisOperation service

type RedisOperationHandler interface {
	SetString(context.Context, *SetStringRequest, *SetStringResponse) error
	GetString(context.Context, *GetStringRequest, *GetStringResponse) error
}

func RegisterRedisOperationHandler(s server.Server, hdlr RedisOperationHandler, opts ...server.HandlerOption) error {
	type redisOperation interface {
		SetString(ctx context.Context, in *SetStringRequest, out *SetStringResponse) error
		GetString(ctx context.Context, in *GetStringRequest, out *GetStringResponse) error
	}
	type RedisOperation struct {
		redisOperation
	}
	h := &redisOperationHandler{hdlr}
	return s.Handle(s.NewHandler(&RedisOperation{h}, opts...))
}

type redisOperationHandler struct {
	RedisOperationHandler
}

func (h *redisOperationHandler) SetString(ctx context.Context, in *SetStringRequest, out *SetStringResponse) error {
	return h.RedisOperationHandler.SetString(ctx, in, out)
}

func (h *redisOperationHandler) GetString(ctx context.Context, in *GetStringRequest, out *GetStringResponse) error {
	return h.RedisOperationHandler.GetString(ctx, in, out)
}
