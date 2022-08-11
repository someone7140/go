package postRepository

import (
	postModel "coda-api/src/model/post"
	"context"

	"coda-api/src/util"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetPostsByIds idを指定したpostの検索
func GetPostsByIds(ids []string) ([]string, error) {
	idBsonA := bson.A{}
	for _, id := range ids {
		idBsonA = append(idBsonA, id)
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: bson.D{{Key: "$in", Value: idBsonA}}}}
		cur, err := col.Find(context.Background(), filter)
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
	var getIds []string
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		getIds = append(getIds, id)
	}
	return getIds, nil
}

// GetNotsetAllPosts ステータス未設定の投稿を取得
func GetNotsetAllPosts(limit int) ([]postModel.PostResponseForAdmin, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage1 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "notset"}},
		}}
		matchStage2 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "post_type", Value: "instagram_gather"}},
		}}
		lookupStage := GetLookupInstagramAccountQuery()
		projectStage := GetProjectInstagramAccountQuery("")
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: 1},
		}}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		cur, err := col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage1, matchStage2, lookupStage, projectStage, sortStage, limitStage,
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
	var posts []postModel.PostResponseForAdmin
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		status, errGet := dproxy.New(v).M("status").String()
		if errGet != nil {
			return nil, errGet
		}
		postInstagramAccountId, errGet := dproxy.New(v).M("post_instagram_account_id").String()
		if errGet != nil {
			return nil, errGet
		}
		postInstagramUserName, errGet := dproxy.New(v["instagram_gather_account"].(primitive.A)[0]).M("instagram_user_name").String()
		if errGet != nil {
			return nil, errGet
		}
		category, errGet := dproxy.New(v["instagram_gather_account"].(primitive.A)[0]).M("category").String()
		if errGet != nil {
			return nil, errGet
		}
		contentURL, errGet := dproxy.New(v).M("content_url").String()
		if errGet != nil {
			return nil, errGet
		}
		genre, _ := dproxy.New(v).M("genre").String()
		postDateInt := int64(v["post_date"].(primitive.DateTime))
		posts = append(posts, postModel.PostResponseForAdmin{
			ID:                     id,
			Status:                 status,
			PostInstagramAccountId: postInstagramAccountId,
			PostInstagramUserName:  postInstagramUserName,
			Category:               category,
			PostGenre:              genre,
			ContentURL:             contentURL,
			PostDate:               util.GetMongoDateIntToTime(postDateInt),
		})
	}
	return posts, nil
}

// GetRecentPosts 最新の投稿を取得
func GetRecentPosts(limit int, genre string) ([]postModel.PostResponseForRecommend, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
		lookupStage := GetLookupInstagramAccountQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$instagram_gather_account"}}
		projectStage := GetProjectInstagramAccountQuery("")
		if genre == "" {
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStage, lookupStage, unwindStage, projectStage, sortStage, limitStage,
				})
		} else {
			matchGenreStage := GetRegexGenreInstagramAccountQuery(genre)
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStage, lookupStage, unwindStage, projectStage, matchGenreStage, sortStage, limitStage,
				})
		}

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
	var posts []postModel.PostResponseForRecommend
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		contentURL, errGet := dproxy.New(v).M("content_url").String()
		if errGet != nil {
			return nil, errGet
		}
		getGenre, _ := dproxy.New(v).M("genre").String()
		category, errGet := dproxy.New(v).M("instagram_gather_account").M("category").String()
		if errGet != nil {
			return nil, errGet
		}
		postDateInt := int64(v["post_date"].(primitive.DateTime))
		favoriteCount, errGet := dproxy.New(v).M("favorite_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		posts = append(posts, postModel.PostResponseForRecommend{
			ID:            id,
			ContentURL:    contentURL,
			Genre:         map[bool]string{true: util.GetGenreFromCategory(category), false: getGenre}[getGenre == ""],
			PostDate:      util.GetMongoDateIntToTime(postDateInt),
			FavoriteCount: int(favoriteCount),
			FavoritedFlg:  false,
		})
	}
	return posts, nil
}

