package service

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func AuthInterceptor(ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler) (interface{}, error) {
	// 認証をスキップするパス
	if info.FullMethod == "/pb.AuthenticationUserService/VerifyGoogleAuthCode" {
		return handler(ctx, req)
	} else {
		userIdSetCtx, err := authorize(ctx)
		if err != nil {
			return nil, err
		}
		return handler(userIdSetCtx, req)
	}
}

func authorize(ctx context.Context) (context.Context, error) {
	// ヘッダーのトークンからユーザIDを取得
	token, err := grpc_auth.AuthFromMD(ctx, "Bearer")
	if err != nil {
		return nil, err
	}
	userId, err := GetUserIdInfoFromJwtToken(token)
	if err != nil {
		return nil, err
	}

	// contextにユーザIDをセット
	return GetUserIdSetContext(ctx, userId), nil
}
