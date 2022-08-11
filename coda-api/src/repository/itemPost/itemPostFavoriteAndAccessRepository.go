package itemPostRepository

import (
	"context"

	itemPostModel "coda-api/src/model/item"
	"coda-api/src/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetStatusItemPostFavorite 投稿IDとユーザIDを指定していいねの状況を取得
func GetStatusItemPostFavorite(postID string, userID string) (*itemPostModel.StatusUserFavoriteItemPost, error) {
	return GetStatusPostFavorite(postID, userID, "item_post")
}

// AddFavoriteItemPost いいねを追加
func AddFavoriteItemPost(postID string, userID string) error {
	return AddFavoritePost(postID, userID, "item_post")
}

// DeleteFavoriteItemPost いいねを削除
func DeleteFavoriteItemPost(postID string, userID string) error {
	return DeleteFavoritePost(postID, userID, "item_post")
}

// DeleteFavoriteItemPostUserAll 指定したユーザのいいねを全て削除
func DeleteFavoriteItemPostUserAll(userID string) error {
	return DeleteFavoritePostUserAll(userID, "item_post")
}

// AddImpressionItemPosts インプレッションを追加
func AddImpressionItemPosts(itemPostIDs []string, userID string) {
	itemPosts, _ := getItemForImpressionOrClick(itemPostIDs)
	var operations []mongo.WriteModel
	for _, r := range itemPosts {
		if r.UserID != userID {
			operation := mongo.NewUpdateOneModel()
			operation.SetFilter(bson.D{{Key: "_id", Value: r.ID}})
			operation.SetUpdate(bson.D{{Key: "$set",
				Value: bson.D{{
					Key: "impression_count", Value: r.ImpressionCount + 1,
				}}}},
			)
			operation.SetUpsert(false)
			operations = append(operations, operation)
		}
	}
	if operations != nil {
		f := func(col *mongo.Collection) ([]bson.M, error) {
			bulkOption := options.BulkWriteOptions{}
			bulkOption.SetOrdered(true)
			_, err := col.BulkWrite(context.Background(), operations, &bulkOption)
			if err != nil {
				return nil, err
			}
			return nil, nil
		}
		_, _ = util.ExecuteDbQuery(f, "item_post")
	}
}

// AddClickItemPost クリックを追加
func AddClickItemPost(itemPostID string, userID string) {
	itemPosts, _ := getItemForImpressionOrClick([]string{itemPostID})
	if itemPosts != nil {
		itemPost := itemPosts[0]
		if itemPost.UserID != userID {
			f := func(col *mongo.Collection) ([]bson.M, error) {
				filter := bson.D{{Key: "_id", Value: itemPost.ID}}
				update := bson.D{{Key: "$set",
					Value: bson.D{
						{Key: "click_count", Value: itemPost.ClickCount + 1},
					}},
				}
				_, err := col.UpdateOne(context.Background(), filter, update)
				if err != nil {
					return nil, err
				}
				return nil, nil
			}
			_, _ = util.ExecuteDbQuery(f, "item_post")
		}
	}
}

// getItemForImpressionOrClick クリック更新用にアイテム投稿IDを指定して取得
func getItemForImpressionOrClick(itemPostIDs []string) ([]itemPostModel.ItemPostResponse, error) {
	postIDBsonA := bson.A{}
	for _, itemPostID := range itemPostIDs {
		postIDBsonA = append(postIDBsonA, itemPostID)
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "$in", Value: postIDBsonA}}}},
		}}
		matchStatusStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		var cur *mongo.Cursor
		var err error
		projectStage := getProjectQueryMyItemPost()
		cur, err = col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage, matchStatusStage, lookUpStage, unwindStage, projectStage,
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
	queryResult, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil || queryResult == nil {
		return nil, err
	}
	var itemPosts []itemPostModel.ItemPostResponse
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			itemPosts = append(itemPosts, *response)
		}
	}
	return itemPosts, err
}