// GetCategoryMatchedPosts カテゴリーにマッチした投稿を取得
func GetCategoryMatchedPosts(
	limit int,
	targetCategories map[string]int,
	userID string,
	genre string) ([]postModel.PostResponseForRecommend, error) {
	// カテゴリー名のリストを抽出
	categoryNameBsonA := bson.A{}
	for categoryKey := range targetCategories {
		categoryNameBsonA = append(categoryNameBsonA, categoryKey)
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage1 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		lookupStage := GetLookupInstagramAccountQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$instagram_gather_account"}}
		projectStage := GetProjectInstagramAccountQuery(userID)
		matchCategoryStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "instagram_gather_account.category", Value: bson.D{{Key: "$in", Value: categoryNameBsonA}}}}}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		// limitの2倍をまずは取得
		limitStage := bson.D{{Key: "$limit", Value: limit * 2}}
		var cur *mongo.Cursor
		var err error
		if genre == "" {
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStage1, lookupStage, unwindStage, projectStage, matchCategoryStage, sortStage, limitStage,
				})
		} else {
			matchGenreStage := GetRegexGenreInstagramAccountQuery(genre)
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStage1, lookupStage, unwindStage, projectStage, matchGenreStage, matchCategoryStage, sortStage, limitStage,
				})
		}

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
	var postInfosWithCategory []postModel.PostInfoWithCategory
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		category, errGet := dproxy.New(v).M("instagram_gather_account").M("category").String()
		if errGet != nil {
			return nil, errGet
		}
		contentURL, errGet := dproxy.New(v).M("content_url").String()
		if errGet != nil {
			return nil, errGet
		}
		getGenre, _ := dproxy.New(v).M("genre").String()
		postDateInt := int64(v["post_date"].(primitive.DateTime))
		favoriteCount, errGet := dproxy.New(v).M("favorite_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		favoriteCountLoginUser, errGet := dproxy.New(v).M("favorite_count_login_user").Int64()
		if errGet != nil {
			return nil, errGet
		}
		postInfosWithCategory = append(postInfosWithCategory, postModel.PostInfoWithCategory{
			ID:            id,
			ContentURL:    contentURL,
			PostDate:      util.GetMongoDateIntToTime(postDateInt),
			Category:      map[bool]string{true: category, false: util.GetPostCategoryMergePostGenre(category, getGenre)}[getGenre == ""],
			FavoriteCount: int(favoriteCount),
			FavoritedFlg:  map[bool]bool{true: true, false: false}[favoriteCountLoginUser > 0],
		})
	}
	return util.GetPostResponseSortedByMatchePoint(postInfosWithCategory, targetCategories, limit), nil
}

// GetFavoritedPosts いいねをした投稿を取得
func GetFavoritedPosts(limit int, userID string) ([]postModel.PostResponseForRecommend, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage1 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		lookupStage := GetLookupInstagramAccountQuery()
		unwindInstagramAccountStage := bson.D{{Key: "$unwind", Value: "$instagram_gather_account"}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "content_url", Value: 1},
			{Key: "post_date", Value: 1},
			{Key: "genre", Value: 1},
			{Key: "favorite_count", Value: GetFavoriteCountQuery()},
			{Key: "favorite_users", Value: GetFavoriteDateQueryByUser(userID)},
			{Key: "instagram_gather_account.category", Value: 1},
		}}}
		unwindFavoriteStage := bson.D{{Key: "$unwind", Value: "$favorite_users"}}
		matchFavoriteStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "favorite_users.user_id", Value: userID}}}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "favorite_users.favorite_date", Value: -1},
		}}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		cur, err := col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage1, lookupStage, unwindInstagramAccountStage, projectStage, unwindFavoriteStage, matchFavoriteStage, sortStage, limitStage,
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
	var posts []postModel.PostResponseForRecommend
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		contentURL, errGet := dproxy.New(v).M("content_url").String()
		if errGet != nil {
			return nil, errGet
		}
		category, errGet := dproxy.New(v).M("instagram_gather_account").M("category").String()
		if errGet != nil {
			return nil, errGet
		}
		genre, _ := dproxy.New(v).M("genre").String()
		postDateInt := int64(v["post_date"].(primitive.DateTime))
		favoriteCount, errGet := dproxy.New(v).M("favorite_count").Int64()
		if errGet != nil {
			return nil, errGet
		}
		posts = append(posts, postModel.PostResponseForRecommend{
			ID:            id,
			ContentURL:    contentURL,
			Genre:         map[bool]string{true: util.GetGenreFromCategory(category), false: genre}[genre == ""],
			PostDate:      util.GetMongoDateIntToTime(postDateInt),
			FavoriteCount: int(favoriteCount),
			FavoritedFlg:  true,
		})
	}
	return posts, nil
}

