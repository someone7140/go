package util

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ExecuteDbQuery DBのクエリ発行
func ExecuteDbQuery(queryFunc func(*mongo.Collection) ([]bson.M, error), colName string) ([]bson.M, error) {
	// MongoDBの接続設定
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	c, _ := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))
	defer c.Disconnect(ctx)

	col := c.Database(os.Getenv("DB_NAME")).Collection(colName)
	// 関数の呼び出し結果を返す
	return queryFunc(col)
}
