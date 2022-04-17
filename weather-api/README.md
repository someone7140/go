`src/proto`配下で以下を実行

```
protoc -I=. -I=${GOPATH}/pkg/mod -I=${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.7 --go_out="plugins=grpc:./src" --validate_out="lang=go:./src" --js_out=import_style=commonjs:${OUTPUT_JS_GRPC_PATH} --grpc-web_out=import_style=commonjs,mode=grpcwebtext:${OUTPUT_JS_GRPC_PATH} ./src/proto/*.proto
```

`grpc_cli ls localhost:50051 pb.AuthenticationUserService`で起動確認
https://dev.classmethod.jp/articles/golang-grpc-sample-project/

sql-migrate を使用して DB の migration を行う
https://qiita.com/k-kurikuri/items/946e2bf8c79176ef3ff0
以下で環境毎に migrate を適用

```
sql-migrate up -config=dbconfig.yml -env="development" -limit=0
```
