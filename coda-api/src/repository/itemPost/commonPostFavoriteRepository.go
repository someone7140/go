package itemPostRepository

import (
	"context"
	"reflect"
	"time"

	itemPostModel "coda-api/src/model/item"
	"coda-api/src/util"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetStatusPostFavorite 投稿IDとユーザIDを指定していいねの状況を取得
func GetStatusPostFavorite(postID string, userID string, colName string) (*itemPostModel.StatusUserFavoriteItemPost, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "_id", Value: postID}},
		}}
		eqBsonA := bson.A{}
		eqBsonA = append(eqBsonA, "$$favorite_users.user_id")
		eqBsonA = append(eqBsonA, userID)
		favoriteFilter := bson.D{{Key: "$filter", Value: bson.D{
			{Key: "input", Value: "$favorite_users"},
			{Key: "as", Value: "favorite_users"},
			{Key: "cond", Value: bson.D{{Key: "$eq", Value: eqBsonA}}}}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "post_user_id", Value: 1},
			{Key: "favorite_users", Value: favoriteFilter},
		}}}
		cur, err := col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage, projectStage,
			})
		if err != nil {
			return nil, err
		}
		var docs []bson.M
		for cur.Next(context.Background()) {
			var doc bson.M
			if err = cur.Decode(&doc); err != nil {
				return nil, err
			}
			docs = append(docs, doc)
		}
		return docs, nil
	}
	queryResult, err := util.ExecuteDbQuery(f, colName)
	if err != nil {
		return nil, err
	}
	if len(queryResult) == 0 {
		return nil, nil
	}
	v := queryResult[0]
	id, errGet := dproxy.New(v).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	userId, errGet := dproxy.New(v).M("post_user_id").String()
	if errGet != nil {
		return nil, errGet
	}
	favoriteUsersLen := 0
	if v["favorite_users"] != nil {
		favoriteUsersLen = reflect.ValueOf(v["favorite_users"].(primitive.A)).Len()
	}
	status := map[bool]string{true: "registered", false: "noregistered"}[favoriteUsersLen > 0]
	return &itemPostModel.StatusUserFavoriteItemPost{
		PostID: id,
		UserID: userId,
		Status: status,
	}, nil
}

// AddFavoritePost いいねを追加
func AddFavoritePost(postID string, userID string, colName string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: postID}}
		update := bson.D{{Key: "$push",
			Value: bson.D{
				{Key: "favorite_users", Value: bson.D{
					{Key: "user_id", Value: userID},
					{Key: "favorite_date", Value: time.Now().UTC()},
				}},
			}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, colName)
	return err
}

// DeleteFavoritePost いいねを削除
func DeleteFavoritePost(postID string, userID string, colName string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: postID}}
		update := bson.D{{Key: "$pull",
			Value: bson.D{
				{Key: "favorite_users", Value: bson.D{
					{Key: "user_id", Value: userID},
				}},
			}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, colName)
	return err
}

// DeleteFavoritePostUserAll 指定したユーザのいいねを全て削除
func DeleteFavoritePostUserAll(userID string, colName string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "favorite_users.user_id", Value: userID}}
		update := bson.D{{Key: "$pull",
			Value: bson.D{
				{Key: "favorite_users", Value: bson.D{
					{Key: "user_id", Value: userID},
				}},
			}},
		}
		_, err := col.UpdateMany(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, colName)
	return err
}
