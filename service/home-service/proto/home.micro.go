// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/home.proto

/*
Package shop_srv_home is a generated protocol buffer package.

It is generated from these files:
	proto/home.proto

It has these top-level messages:
	HomeNavListReq
	HomeNav
	HomeNavListResp
	HomeContentListReq
	HomeContentListResp
*/
package shop_srv_home

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

// Client API for HomeService service

type HomeService interface {
	FindHomeNav(ctx context.Context, in *HomeNavListReq, opts ...client.CallOption) (*HomeNavListResp, error)
	FindHomeList(ctx context.Context, in *HomeContentListReq, opts ...client.CallOption) (*HomeContentListResp, error)
}

type homeService struct {
	c    client.Client
	name string
}

func NewHomeService(name string, c client.Client) HomeService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "shop.srv.home"
	}
	return &homeService{
		c:    c,
		name: name,
	}
}

func (c *homeService) FindHomeNav(ctx context.Context, in *HomeNavListReq, opts ...client.CallOption) (*HomeNavListResp, error) {
	req := c.c.NewRequest(c.name, "HomeService.FindHomeNav", in)
	out := new(HomeNavListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *homeService) FindHomeList(ctx context.Context, in *HomeContentListReq, opts ...client.CallOption) (*HomeContentListResp, error) {
	req := c.c.NewRequest(c.name, "HomeService.FindHomeList", in)
	out := new(HomeContentListResp)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for HomeService service

type HomeServiceHandler interface {
	FindHomeNav(context.Context, *HomeNavListReq, *HomeNavListResp) error
	FindHomeList(context.Context, *HomeContentListReq, *HomeContentListResp) error
}

func RegisterHomeServiceHandler(s server.Server, hdlr HomeServiceHandler, opts ...server.HandlerOption) error {
	type homeService interface {
		FindHomeNav(ctx context.Context, in *HomeNavListReq, out *HomeNavListResp) error
		FindHomeList(ctx context.Context, in *HomeContentListReq, out *HomeContentListResp) error
	}
	type HomeService struct {
		homeService
	}
	h := &homeServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&HomeService{h}, opts...))
}

type homeServiceHandler struct {
	HomeServiceHandler
}

func (h *homeServiceHandler) FindHomeNav(ctx context.Context, in *HomeNavListReq, out *HomeNavListResp) error {
	return h.HomeServiceHandler.FindHomeNav(ctx, in, out)
}

func (h *homeServiceHandler) FindHomeList(ctx context.Context, in *HomeContentListReq, out *HomeContentListResp) error {
	return h.HomeServiceHandler.FindHomeList(ctx, in, out)
}
