// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.7.0
// - protoc             v4.24.2
// source: helloworld/v1/demo.proto

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationDemoSayDemo = "/helloworld.v1.Demo/SayDemo"

type DemoHTTPServer interface {
	SayDemo(context.Context, *Trades) (*Trades, error)
}

func RegisterDemoHTTPServer(s *http.Server, srv DemoHTTPServer) {
	r := s.Route("/")
	r.GET("/trades/{id}", _Demo_SayDemo0_HTTP_Handler(srv))
}

func _Demo_SayDemo0_HTTP_Handler(srv DemoHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in Trades
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationDemoSayDemo)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.SayDemo(ctx, req.(*Trades))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*Trades)
		return ctx.Result(200, reply)
	}
}

type DemoHTTPClient interface {
	SayDemo(ctx context.Context, req *Trades, opts ...http.CallOption) (rsp *Trades, err error)
}

type DemoHTTPClientImpl struct {
	cc *http.Client
}

func NewDemoHTTPClient(client *http.Client) DemoHTTPClient {
	return &DemoHTTPClientImpl{client}
}

func (c *DemoHTTPClientImpl) SayDemo(ctx context.Context, in *Trades, opts ...http.CallOption) (*Trades, error) {
	var out Trades
	pattern := "/trades/{id}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationDemoSayDemo))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}