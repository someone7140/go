package analyticsRepository

import (
	"coda-api/src/util"
	"context"

	analyticsModel "coda-api/src/model/analytics"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddDateAccessAnalytics 1日単位のアクセスデータの追加
func AddDateAccessAnalytics(
	dateKey string,
	totalUserCount int64,
	loginUserCount int64,
	accessUserCount int64,
	instagramFollowerCount int64,
	twitterFollowerCount int64,
) error {
	registerBson := bson.M{
		"_id":                      dateKey,
		"total_user_count":         totalUserCount,
		"login_user_count":         loginUserCount,
		"access_user_count":        accessUserCount,
		"instagram_follower_count": instagramFollowerCount,
		"twitter_follower_count":   twitterFollowerCount,
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.InsertOne(context.Background(), registerBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "analytics_report")
	if err != nil {
		return err
	}
	return nil
}

// GetDateAccessAnalytics 1日単位のアクセスデータの取得
func GetDateAccessAnalytics(limit int) ([]analyticsModel.AccessAnalytics, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "total_user_count", Value: 1},
			{Key: "login_user_count", Value: 1},
			{Key: "access_user_count", Value: 1},
			{Key: "instagram_follower_count", Value: 1},
			{Key: "twitter_follower_count", Value: 1},
		}}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "_id", Value: -1},
		}}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		cur, err := col.Aggregate(context.Background(),
			mongo.Pipeline{
				projectStage, sortStage, limitStage,
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
	queryResult, err := util.ExecuteDbQuery(f, "analytics_report")
	if err != nil {
		return nil, err
	}
	var accessAnalyticsArray []analyticsModel.AccessAnalytics
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		totalUserCount, errGet := dproxy.New(v).M("total_user_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		loginUserCount, errGet := dproxy.New(v).M("login_user_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		accessUserCount, errGet := dproxy.New(v).M("access_user_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		instagramFollowerCount, errGet := dproxy.New(v).M("instagram_follower_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		twitterFollowerCount, errGet := dproxy.New(v).M("twitter_follower_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		accessAnalyticsArray = append(accessAnalyticsArray, analyticsModel.AccessAnalytics{
			ID:                     id,
			TotalUserCount:         totalUserCount,
			LoginUserCount:         loginUserCount,
			AccessUserCount:        accessUserCount,
			InstagramFollowerCount: instagramFollowerCount,
			TwitterFollowerCount:   twitterFollowerCount,
		})
	}
	return accessAnalyticsArray, nil
}
