package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"wasurena-task-api/db"
	"wasurena-task-api/graph"
	"wasurena-task-api/middleware"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/ast"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5"
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

	// DBの接続設定
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_CONNECT"))
	queries := db.New(conn)
	if err != nil {
		log.Fatal("Error Db Connect")
	}
	defer conn.Close(ctx)

	router := chi.NewRouter()
	router.Use(middleware.WithDbQueries(queries))

	// テスト用

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
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
