package custom_middleware

import (
	"context"
	"main/service"
	"strings"

	"github.com/99designs/gqlgen/graphql"
)

type contextAuthKey string

const userAccountIdContextKey contextAuthKey = "userAccountId"

// リクエストヘッダーのトークンからユーザIDを取得してコンテキストに追加するミドルウェア
func SetAuthContextFromHeader(ctx context.Context, next graphql.OperationHandler) graphql.ResponseHandler {
	rc := graphql.GetOperationContext(ctx)
	authHeader := rc.Headers.Get("Authorization")

	if authHeader != "" {
		parts := strings.Split(authHeader, " ")
		if len(parts) == 2 {
			userAccountId, err := service.DecodeAuthToken(parts[1])
			if err == nil && userAccountId != nil {
				ctx = context.WithValue(ctx, userAccountIdContextKey, *userAccountId)
			}
		}
	}

	return next(ctx)
}

// コンテキストからユーザIDを取得する
func GeUserAccountIDFromContext(ctx context.Context) *string {
	userAccountId, ok := ctx.Value(userAccountIdContextKey).(string)
	if !ok {
		return nil
	}
	return &userAccountId
}
