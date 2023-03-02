// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: provider/v1/balance.proto

package providerv1connect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	v1 "github.com/getumbeluzi/umbeluzi-server/gen/provider/v1"
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
	// BalanceServiceName is the fully-qualified name of the BalanceService service.
	BalanceServiceName = "provider.v1.BalanceService"
)

// BalanceServiceClient is a client for the provider.v1.BalanceService service.
type BalanceServiceClient interface {
	ListBalances(context.Context, *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error)
	ValidateAccount(context.Context, *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error)
}

// NewBalanceServiceClient constructs a client for the provider.v1.BalanceService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewBalanceServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) BalanceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &balanceServiceClient{
		listBalances: connect_go.NewClient[v1.ListBalancesRequest, v1.ListBalancesResponse](
			httpClient,
			baseURL+"/provider.v1.BalanceService/ListBalances",
			opts...,
		),
		validateAccount: connect_go.NewClient[v1.GetBalanceRequest, v1.GetBalanceResponse](
			httpClient,
			baseURL+"/provider.v1.BalanceService/ValidateAccount",
			opts...,
		),
	}
}

// balanceServiceClient implements BalanceServiceClient.
type balanceServiceClient struct {
	listBalances    *connect_go.Client[v1.ListBalancesRequest, v1.ListBalancesResponse]
	validateAccount *connect_go.Client[v1.GetBalanceRequest, v1.GetBalanceResponse]
}

// ListBalances calls provider.v1.BalanceService.ListBalances.
func (c *balanceServiceClient) ListBalances(ctx context.Context, req *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error) {
	return c.listBalances.CallUnary(ctx, req)
}

// ValidateAccount calls provider.v1.BalanceService.ValidateAccount.
func (c *balanceServiceClient) ValidateAccount(ctx context.Context, req *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error) {
	return c.validateAccount.CallUnary(ctx, req)
}

// BalanceServiceHandler is an implementation of the provider.v1.BalanceService service.
type BalanceServiceHandler interface {
	ListBalances(context.Context, *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error)
	ValidateAccount(context.Context, *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error)
}

// NewBalanceServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewBalanceServiceHandler(svc BalanceServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	mux := http.NewServeMux()
	mux.Handle("/provider.v1.BalanceService/ListBalances", connect_go.NewUnaryHandler(
		"/provider.v1.BalanceService/ListBalances",
		svc.ListBalances,
		opts...,
	))
	mux.Handle("/provider.v1.BalanceService/ValidateAccount", connect_go.NewUnaryHandler(
		"/provider.v1.BalanceService/ValidateAccount",
		svc.ValidateAccount,
		opts...,
	))
	return "/provider.v1.BalanceService/", mux
}

// UnimplementedBalanceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedBalanceServiceHandler struct{}

func (UnimplementedBalanceServiceHandler) ListBalances(context.Context, *connect_go.Request[v1.ListBalancesRequest]) (*connect_go.Response[v1.ListBalancesResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("provider.v1.BalanceService.ListBalances is not implemented"))
}

func (UnimplementedBalanceServiceHandler) ValidateAccount(context.Context, *connect_go.Request[v1.GetBalanceRequest]) (*connect_go.Response[v1.GetBalanceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("provider.v1.BalanceService.ValidateAccount is not implemented"))
}
