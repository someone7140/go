package authRepository

import (
	"context"
	"time"

	"coda-api/src/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UpdateLoginInfo ログイン情報の更新
func UpdateLoginInfo(userID string, loginCount int64) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: userID}}
		update := bson.D{{Key: "$set",
			Value: bson.D{{Key: "login_count", Value: loginCount}, {Key: "login_date", Value: time.Now().UTC()}}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "user")
	return err
}
