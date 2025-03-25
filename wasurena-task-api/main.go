package main

import (
	"context"
	"log"
	"os"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/graph"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"

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
	ech.Use(custom_middleware.WithDbQueries(queries))

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

	// GraphQLのルート設定
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})

	srv.SetQueryCache(lru.New[*ast.QueryDocument](1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New[string](100),
	})

	ech.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	ech.GET("/playground", func(c echo.Context) error {
		playground.Handler("GraphQL playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	err = ech.Start(":" + port)
	if err != nil {
		panic(err)
	}
}
