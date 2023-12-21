// Code generated by protoc-gen-connect-go. DO NOT EDIT.
//
// Source: proto/placeNotePostPlaceService.proto

package placeNoteconnect

import (
	context "context"
	errors "errors"
	connect_go "github.com/bufbuild/connect-go"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
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
	// PostPlaceServiceName is the fully-qualified name of the PostPlaceService service.
	PostPlaceServiceName = "placeNote.PostPlaceService"
)

// These constants are the fully-qualified names of the RPCs defined in this package. They're
// exposed at runtime as Spec.Procedure and as the final two segments of the HTTP route.
//
// Note that these are different from the fully-qualified method names used by
// google.golang.org/protobuf/reflect/protoreflect. To convert from these constants to
// reflection-formatted method names, remove the leading slash and convert the remaining slash to a
// period.
const (
	// PostPlaceServiceAddPostPlaceProcedure is the fully-qualified name of the PostPlaceService's
	// AddPostPlace RPC.
	PostPlaceServiceAddPostPlaceProcedure = "/placeNote.PostPlaceService/AddPostPlace"
	// PostPlaceServiceUpdatePostPlaceProcedure is the fully-qualified name of the PostPlaceService's
	// UpdatePostPlace RPC.
	PostPlaceServiceUpdatePostPlaceProcedure = "/placeNote.PostPlaceService/UpdatePostPlace"
	// PostPlaceServiceDeletePostPlaceProcedure is the fully-qualified name of the PostPlaceService's
	// DeletePostPlace RPC.
	PostPlaceServiceDeletePostPlaceProcedure = "/placeNote.PostPlaceService/DeletePostPlace"
	// PostPlaceServiceGetPostPlaceListProcedure is the fully-qualified name of the PostPlaceService's
	// GetPostPlaceList RPC.
	PostPlaceServiceGetPostPlaceListProcedure = "/placeNote.PostPlaceService/GetPostPlaceList"
	// PostPlaceServiceGetPostPlaceByIdProcedure is the fully-qualified name of the PostPlaceService's
	// GetPostPlaceById RPC.
	PostPlaceServiceGetPostPlaceByIdProcedure = "/placeNote.PostPlaceService/GetPostPlaceById"
)

