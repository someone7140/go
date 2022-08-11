package authRepository

import (
	"context"
	"time"

	"coda-api/src/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// EmailAuthRegister メール認証の登録
func EmailAuthRegister(
	userID string, email string, password string, token string,
) error {
	now := time.Now().UTC()
	registerBson := bson.M{
		"_id":             userID,
		"email":           email,
		"password":        password,
		"user_setting_id": userID,
		"name":            "dummy",
		"status":          "confirming",
		"user_type":       "normal",
		"email_auth": bson.M{
			"email":       email,
			"token":       token,
			"expire_date": now.AddDate(0, 0, 1),
		},
		"registration_date": now,
		"login_count":       0,
		"categories":        []string{},
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.InsertOne(context.Background(), registerBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return err
	}
	return nil
}

// ChangePassword パスワードの変更
func ChangePassword(
	password string, userID string,
) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: userID}}
		update := bson.D{{Key: "$set",
			Value: bson.D{{Key: "password", Value: password}}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return err
	}
	return nil
}

// DeletePasswordReset パスワードリセット情報の削除
func DeletePasswordReset(
	userID string,
) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: userID}}
		update := bson.D{{Key: "$unset",
			Value: bson.D{{Key: "password_reset", Value: ""}}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return err
	}
	return nil
}

// AddPasswordReset パスワードリセット情報の追加
func AddPasswordReset(token string, userID string) error {
	now := time.Now().UTC()
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: userID}}
		update := bson.D{{Key: "$set",
			Value: bson.D{{Key: "password_reset", Value: bson.M{
				"token":       token,
				"expire_date": now.AddDate(0, 0, 1),
			}}}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "user")
	if err != nil {
		return err
	}
	return nil
}
