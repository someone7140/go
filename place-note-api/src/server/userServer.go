package server

import (
	"context"
	placeNote "placeNote/src/gen/proto"

	"github.com/bufbuild/connect-go"
)

type UserServer struct{}

func (s *UserServer) AuthGoogleCode(
	ctx context.Context,
	req *connect.Request[placeNote.AuthGoogleCodeRequest],
) (*connect.Response[placeNote.AuthGoogleCodeResponse], error) {
	res := connect.NewResponse(&placeNote.AuthGoogleCodeResponse{
		AuthGoogleToken: "token",
	})
	return res, nil
}
