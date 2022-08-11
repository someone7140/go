package itemPostRepository

import (
	authModel "coda-api/src/model/auth"
	itemPostModel "coda-api/src/model/item"
	"context"
	"time"

	"coda-api/src/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddItemPost アイテム投稿データをDB登録
func AddItemPost(
	uuid string,
	itemPost itemPostModel.ItemPostRequest,
	userID string,
	imageURL string,
) error {
	now := time.Now().UTC()
	registerBson := bson.M{
		"_id":              uuid,
		"title":            itemPost.Title,
		"item_type":        itemPost.ItemType,
		"detail":           itemPost.Detail,
		"url":              itemPost.URL,
		"status":           itemPost.Status,
		"post_user_id":     userID,
		"post_date":        now,
		"impression_count": 0,
		"click_count":      0,
		"category":         util.GetItemCategoryFromAttributeInput(itemPost.Gender, itemPost.Silhouette, itemPost.Complex),
		"attribute": bson.M{
			"gender":     itemPost.Gender,
			"silhouette": itemPost.Silhouette,
			"complex":    itemPost.Complex,
		},
		"image_url": imageURL,
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.InsertOne(context.Background(), registerBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "item_post")
	return err
}

// UpdateItemPost アイテム投稿データの更新
func UpdateItemPost(
	updateItemPost itemPostModel.ItemPostUpdateRequest,
	user *authModel.AuthUserResponse,
	imageURL string) error {
	filter := map[bool]bson.D{
		true: {
			{Key: "_id", Value: updateItemPost.ID},
		},
		false: {
			{Key: "_id", Value: updateItemPost.ID},
			{Key: "post_user_id", Value: user.ID},
		},
	}[user.UserType == "admin"]
	updateBson := bson.D{{Key: "$set",
		Value: bson.D{
			{Key: "title", Value: updateItemPost.Title},
			{Key: "item_type", Value: updateItemPost.ItemType},
			{Key: "detail", Value: updateItemPost.Detail},
			{Key: "url", Value: updateItemPost.URL},
			{Key: "status", Value: updateItemPost.Status},
			{Key: "category", Value: util.GetItemCategoryFromAttributeInput(
				updateItemPost.Gender, updateItemPost.Silhouette, updateItemPost.Complex)},
			{Key: "attribute", Value: bson.D{
				{Key: "gender", Value: updateItemPost.Gender},
				{Key: "silhouette", Value: updateItemPost.Silhouette},
				{Key: "complex", Value: updateItemPost.Complex},
			}},
			{Key: "image_url", Value: imageURL},
		}}}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.UpdateOne(context.Background(), filter, updateBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "item_post")
	return err
}

// DeleteItemPost アイテム投稿の削除
func DeleteItemPost(postID string, user *authModel.AuthUserResponse) error {
	filter := map[bool]bson.D{
		true: {
			{Key: "_id", Value: postID},
		},
		false: {
			{Key: "_id", Value: postID},
			{Key: "post_user_id", Value: user.ID},
		},
	}[user.UserType == "admin"]
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.DeleteOne(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllItemPostsByUser ユーザの全アイテム投稿を削除
func DeleteAllItemPostsByUser(userID string) error {
	filter := bson.D{{Key: "post_user_id", Value: userID}}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.DeleteMany(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return err
	}
	return nil
}
