// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: counter.proto

package sstconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	sst "github.com/dylanratcliffe/server-streaming-test/sst"
	http "net/http"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// CounterServiceName is the fully-qualified name of the CounterService service.
	CounterServiceName = "sst.CounterService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// CounterServiceCountToProcedure is the fully-qualified name of the CounterService's CountTo RPC.
	CounterServiceCountToProcedure = "/sst.CounterService/CountTo"
)

// CounterServiceClient is a client for the sst.CounterService service.
type CounterServiceClient interface {
	// Counts to a certain number, returning every integer in the process.
	CountTo(context.Context, *connect_go.Request[sst.CountToRequest]) (*connect_go.ServerStreamForClient[sst.CountToResponse], error)
}

// NewCounterServiceClient constructs a client for the sst.CounterService service. By default, it
// uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses, and sends
// uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the connect.WithGRPC() or
// connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewCounterServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) CounterServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &counterServiceClient{
		countTo: connect_go.NewClient[sst.CountToRequest, sst.CountToResponse](
			httpClient,
			baseURL+CounterServiceCountToProcedure,
			opts...,
		),
	}
}

// counterServiceClient implements CounterServiceClient.
type counterServiceClient struct {
	countTo *connect_go.Client[sst.CountToRequest, sst.CountToResponse]
}

// CountTo calls sst.CounterService.CountTo.
func (c *counterServiceClient) CountTo(ctx context.Context, req *connect_go.Request[sst.CountToRequest]) (*connect_go.ServerStreamForClient[sst.CountToResponse], error) {
	return c.countTo.CallServerStream(ctx, req)
}

// CounterServiceHandler is an implementation of the sst.CounterService service.
type CounterServiceHandler interface {
	// Counts to a certain number, returning every integer in the process.
	CountTo(context.Context, *connect_go.Request[sst.CountToRequest], *connect_go.ServerStream[sst.CountToResponse]) error
}

// NewCounterServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewCounterServiceHandler(svc CounterServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle(CounterServiceCountToProcedure, connect_go.NewServerStreamHandler(
		CounterServiceCountToProcedure,
		svc.CountTo,
		opts...,
	))
	return "/sst.CounterService/", mux
}

// UnimplementedCounterServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedCounterServiceHandler struct{}

func (UnimplementedCounterServiceHandler) CountTo(context.Context, *connect_go.Request[sst.CountToRequest], *connect_go.ServerStream[sst.CountToResponse]) error {
	return connect_go.NewError(connect_go.CodeUnimplemented, errors.New("sst.CounterService.CountTo is not implemented"))
}
