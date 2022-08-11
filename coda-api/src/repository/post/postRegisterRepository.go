package postRepository

import (
	instagramAccountModel "coda-api/src/model/instagramAccount"
	postModel "coda-api/src/model/post"
	"context"

	"coda-api/src/util"

	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// RegisterPosts 引数の投稿データをDBに登録
func RegisterPosts(posts []instagramAccountModel.InstagramPost, userID string) error {
	registerBson := funk.Map(posts, func(p instagramAccountModel.InstagramPost) interface{} {
		return bson.M{
			"_id":                       p.ID,
			"post_type":                 "instagram_gather",
			"status":                    "notset",
			"post_instagram_account_id": userID,
			"post_date":                 p.PostDate,
			"content_url":               p.Permalink,
		}
	}).([]interface{})
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.InsertMany(context.Background(), registerBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "post")
	return err
}

// SetStatusPosts 投稿に対するステータス変更
func SetStatusPosts(request []postModel.PostStatusUpdateRequest) error {
	var operations []mongo.WriteModel
	for _, r := range request {
		operation := mongo.NewUpdateOneModel()
		operation.SetFilter(bson.D{{Key: "_id", Value: r.PostID}})
		operation.SetUpdate(bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "status", Value: r.Status},
				{Key: "genre", Value: map[bool]*string{true: nil, false: &r.Genre}[r.Genre == ""]},
			}}},
		)
		operation.SetUpsert(false)
		operations = append(operations, operation)
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		bulkOption := options.BulkWriteOptions{}
		bulkOption.SetOrdered(true)
		_, err := col.BulkWrite(context.Background(), operations, &bulkOption)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "post")
	return err
}
