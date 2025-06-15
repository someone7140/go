package custom_middleware

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"

	"wasurena-task-api/db"
)

type contextDbKey string

const dbQueriesContextKey contextDbKey = "pgxQueries"

// pgxのクエリをコンテキストに追加するミドルウェア
func WithDbQueries(queries *db.Queries) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		{
			return func(c echo.Context) error {
				ctx := context.WithValue(c.Request().Context(), dbQueriesContextKey, queries)
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

// コネクションプールを取得する
func GetConnectionPool() (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(os.Getenv("DB_CONNECT"))
	if err != nil {
		return nil, err
	}

	// 接続プール設定
	maxConn, err := strconv.Atoi(os.Getenv("DB_MAX_CONNS"))
	if err != nil {
		return nil, err
	}
	minConn, err := strconv.Atoi(os.Getenv("DB_MIN_CONNS"))
	if err != nil {
		return nil, err
	}
	config.MaxConns = int32(maxConn) // 最大接続数
	config.MinConns = int32(minConn) // 最小接続数
	config.MaxConnLifetime = time.Hour * 2
	config.MaxConnIdleTime = time.Minute * 30

	pool, err := pgxpool.New(context.Background(), config.ConnString())
	if err != nil {
		return nil, err
	}

	return pool, nil
}