// GetFavoriteCountQuery いいね数をカウントするクエリ定義
func GetFavoriteCountQuery() bson.D {
	condBsonA := bson.A{}
	condNotBsonA := bson.A{}
	condNotBsonA = append(condNotBsonA, "$favorite_users")
	condBsonA = append(condBsonA, bson.D{{Key: "$not", Value: condNotBsonA}})
	condBsonA = append(condBsonA, 0)
	condBsonA = append(condBsonA, bson.D{{Key: "$size", Value: "$favorite_users"}})
	return bson.D{{Key: "$cond", Value: condBsonA}}
}

// GetFavoriteCountQueryByUser ユーザを指定していいね数をカウントするクエリ定義
func GetFavoriteCountQueryByUser(userID string) bson.D {
	eqBsonA := bson.A{}
	eqBsonA = append(eqBsonA, "$$favorite_users.user_id")
	eqBsonA = append(eqBsonA, userID)
	filterBsonD := bson.D{{Key: "$filter", Value: bson.D{
		{Key: "input", Value: "$favorite_users"},
		{Key: "as", Value: "favorite_users"},
		{Key: "cond", Value: bson.D{
			{Key: "$eq", Value: eqBsonA},
		},
		}}}}
	condBsonA := bson.A{}
	condNotBsonA := bson.A{}
	condNotBsonA = append(condNotBsonA, "$favorite_users")
	condBsonA = append(condBsonA, bson.D{{Key: "$not", Value: condNotBsonA}})
	condBsonA = append(condBsonA, 0)
	condBsonA = append(condBsonA, bson.D{{Key: "$size", Value: filterBsonD}})
	return bson.D{{Key: "$cond", Value: condBsonA}}
}

// GetFavoriteDateQueryByUser ユーザを指定していいね日付を取得するクエリ定義
func GetFavoriteDateQueryByUser(userID string) bson.D {
	eqBsonA := bson.A{}
	eqBsonA = append(eqBsonA, "$$favorite_users.user_id")
	eqBsonA = append(eqBsonA, userID)
	filterBsonD := bson.D{{Key: "$filter", Value: bson.D{
		{Key: "input", Value: "$favorite_users"},
		{Key: "as", Value: "favorite_users"},
		{Key: "cond", Value: bson.D{
			{Key: "$eq", Value: eqBsonA},
		},
		}}}}
	return filterBsonD
}

// GetLookupInstagramAccountQuery instagram_gather_accountと結合するためのクエリ
func GetLookupInstagramAccountQuery() bson.D {
	return bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "instagram_gather_account"},
		{Key: "localField", Value: "post_instagram_account_id"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "instagram_gather_account"},
	}}}
}

// GetProjectInstagramAccountQuery instagram_gather_accountと結合して列指定をするためのクエリ
func GetProjectInstagramAccountQuery(userID string) bson.D {
	if userID != "" {
		return bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "status", Value: 1},
			{Key: "post_instagram_account_id", Value: 1},
			{Key: "content_url", Value: 1},
			{Key: "post_date", Value: 1},
			{Key: "genre", Value: 1},
			{Key: "instagram_gather_account.instagram_user_name", Value: 1},
			{Key: "instagram_gather_account.category", Value: 1},
			{Key: "favorite_count", Value: GetFavoriteCountQuery()},
			{Key: "favorite_count_login_user", Value: GetFavoriteCountQueryByUser(userID)},
		}}}
	}
	return bson.D{{Key: "$project", Value: bson.D{
		{Key: "_id", Value: 1},
		{Key: "status", Value: 1},
		{Key: "post_instagram_account_id", Value: 1},
		{Key: "content_url", Value: 1},
		{Key: "post_date", Value: 1},
		{Key: "genre", Value: 1},
		{Key: "instagram_gather_account.instagram_user_name", Value: 1},
		{Key: "instagram_gather_account.category", Value: 1},
		{Key: "favorite_count", Value: GetFavoriteCountQuery()},
	}}}
}

// GetRegexGenreInstagramAccountQuery instagram_gather_accountからジャンルを指定するクエリ
func GetRegexGenreInstagramAccountQuery(genre string) bson.D {
	return bson.D{{Key: "$match", Value: bson.D{
		{Key: "$or", Value: bson.A{
			bson.M{"$and": bson.A{
				bson.M{"genre": bson.D{{Key: "$ne", Value: nil}}},
				bson.M{"genre": genre}}},
			bson.M{"$and": bson.A{
				bson.M{"genre": nil},
				bson.M{"instagram_gather_account.category": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: genre, Options: "i"}}}}},
			},
		}},
	}},
	}
}
