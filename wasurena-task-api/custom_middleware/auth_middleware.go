package custom_middleware

import (
	"context"
	"strings"
	"wasurena-task-api/domain"

	"github.com/labstack/echo/v4"
)

type contextAuthKey string

const userAccountIdContextKey contextAuthKey = "userAccountId"

// リクエストヘッダーのトークンからユーザIDを取得してコンテキストに追加するミドルウェア
func WithUserAccountId() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		{
			return func(c echo.Context) error {
				authorizationHeader := c.Request().Header.Get("Authorization")
				authorizationHeaderSplits := strings.Split(authorizationHeader, " ")
				// 認証ヘッダーが設定されてなければそのまま次へ
				if len(authorizationHeaderSplits) < 2 {
					err := next(c)
					return err
				}

				// tokenを複合化してコンテキストに設定
				userAccountId, err := domain.GetUserAccountIdFromToken(authorizationHeaderSplits[1])
				if err != nil {
					err := next(c)
					return err
				}
				ctx := context.WithValue(c.Request().Context(), userAccountIdContextKey, userAccountId)
				c.SetRequest(c.Request().WithContext(ctx))
				err = next(c)
				return err
			}
		}
	}
}

// コンテキストからユーザIDを取得する
func GeUserAccountId(ctx context.Context) *string {
	userAccountId, ok := ctx.Value(userAccountIdContextKey).(*string)
	if !ok {
		return nil
	}
	return userAccountId
}
