package server

import (
	"context"
	placeNote "placeNote/src/gen/proto"

	"github.com/bufbuild/connect-go"
)

type UserServer struct{}

func (s *UserServer) RegisterUser(
	ctx context.Context,
	req *connect.Request[placeNote.RegsiterUserRequest],
) (*connect.Response[placeNote.UserResponse], error) {

	res := connect.NewResponse(&placeNote.UserResponse{
		Token:         "saaaa",
		AuthMethod:    placeNote.AuthMethod_EMAIL,
		UserSettingId: "id",
		Name:          "name",
	})
	return res, nil

	/*
		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("AlreadyExists userId"))
	*/
}
