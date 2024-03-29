// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/placeNoteGeolocationService.proto

package placeNoteconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	http "net/http"
	proto "placeNote/src/gen/proto"
	strings "strings"
)

// This is a compile-time assertion to ensure that this generated file and the connect package are
// compatible. If you get a compiler error that this constant is not defined, this code was
// generated with a version of connect newer than the one compiled into your binary. You can fix the
// problem by either regenerating this code with an older version of connect or updating the connect
// version compiled into your binary.
const _ = connect_go.IsAtLeastVersion0_1_0

const (
	// GeolocationServiceName is the fully-qualified name of the GeolocationService service.
	GeolocationServiceName = "placeNote.GeolocationService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// GeolocationServiceGetLatLonFromAddressProcedure is the fully-qualified name of the
	// GeolocationService's GetLatLonFromAddress RPC.
	GeolocationServiceGetLatLonFromAddressProcedure = "/placeNote.GeolocationService/GetLatLonFromAddress"
)

// GeolocationServiceClient is a client for the placeNote.GeolocationService service.
type GeolocationServiceClient interface {
	GetLatLonFromAddress(context.Context, *connect_go.Request[proto.GetLatLonFromAddressRequest]) (*connect_go.Response[proto.GetLatLonFromAddressResponse], error)
}

// NewGeolocationServiceClient constructs a client for the placeNote.GeolocationService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewGeolocationServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) GeolocationServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &geolocationServiceClient{
		getLatLonFromAddress: connect_go.NewClient[proto.GetLatLonFromAddressRequest, proto.GetLatLonFromAddressResponse](
			httpClient,
			baseURL+GeolocationServiceGetLatLonFromAddressProcedure,
			opts...,
		),
	}
}

// geolocationServiceClient implements GeolocationServiceClient.
type geolocationServiceClient struct {
	getLatLonFromAddress *connect_go.Client[proto.GetLatLonFromAddressRequest, proto.GetLatLonFromAddressResponse]
}

// GetLatLonFromAddress calls placeNote.GeolocationService.GetLatLonFromAddress.
func (c *geolocationServiceClient) GetLatLonFromAddress(ctx context.Context, req *connect_go.Request[proto.GetLatLonFromAddressRequest]) (*connect_go.Response[proto.GetLatLonFromAddressResponse], error) {
	return c.getLatLonFromAddress.CallUnary(ctx, req)
}

// GeolocationServiceHandler is an implementation of the placeNote.GeolocationService service.
type GeolocationServiceHandler interface {
	GetLatLonFromAddress(context.Context, *connect_go.Request[proto.GetLatLonFromAddressRequest]) (*connect_go.Response[proto.GetLatLonFromAddressResponse], error)
}

// NewGeolocationServiceHandler builds an HTTP handler from the service implementation. It returns
// the path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewGeolocationServiceHandler(svc GeolocationServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	geolocationServiceGetLatLonFromAddressHandler := connect_go.NewUnaryHandler(
		GeolocationServiceGetLatLonFromAddressProcedure,
		svc.GetLatLonFromAddress,
		opts...,
	)
	return "/placeNote.GeolocationService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case GeolocationServiceGetLatLonFromAddressProcedure:
			geolocationServiceGetLatLonFromAddressHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedGeolocationServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedGeolocationServiceHandler struct{}

func (UnimplementedGeolocationServiceHandler) GetLatLonFromAddress(context.Context, *connect_go.Request[proto.GetLatLonFromAddressRequest]) (*connect_go.Response[proto.GetLatLonFromAddressResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("placeNote.GeolocationService.GetLatLonFromAddress is not implemented"))
}
