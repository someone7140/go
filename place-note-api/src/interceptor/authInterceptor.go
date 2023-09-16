package interceptor

import (
	"context"

	"github.com/bufbuild/connect-go"
	"golang.org/x/exp/slices"

	userUseCase "placeNote/src/useCase/userUseCase"
)

// 認証が不要なメソッドの名前
var authNotRequiredMethods = []string{
	"/placeNote.UserAccountService/AuthGoogleAccount",
	"/placeNote.UserAccountService/RegisterUserAccount",
	"/placeNote.UserAccountService/LoginByGoogle",
}

func AuthInterceptor() connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			if req.Spec().IsClient {
				// do nothing
				return next(ctx, req)
			}
			// 認証不要なメソッド判定
			if slices.Contains(authNotRequiredMethods, req.Spec().Procedure) {
				// do nothing
				return next(ctx, req)
			}

			// 認証トークンからアカウントのID取得
			authorizationToken := req.Header().Get("Authorization")
			userAccountId, err := userUseCase.GetStrInfoFromToken(userUseCase.UserAccountJwtPropertyName, authorizationToken)
			if err != nil {
				return nil, err
			}
			// コンテキストにアカウントのIDとtokenをセットする
			newCtx := context.WithValue(ctx, userUseCase.UserAccountIdContextKey, userAccountId)
			newCtx = context.WithValue(newCtx, userUseCase.UserAccountTokenContextKey, authorizationToken)
			return next(newCtx, req)
		}
	}
}
