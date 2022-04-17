package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"weather-api/src/pb"
	"weather-api/src/service"

	db "weather-api/src/db"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/improbable-eng/grpc-web/go/grpcweb"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
)

// envLoad 環境毎の設定を読み込み
func envLoad() {
	// 環境変数GO_ENV
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "local")
	}

	err := godotenv.Load(fmt.Sprintf("./config/%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	envLoad()
	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_validator.UnaryServerInterceptor(),
				service.AuthInterceptor,
			),
		),
		grpc.StreamInterceptor(
			grpc_middleware.ChainStreamServer(
				grpc_validator.StreamServerInterceptor(),
			),
		),
	)
	dbEngine, err := db.GetOracleDbEngine()
	if err != nil {
		log.Fatal(err)
	}
	pb.RegisterAuthenticationUserServiceServer(grpcServer, service.NewAuthenticationUserService(dbEngine))
	pb.RegisterGeographicPointServiceServer(grpcServer, service.NewGeographicPointService(dbEngine))

	/*
		port := os.Getenv("PORT")
		if port == "" {
			port = "50051"
		}

		grpcEndpoint := fmt.Sprintf(":%s", port)
		log.Printf("gRPC endpoint [%s]", grpcEndpoint)

		listen, err := net.Listen("tcp", grpcEndpoint)
		if err != nil {
			log.Fatal(err)
		}
		reflection.Register(grpcServer)
		log.Printf("Starting: gRPC Listener [%s]\n", grpcEndpoint)
		log.Fatal(grpcServer.Serve(listen))
	*/

	wrappedServer := grpcweb.WrapServer(
		grpcServer,
		grpcweb.WithOriginFunc(func(origin string) bool {
			return origin == os.Getenv("VIEW_DOMAIN")
		}),
	)
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(wrappedServer.ServeHTTP))
	hs := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	log.Fatal(hs.ListenAndServe())
}
