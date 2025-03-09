package middleware

import (
	"context"
	"net/http"
	"wasurena-task-api/db"
)

type contextKey string

const dbQueriesContextKey contextKey = "pgx_queries"

// pgxのクエリをコンテキストに追加するミドルウェア
func WithDbQueries(queries *db.Queries) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), dbQueriesContextKey, queries)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
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
