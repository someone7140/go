package postRepository

import (
	"context"
	"reflect"
	"sort"
	"strings"
	"time"

	analyticsModel "coda-api/src/model/analytics"
	postModel "coda-api/src/model/post"
	"coda-api/src/util"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetStatusFavorite 投稿IDとユーザIDを指定していいねの状況を取得
func GetStatusFavorite(postID string, userID string) (*postModel.StatusUserFavorite, error) {
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
	queryResult, err := util.ExecuteDbQuery(f, "post")
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
	favoriteUsersLen := 0
	if v["favorite_users"] != nil {
		favoriteUsersLen = reflect.ValueOf(v["favorite_users"].(primitive.A)).Len()
	}
	status := map[bool]string{true: "registered", false: "noregistered"}[favoriteUsersLen > 0]
	return &postModel.StatusUserFavorite{
		PostID: id,
		Status: status,
	}, nil
}

// AddFavorite いいねを追加
func AddFavorite(postID string, userID string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) { // 収集した日を更新＆エラー日付を消去
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
	_, err := util.ExecuteDbQuery(f, "post")
	return err
}

// DeleteFavorite いいねを削除
func DeleteFavorite(postID string, userID string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) { // 収集した日を更新＆エラー日付を消去
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
	_, err := util.ExecuteDbQuery(f, "post")
	return err
}

// DeleteFavoriteUseAll 指定したユーザのいいねを全て削除
func DeleteFavoriteUseAll(userID string) error {
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
	_, err := util.ExecuteDbQuery(f, "post")
	return err
}

// GetFavoriteAnalytics いいねの分析データを取得
func GetFavoriteAnalytics(sortInput string, limit int64) ([]analyticsModel.FavoriteAnalytics, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "favorite_users", Value: bson.D{
				{Key: "$exists", Value: true}}}},
		}}
		lookupInstagramAccountStage := GetLookupInstagramAccountQuery()
		unwindInstagramAccountStage := bson.D{{Key: "$unwind", Value: "$instagram_gather_account"}}
		lookupFavoriteUserStage := bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "user"},
			{Key: "localField", Value: "favorite_users.user_id"},
			{Key: "foreignField", Value: "_id"},
			{Key: "as", Value: "users"},
		}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "status", Value: 1},
			{Key: "content_url", Value: 1},
			{Key: "post_date", Value: 1},
			{Key: "genre", Value: 1},
			{Key: "post_instagram_account_id", Value: 1},
			{Key: "instagram_gather_account.instagram_user_name", Value: 1},
			{Key: "instagram_gather_account.category", Value: 1},
			{Key: "favorite_count", Value: GetFavoriteCountQuery()},
			{Key: "users", Value: 1},
		}}}
		matchStageFavoriteCount := bson.D{{Key: "$match", Value: bson.D{
			{Key: "favorite_count", Value: bson.D{
				{Key: "$gte", Value: 1}}}},
		}}
		sortStage := map[bool]bson.D{
			true: {{Key: "$sort", Value: bson.D{
				{Key: "post_date", Value: -1},
				{Key: "favorite_count", Value: -1},
			}}}, false: {{Key: "$sort", Value: bson.D{
				{Key: "favorite_count", Value: -1},
				{Key: "post_date", Value: -1},
			}}},
		}[sortInput == "recent"]
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		cur, err := col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage,
				lookupInstagramAccountStage,
				unwindInstagramAccountStage,
				lookupFavoriteUserStage,
				projectStage,
				matchStageFavoriteCount,
				sortStage,
				limitStage,
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
	queryResult, err := util.ExecuteDbQuery(f, "post")
	if err != nil {
		return nil, err
	}
	if len(queryResult) == 0 {
		return nil, nil
	}
	var posts []analyticsModel.FavoriteAnalytics
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		status, errGet := dproxy.New(v).M("status").String()
		if errGet != nil {
			return nil, errGet
		}
		postDateInt := int64(v["post_date"].(primitive.DateTime))
		genre, _ := dproxy.New(v).M("genre").String()
		contentUrl, errGet := dproxy.New(v).M("content_url").String()
		if errGet != nil {
			return nil, errGet
		}
		postInstagramAccountId, errGet := dproxy.New(v).M("post_instagram_account_id").String()
		if errGet != nil {
			return nil, errGet
		}
		postInstagramUserName, errGet := dproxy.New(v).M("instagram_gather_account").M("instagram_user_name").String()
		if errGet != nil {
			return nil, errGet
		}
		category, errGet := dproxy.New(v).M("instagram_gather_account").M("category").String()
		if errGet != nil {
			return nil, errGet
		}
		favoriteCount, errGet := dproxy.New(v).M("favorite_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		var favoriteDetails []analyticsModel.FavoriteAnalyticsDetail
		var detailKeyCount map[string]int = map[string]int{}
		if v["users"] != nil {
			users := v["users"].(primitive.A)
			for i := 0; i < reflect.ValueOf(users).Len(); i++ {
				attribute := users[i].(primitive.M)["attribute"].(primitive.M)
				gender, _ := dproxy.New(attribute).M("gender").String()
				silhouette, _ := dproxy.New(attribute).M("silhouette").String()
				height, _ := dproxy.New(attribute).M("height").Int64()
				heightCategory := map[bool]string{
					true: util.GetHeightCategory(int(height), gender), false: "none",
				}[height > 0]
				genreCategory := "none"
				if attribute["genres"] != nil {
					var genres []string
					genresInput := attribute["genres"].(primitive.A)
					for i := 0; i < reflect.ValueOf(genresInput).Len(); i++ {
						genre, errGet := dproxy.New(genresInput[i]).String()
						if errGet != nil {
							return nil, errGet
						}
						genres = append(genres, genre)
					}
					sort.Strings(genres)
					genreCategory = strings.Join(genres, "|")
				}
				complexCategory := "none"
				if attribute["complexes"] != nil {
					var complexes []string
					complexessInput := attribute["complexes"].(primitive.A)
					for i := 0; i < reflect.ValueOf(complexessInput).Len(); i++ {
						complex, errGet := dproxy.New(complexessInput[i]).String()
						if errGet != nil {
							return nil, errGet
						}
						complexes = append(complexes, complex)
					}
					sort.Strings(complexes)
					complexCategory = strings.Join(complexes, "|")
				}
				detailKey := gender + "-" + silhouette + "-" + heightCategory + "-" +
					genreCategory + "-" + complexCategory
				if val, ok := detailKeyCount[detailKey]; ok {
					detailKeyCount[detailKey] = val + 1
				} else {
					detailKeyCount[detailKey] = 1
				}
			}
			for key, value := range detailKeyCount {
				favoriteDetails = append(favoriteDetails,
					analyticsModel.FavoriteAnalyticsDetail{
						UserCategory:  key,
						FavoriteCount: value,
					},
				)
			}
		}
		posts = append(posts, analyticsModel.FavoriteAnalytics{
			PostID:                 id,
			Status:                 status,
			Category:               map[bool]string{true: category, false: util.GetPostCategoryMergePostGenre(category, genre)}[genre == ""],
			PostInstagramAccountId: postInstagramAccountId,
			PostInstagramUserName:  postInstagramUserName,
			ContentURL:             contentUrl,
			PostDate:               util.GetMongoDateIntToTime(postDateInt),
			FavoriteTotalCount:     favoriteCount,
			FavoriteDetails:        favoriteDetails,
		})
	}
	return posts, nil
}
