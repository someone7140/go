package server

import (
	"context"
	placeNote "placeNote/src/gen/proto"
	postUseCase "placeNote/src/useCase/postUseCase"
	userUseCase "placeNote/src/useCase/userUseCase"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostPlaceServer struct{}

func (s *PostPlaceServer) AddPostPlace(
	ctx context.Context,
	req *connect.Request[placeNote.AddPostPlaceRequest],
) (*connect.Response[emptypb.Empty], error) {
	_, err := postUseCase.AddPostPlace(req.Msg, ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *PostPlaceServer) UpdatePostPlace(
	ctx context.Context,
	req *connect.Request[placeNote.UpdatePostPlaceRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := postUseCase.UpdatePostPlace(req.Msg, ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *PostPlaceServer) DeletePostPlace(
	ctx context.Context,
	req *connect.Request[placeNote.DeletePostPlaceRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := postUseCase.DeletePostPlace(req.Msg, ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *PostPlaceServer) GetPostPlaceList(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[placeNote.GetPostPlaceListResponse], error) {
	response, err := postUseCase.GetPostPlaceList(ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (s *PostPlaceServer) GetPostPlaceById(
	ctx context.Context,
	req *connect.Request[placeNote.GetPostPlaceByIdRequest],
) (*connect.Response[placeNote.PostPlaceResponse], error) {
	response, err := postUseCase.GetPostPlaceById(ctx.Value(userUseCase.UserAccountIdContextKey).(string), req.Msg.Id)
	if err != nil {
		return nil, err
	}
	return response, nil
}
