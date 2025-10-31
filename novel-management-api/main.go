package main

import (
	"log"
	"main/db/query"
	"main/graph"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECT")), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	query.SetDefault(db)

	// GraphQLハンドラーの作成
	playgroundHandler := playground.Handler("GraphQL", "/query")
	app.Get("/playground", adaptor.HTTPHandler(playgroundHandler))
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{
		Resolvers: &graph.Resolver{},
	}))
	app.Post("/query", adaptor.HTTPHandler(srv))

	log.Fatal(app.Listen(":" + port))
}
