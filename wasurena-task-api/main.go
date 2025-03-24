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
	"github.com/go-chi/cors"
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

	router := chi.NewRouter()

	// DBの接続設定
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, os.Getenv("DB_CONNECT"))
	queries := db.New(conn)
	if err != nil {
		log.Fatal("Error Db Connect")
	}
	defer conn.Close(ctx)
	router.Use(middleware.WithDbQueries(queries))
	// CORSの設定
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"DELETE",
			"GET",
			"OPTIONS",
			"PATCH",
			"POST",
			"PUT"},
		AllowedHeaders: []string{"accept",
			"authorization",
			"content-type",
			"user-agent",
			"x-csrftoken",
			"x-requested-with"},
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
	router.Handle("/query", srv)
	router.Handle("/", playground.Handler("GraphQL playground", "/query"))

	err = http.ListenAndServe(":"+port, router)
	if err != nil {
		panic(err)
	}
}
