package server

import (
	"context"
	placeNote "placeNote/src/gen/proto"
	"placeNote/src/useCase/userUseCase"

	"github.com/bufbuild/connect-go"
)

type UserAccountServer struct{}

func (s *UserAccountServer) AuthGoogleAccount(
	ctx context.Context,
	req *connect.Request[placeNote.AuthGoogleAccountRequest],
) (*connect.Response[placeNote.AuthGoogleAccountResponse], error) {
	authGoogleAccountResponse, err := userUseCase.AuthGoogleUserAccount(req.Msg.AuthCode)
	if err != nil {
		return nil, err
	}
	return authGoogleAccountResponse, nil
}

func (s *UserAccountServer) RegisterUserAccount(
	ctx context.Context,
	req *connect.Request[placeNote.RegisterUserAccountRequest],
) (*connect.Response[placeNote.UserAccountResponse], error) {
	userAccountResponse, err := userUseCase.NewRegistrationUserAccountUseCase(req)
	if err != nil {
		return nil, err
	}
	return userAccountResponse, nil
}