// PostPlaceServiceClient is a client for the placeNote.PostPlaceService service.
type PostPlaceServiceClient interface {
	AddPostPlace(context.Context, *connect_go.Request[proto.AddPostPlaceRequest]) (*connect_go.Response[proto.AddPostPlaceResponse], error)
	UpdatePostPlace(context.Context, *connect_go.Request[proto.UpdatePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error)
	DeletePostPlace(context.Context, *connect_go.Request[proto.DeletePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetPostPlaceList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.GetPostPlaceListResponse], error)
	GetPostPlaceById(context.Context, *connect_go.Request[proto.GetPostPlaceByIdRequest]) (*connect_go.Response[proto.PostPlaceResponse], error)
}

// NewPostPlaceServiceClient constructs a client for the placeNote.PostPlaceService service. By
// default, it uses the Connect protocol with the binary Protobuf Codec, asks for gzipped responses,
// and sends uncompressed requests. To use the gRPC or gRPC-Web protocols, supply the
// connect.WithGRPC() or connect.WithGRPCWeb() options.
//
// The URL supplied here should be the base URL for the Connect or gRPC server (for example,
// http://api.acme.com or https://acme.com/grpc).
func NewPostPlaceServiceClient(httpClient connect_go.HTTPClient, baseURL string, opts ...connect_go.ClientOption) PostPlaceServiceClient {
	baseURL = strings.TrimRight(baseURL, "/")
	return &postPlaceServiceClient{
		addPostPlace: connect_go.NewClient[proto.AddPostPlaceRequest, proto.AddPostPlaceResponse](
			httpClient,
			baseURL+PostPlaceServiceAddPostPlaceProcedure,
			opts...,
		),
		updatePostPlace: connect_go.NewClient[proto.UpdatePostPlaceRequest, emptypb.Empty](
			httpClient,
			baseURL+PostPlaceServiceUpdatePostPlaceProcedure,
			opts...,
		),
		deletePostPlace: connect_go.NewClient[proto.DeletePostPlaceRequest, emptypb.Empty](
			httpClient,
			baseURL+PostPlaceServiceDeletePostPlaceProcedure,
			opts...,
		),
		getPostPlaceList: connect_go.NewClient[emptypb.Empty, proto.GetPostPlaceListResponse](
			httpClient,
			baseURL+PostPlaceServiceGetPostPlaceListProcedure,
			opts...,
		),
		getPostPlaceById: connect_go.NewClient[proto.GetPostPlaceByIdRequest, proto.PostPlaceResponse](
			httpClient,
			baseURL+PostPlaceServiceGetPostPlaceByIdProcedure,
			opts...,
		),
	}
}

// postPlaceServiceClient implements PostPlaceServiceClient.
type postPlaceServiceClient struct {
	addPostPlace     *connect_go.Client[proto.AddPostPlaceRequest, proto.AddPostPlaceResponse]
	updatePostPlace  *connect_go.Client[proto.UpdatePostPlaceRequest, emptypb.Empty]
	deletePostPlace  *connect_go.Client[proto.DeletePostPlaceRequest, emptypb.Empty]
	getPostPlaceList *connect_go.Client[emptypb.Empty, proto.GetPostPlaceListResponse]
	getPostPlaceById *connect_go.Client[proto.GetPostPlaceByIdRequest, proto.PostPlaceResponse]
}

// AddPostPlace calls placeNote.PostPlaceService.AddPostPlace.
func (c *postPlaceServiceClient) AddPostPlace(ctx context.Context, req *connect_go.Request[proto.AddPostPlaceRequest]) (*connect_go.Response[proto.AddPostPlaceResponse], error) {
	return c.addPostPlace.CallUnary(ctx, req)
}

// UpdatePostPlace calls placeNote.PostPlaceService.UpdatePostPlace.
func (c *postPlaceServiceClient) UpdatePostPlace(ctx context.Context, req *connect_go.Request[proto.UpdatePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.updatePostPlace.CallUnary(ctx, req)
}

// DeletePostPlace calls placeNote.PostPlaceService.DeletePostPlace.
func (c *postPlaceServiceClient) DeletePostPlace(ctx context.Context, req *connect_go.Request[proto.DeletePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return c.deletePostPlace.CallUnary(ctx, req)
}

// GetPostPlaceList calls placeNote.PostPlaceService.GetPostPlaceList.
func (c *postPlaceServiceClient) GetPostPlaceList(ctx context.Context, req *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.GetPostPlaceListResponse], error) {
	return c.getPostPlaceList.CallUnary(ctx, req)
}

// GetPostPlaceById calls placeNote.PostPlaceService.GetPostPlaceById.
func (c *postPlaceServiceClient) GetPostPlaceById(ctx context.Context, req *connect_go.Request[proto.GetPostPlaceByIdRequest]) (*connect_go.Response[proto.PostPlaceResponse], error) {
	return c.getPostPlaceById.CallUnary(ctx, req)
}

// PostPlaceServiceHandler is an implementation of the placeNote.PostPlaceService service.
type PostPlaceServiceHandler interface {
	AddPostPlace(context.Context, *connect_go.Request[proto.AddPostPlaceRequest]) (*connect_go.Response[proto.AddPostPlaceResponse], error)
	UpdatePostPlace(context.Context, *connect_go.Request[proto.UpdatePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error)
	DeletePostPlace(context.Context, *connect_go.Request[proto.DeletePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error)
	GetPostPlaceList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.GetPostPlaceListResponse], error)
	GetPostPlaceById(context.Context, *connect_go.Request[proto.GetPostPlaceByIdRequest]) (*connect_go.Response[proto.PostPlaceResponse], error)
}

// NewPostPlaceServiceHandler builds an HTTP handler from the service implementation. It returns the
// path on which to mount the handler and the handler itself.
//
// By default, handlers support the Connect, gRPC, and gRPC-Web protocols with the binary Protobuf
// and JSON codecs. They also support gzip compression.
func NewPostPlaceServiceHandler(svc PostPlaceServiceHandler, opts ...connect_go.HandlerOption) (string, http.Handler) {
	postPlaceServiceAddPostPlaceHandler := connect_go.NewUnaryHandler(
		PostPlaceServiceAddPostPlaceProcedure,
		svc.AddPostPlace,
		opts...,
	)
	postPlaceServiceUpdatePostPlaceHandler := connect_go.NewUnaryHandler(
		PostPlaceServiceUpdatePostPlaceProcedure,
		svc.UpdatePostPlace,
		opts...,
	)
	postPlaceServiceDeletePostPlaceHandler := connect_go.NewUnaryHandler(
		PostPlaceServiceDeletePostPlaceProcedure,
		svc.DeletePostPlace,
		opts...,
	)
	postPlaceServiceGetPostPlaceListHandler := connect_go.NewUnaryHandler(
		PostPlaceServiceGetPostPlaceListProcedure,
		svc.GetPostPlaceList,
		opts...,
	)
	postPlaceServiceGetPostPlaceByIdHandler := connect_go.NewUnaryHandler(
		PostPlaceServiceGetPostPlaceByIdProcedure,
		svc.GetPostPlaceById,
		opts...,
	)
	return "/placeNote.PostPlaceService/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case PostPlaceServiceAddPostPlaceProcedure:
			postPlaceServiceAddPostPlaceHandler.ServeHTTP(w, r)
		case PostPlaceServiceUpdatePostPlaceProcedure:
			postPlaceServiceUpdatePostPlaceHandler.ServeHTTP(w, r)
		case PostPlaceServiceDeletePostPlaceProcedure:
			postPlaceServiceDeletePostPlaceHandler.ServeHTTP(w, r)
		case PostPlaceServiceGetPostPlaceListProcedure:
			postPlaceServiceGetPostPlaceListHandler.ServeHTTP(w, r)
		case PostPlaceServiceGetPostPlaceByIdProcedure:
			postPlaceServiceGetPostPlaceByIdHandler.ServeHTTP(w, r)
		default:
			http.NotFound(w, r)
		}
	})
}

