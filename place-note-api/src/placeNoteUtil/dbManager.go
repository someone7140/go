package placeNoteUtil

import (
	"context"
	"errors"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoConnection *mongo.Client
)

// SetDbConnection DBの接続
func SetDbConnection() func() {
	// MongoDBの接続設定
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	c, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DB_CONNECTION")))
	if err != nil {
		panic(err)
	}

	mongoConnection = c
	deferFunc := func() {
		c.Disconnect(ctx)
		cancel()
	}

	return deferFunc
}

// GetDbCollection DBの接続
func GetDbCollection(colName string) mongo.Collection {
	col := mongoConnection.Database(os.Getenv("DB_NAME")).Collection(colName)
	if col == nil {
		panic(errors.New("can not get collection"))
	}
	return *col
}
