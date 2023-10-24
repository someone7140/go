package server

import (
	"context"
	placeNote "placeNote/src/gen/proto"
	postUseCase "placeNote/src/useCase/postUseCase"
	userUseCase "placeNote/src/useCase/userUseCase"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PostCategoryServer struct{}

func (s *PostCategoryServer) AddPostCategory(
	ctx context.Context,
	req *connect.Request[placeNote.AddPostCategoryRequest],
) (*connect.Response[emptypb.Empty], error) {
	_, err := postUseCase.AddPostCategory(req.Msg, ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *PostCategoryServer) UpdatePostCategory(
	ctx context.Context,
	req *connect.Request[placeNote.UpdatePostCategoryRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := postUseCase.UpdatePostCategory(req.Msg, ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *PostCategoryServer) DeletePostCategory(
	ctx context.Context,
	req *connect.Request[placeNote.DeletePostCategoryRequest],
) (*connect.Response[emptypb.Empty], error) {
	err := postUseCase.DeletePostCategory(req.Msg.Id, ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&emptypb.Empty{}), nil
}

func (s *PostCategoryServer) GetPostCategoryList(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[placeNote.GetPostCategoryListResponse], error) {
	response, err := postUseCase.GetMyPostCategoryList(ctx.Value(userUseCase.UserAccountIdContextKey).(string))
	if err != nil {
		return nil, err
	}
	return response, nil
}
