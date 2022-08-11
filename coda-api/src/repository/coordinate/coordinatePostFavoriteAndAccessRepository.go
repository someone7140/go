package coordinateRepository

import (
	"context"

	coordinateModel "coda-api/src/model/coordinate"
	itemPostModel "coda-api/src/model/item"
	itemPostRepository "coda-api/src/repository/itemPost"
	"coda-api/src/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetStatusCoordinatePostFavorite 投稿IDとユーザIDを指定していいねの状況を取得
func GetStatusCoordinatePostFavorite(postID string, userID string) (*itemPostModel.StatusUserFavoriteItemPost, error) {
	return itemPostRepository.GetStatusPostFavorite(postID, userID, "coordinate_post")
}

// AddFavoriteCoordinatePost いいねを追加
func AddFavoriteCoordinatePost(postID string, userID string) error {
	return itemPostRepository.AddFavoritePost(postID, userID, "coordinate_post")
}

// DeleteFavoriteCoordinatePost いいねを削除
func DeleteFavoriteCoordinatePost(postID string, userID string) error {
	return itemPostRepository.DeleteFavoritePost(postID, userID, "coordinate_post")
}

// DeleteFavoriteCoordinatePostUserAll 指定したユーザのいいねを全て削除
func DeleteFavoriteCoordinatePostUserAll(userID string) error {
	return itemPostRepository.DeleteFavoritePostUserAll(userID, "coordinate_post")
}

// AddImpressionCoordinatePosts インプレッションを追加
func AddImpressionCoordinatePosts(coordinatePostIDs []string, userID string) {
	coordinatePosts, _ := getCoordinateForImpressionOrClick(coordinatePostIDs)
	var operations []mongo.WriteModel
	for _, r := range coordinatePosts {
		if r.PostUserID != userID {
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
		_, _ = util.ExecuteDbQuery(f, "coordinate_post")
	}
}

// AddPurchaseRequestCount 購入申請の押された回数を追加
func AddPurchaseRequestCount(coordinatePostID string, userID string) {
	coordinatePosts, _ := getCoordinateForImpressionOrClick([]string{coordinatePostID})
	if coordinatePosts != nil {
		coordinatePost := coordinatePosts[0]
		if coordinatePost.PostUserID != userID {
			f := func(col *mongo.Collection) ([]bson.M, error) {
				filter := bson.D{{Key: "_id", Value: coordinatePost.ID}}
				update := bson.D{{Key: "$set",
					Value: bson.D{
						{Key: "purchase_request_count", Value: coordinatePost.PurchaseRequestCount + 1},
					}},
				}
				_, err := col.UpdateOne(context.Background(), filter, update)
				if err != nil {
					return nil, err
				}
				return nil, nil
			}
			_, _ = util.ExecuteDbQuery(f, "coordinate_post")
		}
	}
}

// AddClickCoordinatePost 購入申請を追加
func AddClickCoordinatePost(coordinatePostID string, userID string) {
	coordinatePosts, _ := getCoordinateForImpressionOrClick([]string{coordinatePostID})
	if coordinatePosts != nil {
		coordinatePost := coordinatePosts[0]
		if coordinatePost.PostUserID != userID {
			f := func(col *mongo.Collection) ([]bson.M, error) {
				filter := bson.D{{Key: "_id", Value: coordinatePost.ID}}
				update := bson.D{{Key: "$set",
					Value: bson.D{
						{Key: "click_count", Value: coordinatePost.ClickCount + 1},
					}},
				}
				_, err := col.UpdateOne(context.Background(), filter, update)
				if err != nil {
					return nil, err
				}
				return nil, nil
			}
			_, _ = util.ExecuteDbQuery(f, "coordinate_post")
		}
	}
}

// getCoordinateForImpressionOrClick クリック更新用にコーデ投稿IDを指定して取得
func getCoordinateForImpressionOrClick(coordinatePostIDs []string) ([]coordinateModel.CoordinatePostInfo, error) {
	postIDBsonA := bson.A{}
	for _, coordinatePostID := range coordinatePostIDs {
		postIDBsonA = append(postIDBsonA, coordinatePostID)
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "_id", Value: bson.D{{Key: "$in", Value: postIDBsonA}}}},
		}}
		matchStatusStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		var cur *mongo.Cursor
		var err error
		projectStage := getProjectQueryMyCoordinatePost()
		lookUpStage := GetLookupShopQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$shop"}}
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
	queryResult, err := util.ExecuteDbQuery(f, "coordinate_post")
	if err != nil || queryResult == nil {
		return nil, err
	}
	var coordinatePosts []coordinateModel.CoordinatePostInfo
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			coordinatePosts = append(coordinatePosts, *response)
		}
	}
	return coordinatePosts, err
}
