package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"placeNote/src/gen/proto/placeNoteconnect"
	"placeNote/src/interceptor"
	"placeNote/src/placeNoteUtil"
	"placeNote/src/server"

	"github.com/bufbuild/connect-go"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

// envLoad 環境毎の設定を読み込み
func envLoad() {
	// 環境変数GO_ENV
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "local")
	}

	err := godotenv.Load(fmt.Sprintf("../env/%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	envLoad()
	mux := http.NewServeMux()
	interceptors := connect.WithInterceptors(interceptor.NewValidationInterceptor(), interceptor.AuthInterceptor())

	// 各種サーバーの登録
	userAccountServerPath, userAccountServerHandler := placeNoteconnect.NewUserAccountServiceHandler(&server.UserAccountServer{}, interceptors)
	mux.Handle(userAccountServerPath, userAccountServerHandler)
	geolocationServerPath, geolocationServerHandler := placeNoteconnect.NewGeolocationServiceHandler(&server.GeoLocationServer{}, interceptors)
	mux.Handle(geolocationServerPath, geolocationServerHandler)
	postCategoryServerPath, postCategoryServerHandler := placeNoteconnect.NewPostCategoryServiceHandler(&server.PostCategoryServer{}, interceptors)
	mux.Handle(postCategoryServerPath, postCategoryServerHandler)

	// CORSの設定
	c := cors.New(cors.Options{
		// 許可したいHTTPメソッドの一覧
		AllowedMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowedHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
			"Connect-Protocol-Version",
		},
		// 許可したいアクセス元の一覧
		AllowedOrigins:   []string{os.Getenv("VIEW_DOMAIN")},
		AllowCredentials: true,
	}).Handler(mux)

	// DB接続
	deferFunc := placeNoteUtil.SetDbConnection()
	defer deferFunc()

	// サーバの起動
	http.ListenAndServe(
		"localhost:8080", c,
	)
}
