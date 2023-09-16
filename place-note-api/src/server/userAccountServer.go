package server

import (
	"context"
	placeNote "placeNote/src/gen/proto"
	userUseCase "placeNote/src/useCase/userUseCase"

	"github.com/bufbuild/connect-go"
	"google.golang.org/protobuf/types/known/emptypb"
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

func (s *UserAccountServer) GetUserAccountFromAuthToken(
	ctx context.Context,
	req *connect.Request[emptypb.Empty],
) (*connect.Response[placeNote.UserAccountResponse], error) {
	userAccountResponse, err := userUseCase.GetUserAccountFromUserAccountIdUseCase(
		ctx.Value(userUseCase.UserAccountIdContextKey).(string),
		ctx.Value(userUseCase.UserAccountTokenContextKey).(string),
	)
	if err != nil {
		return nil, err
	}
	return userAccountResponse, nil
}

func (s *UserAccountServer) LoginByGoogle(
	ctx context.Context,
	req *connect.Request[placeNote.AuthGoogleAccountRequest],
) (*connect.Response[placeNote.UserAccountResponse], error) {
	userAccountResponse, err := userUseCase.LoginGoogleUserAccount(req.Msg.AuthCode)
	if err != nil {
		return nil, err
	}
	return userAccountResponse, nil
}
