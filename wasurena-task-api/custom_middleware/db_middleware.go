package custom_middleware

import (
	"context"

	"github.com/labstack/echo/v4"

	"wasurena-task-api/db"
)

type contextKey string

const dbQueriesContextKey contextKey = "pgx_queries"

// pgxのクエリをコンテキストに追加するミドルウェア
func WithDbQueries(queries *db.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		{
			return func(c echo.Context) error {
				ctx := context.WithValue(c.Request().Context(), dbQueriesContextKey, c)
				c.SetRequest(c.Request().WithContext(ctx))
				err := next(c)
				return err
			}
		}
	}
}

// コンテキストからpgxのクエリを取得する
func GetDbQueries(ctx context.Context) *db.Queries {
	queries, ok := ctx.Value(dbQueriesContextKey).(*db.Queries)
	if !ok {
		panic("PGX DB not found in context")
	}
	return queries
}
