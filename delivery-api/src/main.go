package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"./handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// EnvLoad 環境毎の設定を読み込み
func EnvLoad() {
	// 環境変数GO_ENV
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "local")
	}

	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	EnvLoad()
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			os.Getenv("VIEW_DOMAIN"),
		},
		MaxAge: 24 * time.Hour,
	}))
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.POST("/searchStore", handler.StoreSearch)

	router.Run()
}
