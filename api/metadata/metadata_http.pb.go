// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.0.0

package metadata

import (
	context "context"

	http "github.com/go-warrior/pkg/transport/http"
	binding "github.com/go-warrior/pkg/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

type MetadataHTTPServer interface {
	GetServiceDesc(context.Context, *GetServiceDescRequest) (*GetServiceDescReply, error)
	ListServices(context.Context, *ListServicesRequest) (*ListServicesReply, error)
}

func RegisterMetadataHTTPServer(s *http.Server, srv MetadataHTTPServer) {
	r := s.Route("/")
	r.GET("/services", _Metadata_ListServices0_HTTP_Handler(srv))
	r.GET("/services/{name}", _Metadata_GetServiceDesc0_HTTP_Handler(srv))
}

func _Metadata_ListServices0_HTTP_Handler(srv MetadataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListServicesRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kratos.api.Metadata/ListServices")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListServices(ctx, req.(*ListServicesRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListServicesReply)
		return ctx.Result(200, reply)
	}
}

func _Metadata_GetServiceDesc0_HTTP_Handler(srv MetadataHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in GetServiceDescRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, "/kratos.api.Metadata/GetServiceDesc")
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetServiceDesc(ctx, req.(*GetServiceDescRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*GetServiceDescReply)
		return ctx.Result(200, reply)
	}
}

type MetadataHTTPClient interface {
	GetServiceDesc(ctx context.Context, req *GetServiceDescRequest, opts ...http.CallOption) (rsp *GetServiceDescReply, err error)
	ListServices(ctx context.Context, req *ListServicesRequest, opts ...http.CallOption) (rsp *ListServicesReply, err error)
}

type MetadataHTTPClientImpl struct {
	cc *http.Client
}

func NewMetadataHTTPClient(client *http.Client) MetadataHTTPClient {
	return &MetadataHTTPClientImpl{client}
}

func (c *MetadataHTTPClientImpl) GetServiceDesc(ctx context.Context, in *GetServiceDescRequest, opts ...http.CallOption) (*GetServiceDescReply, error) {
	var out GetServiceDescReply
	pattern := "/services/{name}"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/kratos.api.Metadata/GetServiceDesc"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *MetadataHTTPClientImpl) ListServices(ctx context.Context, in *ListServicesRequest, opts ...http.CallOption) (*ListServicesReply, error) {
	var out ListServicesReply
	pattern := "/services"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation("/kratos.api.Metadata/ListServices"))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
