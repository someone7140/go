package main

import (
	"context"
	"log"
	"os"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph"
	"wasurena-task-api/rest_handler"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"
	"github.com/vektah/gqlparser/v2/gqlerror"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

const defaultPort = "8080"

func main() {
	// envの読み込み
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	ech := echo.New()
	ech.Use(middleware.Recover())

	// DBの接続設定
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_CONNECT"))
	queries := db.New(conn)
	if err != nil {
		log.Fatal("Error Db Connect")
	}
	defer conn.Close(ctx)

	// ミドルウェアの設定
	ech.Use(custom_middleware.WithDbQueries(queries))
	ech.Use(custom_middleware.WithUserAccountId())

	// CORSの設定
	ech.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("FRONTEND_DOMAIN")},
		AllowMethods: []string{
			"DELETE",
			"GET",
			"OPTIONS",
			"PATCH",
			"POST",
			"PUT",
		},
		AllowHeaders: []string{
			"accept",
			"authorization",
			"content-type",
			"user-agent",
			"x-csrftoken",
			"x-requested-with",
		},
		AllowCredentials: true,
	}))

	// GraphQLのディレクティブ設定
	graphQLConfig := graph.Config{Resolvers: &graph.Resolver{}}
	graphQLConfig.Directives.IsAuthenticated = func(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
		if custom_middleware.GeUserAccountID(ctx) == nil {
			return nil, &gqlerror.Error{
				Message: "Authentication Error",
				Extensions: map[string]any{
					"code": 401,
				}}
		}
		return next(ctx)
	}
	srv := handler.New(graph.NewExecutableSchema(graphQLConfig))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	// GraphQLのパス設定
	ech.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	ech.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})
	// LINEのコールバック用パス
	ech.POST("/line_callback", rest_handler.LineCallBackHandler)

	err = ech.Start(":" + port)
	if err != nil {
		panic(err)
	}
}