// UnimplementedPostPlaceServiceHandler returns CodeUnimplemented from all methods.
type UnimplementedPostPlaceServiceHandler struct{}

func (UnimplementedPostPlaceServiceHandler) AddPostPlace(context.Context, *connect_go.Request[proto.AddPostPlaceRequest]) (*connect_go.Response[proto.AddPostPlaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("placeNote.PostPlaceService.AddPostPlace is not implemented"))
}

func (UnimplementedPostPlaceServiceHandler) UpdatePostPlace(context.Context, *connect_go.Request[proto.UpdatePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("placeNote.PostPlaceService.UpdatePostPlace is not implemented"))
}

func (UnimplementedPostPlaceServiceHandler) DeletePostPlace(context.Context, *connect_go.Request[proto.DeletePostPlaceRequest]) (*connect_go.Response[emptypb.Empty], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("placeNote.PostPlaceService.DeletePostPlace is not implemented"))
}

func (UnimplementedPostPlaceServiceHandler) GetPostPlaceList(context.Context, *connect_go.Request[emptypb.Empty]) (*connect_go.Response[proto.GetPostPlaceListResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("placeNote.PostPlaceService.GetPostPlaceList is not implemented"))
}

func (UnimplementedPostPlaceServiceHandler) GetPostPlaceById(context.Context, *connect_go.Request[proto.GetPostPlaceByIdRequest]) (*connect_go.Response[proto.PostPlaceResponse], error) {
	return nil, connect_go.NewError(connect_go.CodeUnimplemented, errors.New("placeNote.PostPlaceService.GetPostPlaceById is not implemented"))
}