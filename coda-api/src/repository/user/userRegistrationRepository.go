package userRepository

import (
	userModel "coda-api/src/model/user"
	"coda-api/src/util"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewRegistrationUser ユーザの新規登録
func NewRegistrationUser(userEntity userModel.UserRegisterEntity) error {
	now := time.Now().UTC()
	registerBson := bson.M{
		"_id":             userEntity.ID,
		"email":           map[bool]*string{true: &userEntity.Email, false: nil}[userEntity.Email != ""],
		"password":        map[bool]*string{true: &userEntity.Password, false: nil}[userEntity.Password != ""],
		"facebook_id":     map[bool]*string{true: &userEntity.FacebookID, false: nil}[userEntity.FacebookID != ""],
		"google_id":       map[bool]*string{true: &userEntity.GoogleID, false: nil}[userEntity.GoogleID != ""],
		"user_setting_id": userEntity.UserSettingID,
		"name":            userEntity.Name,
		"status":          userEntity.Status,
		"user_type":       userEntity.UserType,
		"icon_url":        userEntity.IconURL,
		"attribute": bson.M{
			"gender":     userEntity.Gender,
			"birth_date": userEntity.BirthDate,
			"silhouette": userEntity.Silhouette,
			"height":     map[bool]*int{true: &userEntity.Height, false: nil}[userEntity.Height > 0],
			"weight":     map[bool]*int{true: &userEntity.Weight, false: nil}[userEntity.Weight > 0],
			"genres":     userEntity.Genres,
			"complexes":  userEntity.Complexes,
		},
		"categories":        userEntity.Categories,
		"login_date":        now,
		"login_count":       1,
		"registration_date": now,
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

// DeleteUser ユーザの削除
func DeleteUser(userID string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: userID}})
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

// UpdateUser ユーザの更新
func UpdateUser(userID string, userEntity userModel.UserRegisterEntity) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: userID}}
		update := bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "user_setting_id", Value: userEntity.UserSettingID},
				{Key: "name", Value: userEntity.Name},
				{Key: "categories", Value: userEntity.Categories},
				{Key: "icon_url", Value: userEntity.IconURL},
				{Key: "attribute", Value: bson.D{
					{Key: "gender", Value: userEntity.Gender},
					{Key: "birth_date", Value: userEntity.BirthDate},
					{Key: "silhouette", Value: userEntity.Silhouette},
					{Key: "height", Value: userEntity.Height},
					{Key: "weight", Value: userEntity.Weight},
					{Key: "genres", Value: userEntity.Genres},
					{Key: "complexes", Value: userEntity.Complexes},
				},
				},
			}},
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
