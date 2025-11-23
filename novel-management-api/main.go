package main

import (
	"context"
	"log"
	"main/custom_middleware"
	"main/db/query"
	"main/graph"
	"os"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const defaultPort = "8080"

func main() {
	// envの読み込み
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// データベース接続
	gormConfig := &gorm.Config{Logger: logger.Default.LogMode(logger.Silent)}
	if os.Getenv("DEV_MODE") == "true" {
		gormConfig = &gorm.Config{}
	}
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECT")), gormConfig)
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)

	// GraphQLハンドラーの作成
	playgroundHandler := playground.Handler("GraphQL", "/query")
	app.Get("/playground", adaptor.HTTPHandler(playgroundHandler))
	// GraphQLのディレクティブ設定
	graphQLConfig := graph.Config{Resolvers: &graph.Resolver{}}
	graphQLConfig.Directives.IsAuthenticated = func(ctx context.Context, obj any, next graphql.Resolver) (any, error) {
		if custom_middleware.GeUserAccountIDFromContext(ctx) == nil {
			return nil, &gqlerror.Error{
				Message: "Authentication Error",
				Extensions: map[string]any{
					"code": 401,
				}}
		}
		return next(ctx)
	}

	app.Use(cors.New(cors.Config{
		AllowOrigins:     os.Getenv("FRONTEND_DOMAIN"),
		AllowMethods:     "GET,POST,HEAD,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization,Apollo-Require-Preflight",
		AllowCredentials: true,
	}))

	srv := handler.New(graph.NewExecutableSchema(graphQLConfig))
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.Options{})
	if os.Getenv("DEV_MODE") == "true" {
		srv.Use(extension.Introspection{})
	}
	srv.AroundOperations(custom_middleware.SetAuthContextFromHeader)
	app.Post("/graphql", adaptor.HTTPHandler(srv))

	log.Fatal(app.Listen(":" + port))
}
