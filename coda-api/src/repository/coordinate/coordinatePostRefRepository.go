package coordinateRepository

import (
	"context"
	"math"
	"reflect"
	"time"

	errorConstants "coda-api/src/constants"
	analyticsModel "coda-api/src/model/analytics"
	coordinateModel "coda-api/src/model/coordinate"
	postRepository "coda-api/src/repository/post"
	"coda-api/src/util"

	"github.com/koron/go-dproxy"
	"github.com/thoas/go-funk"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetCoordinatePostByID コーデ投稿IDを指定して取得
func GetCoordinatePostByID(postID string, myUserID string, myPostFlag bool) (*coordinateModel.CoordinatePostInfo, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "_id", Value: postID}},
		}}
		matchStageStatus := map[bool]bson.D{
			true: {{Key: "$match", Value: bson.M{"status": bson.D{{Key: "$ne", Value: nil}}}}},
			false: {{Key: "$match", Value: bson.D{
				{Key: "status", Value: "open"}},
			}}}[myPostFlag]
		lookUpStage := GetLookupShopQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$shop"}}
		var cur *mongo.Cursor
		var err error
		projectStage := map[bool]bson.D{
			true:  getProjectQueryMyCoordinatePost(),
			false: getProjectQueryOtherCoordinatePost(myUserID),
		}[myPostFlag]
		cur, err = col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage, matchStageStatus, lookUpStage, unwindStage, projectStage,
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
	res, err := getQueryResultToResponse(queryResult[0])
	if err != nil {
		return nil, err
	}
	if myPostFlag && myUserID != res.PostUserID {
		return nil, errorConstants.ErrForbidden
	}
	return res, err
}

// GetFavoritedCoordinatePosts いいねをしたコーデ投稿を取得
func GetFavoritedCoordinatePosts(limit int, userID string) ([]coordinateModel.CoordinatePostInfo, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage1 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		lookUpStage := GetLookupShopQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$shop"}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "title", Value: 1},
			{Key: "status", Value: 1},
			{Key: "shop_id", Value: 1},
			{Key: "shop.name", Value: 1},
			{Key: "shop.shop_setting_id", Value: 1},
			{Key: "detail", Value: 1},
			{Key: "url", Value: 1},
			{Key: "post_date", Value: 1},
			{Key: "impression_count", Value: 1},
			{Key: "click_count", Value: 1},
			{Key: "model_category", Value: 1},
			{Key: "model_attribute.gender", Value: 1},
			{Key: "model_attribute.silhouette", Value: 1},
			{Key: "model_attribute.height", Value: 1},
			{Key: "model_attribute.weight", Value: 1},
			{Key: "images.key", Value: 1},
			{Key: "images.url", Value: 1},
			{Key: "favorite_count", Value: postRepository.GetFavoriteCountQuery()},
			{Key: "favorite_count_login_user", Value: postRepository.GetFavoriteCountQueryByUser(userID)},
			{Key: "favorite_users", Value: postRepository.GetFavoriteDateQueryByUser(userID)},
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
				matchStage1, lookUpStage, unwindStage, projectStage, unwindFavoriteStage, matchFavoriteStage, sortStage, limitStage,
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
	if err != nil {
		return nil, err
	}
	var coordinatePosts = []coordinateModel.CoordinatePostInfo{}

	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			coordinatePosts = append(coordinatePosts, *response)
		}
	}
	return coordinatePosts, err
}

// GetRecentCoordinatePosts 最新のコーデ投稿を取得
func GetRecentCoordinatePosts(limit int, allFlag bool, myUserID string) ([]coordinateModel.CoordinatePostInfo, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStageStatus := map[bool]bson.D{
			true: {{Key: "$match", Value: bson.M{"status": bson.D{{Key: "$ne", Value: nil}}}}},
			false: {{Key: "$match", Value: bson.D{
				{Key: "status", Value: "open"}},
			}}}[allFlag]
		lookUpStage := GetLookupShopQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$shop"}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
		projectStage := map[bool]bson.D{
			true:  getProjectQueryMyCoordinatePost(),
			false: getProjectQueryOtherCoordinatePost(myUserID),
		}[allFlag]
		cur, err = col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStageStatus, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
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
	if err != nil {
		return nil, err
	}
	var cordinatePosts = []coordinateModel.CoordinatePostInfo{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			cordinatePosts = append(cordinatePosts, *response)
		}
	}
	return cordinatePosts, err
}

// GetRecentCoordinateWithImagePosts 最新の画像があるコーデ投稿を取得
func GetRecentCoordinateWithImagePosts(limit int, myUserID string, saleOnly bool) ([]coordinateModel.CoordinatePostInfo, error) {
	matchArray := bson.A{}
	matchArray = append(matchArray, bson.M{"status": "open"})
	matchArray = append(matchArray, bson.M{"images.0": bson.D{
		{Key: "$exists", Value: true}},
	})
	if saleOnly {
		now := time.Now().UTC()
		saleFilter := bson.D{
			{Key: "$and", Value: bson.A{
				bson.M{"sale.sale_price": bson.M{"$gt": 0}},
				bson.M{"sale.start_date": bson.M{"$lte": now}},
				bson.M{"sale.end_date": bson.M{"$gte": now}},
			}},
		}
		matchArray = append(matchArray, saleFilter)
	}

	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "$and", Value: matchArray}}}}
		lookUpStage := GetLookupShopQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$shop"}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
		projectStage := getProjectQueryOtherCoordinatePost(myUserID)
		cur, err = col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
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
	if err != nil {
		return nil, err
	}
	var coordinatePosts = []coordinateModel.CoordinatePostInfo{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			coordinatePosts = append(coordinatePosts, *response)
		}
	}
	return coordinatePosts, err
}

// GetSearchCoordinatePosts コーデ投稿の検索
func GetSearchCoordinatePosts(
	limit int,
	allFlag bool,
	myUserID string,
	searchRequest coordinateModel.CoordinatePostSearchRequest,
) ([]coordinateModel.CoordinatePostInfo, error) {
	now := time.Now().UTC()
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStageStatus := map[bool]bson.D{
			true: {{Key: "$match", Value: bson.M{"status": bson.D{{Key: "$ne", Value: nil}}}}},
			false: {{Key: "$match", Value: bson.D{
				{Key: "status", Value: "open"}},
			}}}[allFlag]
		// 属性の検索条件
		searchAttributeBsonA := bson.A{}
		if searchRequest.Gender != "" {
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.gender": searchRequest.Gender})
		}
		if searchRequest.Silhouette != "" {
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.silhouette": searchRequest.Silhouette})
		}
		if searchRequest.MinHeight > 0 {
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.height": bson.M{"$gte": searchRequest.MinHeight}})
		}
		if searchRequest.MaxHeight > 0 {
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.height": bson.M{"$lte": searchRequest.MaxHeight}})
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.height": bson.M{"$ne": 0}})
		}
		if searchRequest.MinWeight > 0 {
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.weight": bson.M{"$gte": searchRequest.MinWeight}})
		}
		if searchRequest.MaxWeight > 0 {
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.weight": bson.M{"$lte": searchRequest.MaxWeight}})
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"model_attribute.weight": bson.M{"$ne": 0}})
		}
		if searchRequest.MinPrice > 0 {
			minPriceCondition := bson.D{{Key: "$or",
				Value: bson.A{
					bson.M{"price": bson.M{"$gte": searchRequest.MinPrice}},
					bson.M{"$and": bson.A{
						bson.M{"sale.sale_price": bson.M{"$gte": searchRequest.MinPrice}},
						bson.M{"sale.start_date": bson.M{"$lte": now}},
						bson.M{"sale.end_date": bson.M{"$gte": now}},
					}},
				}},
			}
			searchAttributeBsonA = append(searchAttributeBsonA, minPriceCondition)
		}
		if searchRequest.MaxPrice > 0 {
			maxPriceCondition := bson.D{{Key: "$or",
				Value: bson.A{
					bson.M{"$and": bson.A{
						bson.M{"price": bson.M{"$lte": searchRequest.MaxPrice}},
						bson.M{"price": bson.M{"$ne": 0}},
					}},
					bson.M{"$and": bson.A{
						bson.M{"sale.sale_price": bson.M{"$lte": searchRequest.MaxPrice}},
						bson.M{"sale.sale_price": bson.M{"$ne": 0}},
						bson.M{"sale.start_date": bson.M{"$lte": now}},
						bson.M{"sale.end_date": bson.M{"$gte": now}},
					}},
				}},
			}
			searchAttributeBsonA = append(searchAttributeBsonA, maxPriceCondition)
		}
		if searchRequest.Category != "" {
			searchAttributeBsonA = append(searchAttributeBsonA, bson.M{"category": searchRequest.Category})
		}
		// キーワードとショップ検索
		searchKeyWordAndShopBsonA := bson.A{}
		if searchRequest.Keyword != "" {
			searchKeyWordAndShopBsonA = append(searchKeyWordAndShopBsonA, bson.M{"$or": bson.A{
				bson.M{"title": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: searchRequest.Keyword, Options: "i"}}}},
				bson.M{"detail": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: searchRequest.Keyword, Options: "i"}}}},
				bson.M{"shop.name": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: searchRequest.Keyword, Options: "i"}}}},
			}})
		}
		if searchRequest.ShopSettingId != "" {
			searchKeyWordAndShopBsonA = append(searchKeyWordAndShopBsonA, bson.M{"shop.shop_setting_id": searchRequest.ShopSettingId})
		}
		// クエリの発行
		lookUpStage := GetLookupShopQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$shop"}}
		limitStage := bson.D{{
			Key: "$limit",
			Value: map[bool]int{
				true:  limit,
				false: math.MaxInt32,
			}[limit > 0],
		}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
		projectStage := map[bool]bson.D{
			true:  getProjectQueryMyCoordinatePost(),
			false: getProjectQueryOtherCoordinatePost(myUserID),
		}[allFlag]
		// 検索条件によってクエリの発行形式を変える
		if reflect.ValueOf(searchAttributeBsonA).Len() > 0 {
			matchStageSearchAttribute := bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: searchAttributeBsonA}}}}
			if reflect.ValueOf(searchKeyWordAndShopBsonA).Len() > 0 {
				matchStageSearchKeyWordAndShop := bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: searchKeyWordAndShopBsonA}}}}
				cur, err = col.Aggregate(context.Background(),
					mongo.Pipeline{
						matchStageStatus, matchStageSearchAttribute, lookUpStage, unwindStage, projectStage, matchStageSearchKeyWordAndShop, sortStage, limitStage,
					})
			} else {
				cur, err = col.Aggregate(context.Background(),
					mongo.Pipeline{
						matchStageStatus, matchStageSearchAttribute, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
					})
			}
		} else {
			if reflect.ValueOf(searchKeyWordAndShopBsonA).Len() > 0 {
				matchStageSearchKeyWordAndShop := bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: searchKeyWordAndShopBsonA}}}}
				cur, err = col.Aggregate(context.Background(),
					mongo.Pipeline{
						matchStageStatus, lookUpStage, unwindStage, projectStage, matchStageSearchKeyWordAndShop, sortStage, limitStage,
					})
			} else {
				cur, err = col.Aggregate(context.Background(),
					mongo.Pipeline{
						matchStageStatus, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
					})
			}

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
	queryResult, err := util.ExecuteDbQuery(f, "coordinate_post")
	if err != nil {
		return nil, err
	}
	var cordinatePosts = []coordinateModel.CoordinatePostInfo{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			cordinatePosts = append(cordinatePosts, *response)
		}
	}
	return cordinatePosts, err
}

// GetCoordinatePostAnalysis 分析用のコーデ投稿取得
func GetCoordinatePostAnalysis(analysisSpan string) ([]analyticsModel.CoordinatePostAnalytics, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		lookUpStage := GetLookupShopQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$shop"}}
		limitStage := bson.D{{Key: "$limit", Value: 200}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "purchase_request_count", Value: -1},
			{Key: "click_count", Value: -1},
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
		projectStage := getProjectQueryMyCoordinatePost()

		filterBsonA := bson.A{}

		if analysisSpan == "one_month" {
			oneMonthBefore := time.Now().UTC().AddDate(0, -1, 0)
			filterBsonA = append(filterBsonA, bson.M{"post_date": bson.M{"$gte": oneMonthBefore}})
		} else if analysisSpan == "three_month" {
			threeMonthBefore := time.Now().UTC().AddDate(0, -3, 0)
			filterBsonA = append(filterBsonA, bson.M{"post_date": bson.M{"$gte": threeMonthBefore}})
		}
		if reflect.ValueOf(filterBsonA).Len() > 0 {
			matchFilter := bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: filterBsonA}}}}
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchFilter, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
				})
		} else {
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					lookUpStage, unwindStage, projectStage, sortStage, limitStage,
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
	queryResult, err := util.ExecuteDbQuery(f, "coordinate_post")
	if err != nil {
		return nil, err
	}
	var cordinatePosts = []coordinateModel.CoordinatePostInfo{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			cordinatePosts = append(cordinatePosts, *response)
		}
	}
	return funk.Map(cordinatePosts, func(c coordinateModel.CoordinatePostInfo) analyticsModel.CoordinatePostAnalytics {
		return analyticsModel.CoordinatePostAnalytics{
			ID:                   c.ID,
			Title:                c.Title,
			ShopID:               c.ShopID,
			ShopName:             c.ShopName,
			ShopSettingID:        c.ShopSettingID,
			PostDate:             c.PostDate,
			ImpressionCount:      c.ImpressionCount,
			ClickCount:           c.ClickCount,
			PurchaseRequestCount: c.PurchaseRequestCount,
		}
	}).([]analyticsModel.CoordinatePostAnalytics), err
}

// GetLookupShopQuery shopと結合するためのクエリ
func GetLookupShopQuery() bson.D {
	return bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "shop"},
		{Key: "localField", Value: "shop_id"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "shop"},
	}}}
}

// getProjectQueryMyCoordinatePost 自分のコーデ投稿情報を取得するプロジェクトのクエリ
func getProjectQueryMyCoordinatePost() bson.D {
	return bson.D{{Key: "$project", Value: bson.D{
		{Key: "_id", Value: 1},
		{Key: "title", Value: 1},
		{Key: "status", Value: 1},
		{Key: "shop_id", Value: 1},
		{Key: "shop.name", Value: 1},
		{Key: "shop.shop_setting_id", Value: 1},
		{Key: "detail", Value: 1},
		{Key: "url", Value: 1},
		{Key: "post_date", Value: 1},
		{Key: "post_user_id", Value: 1},
		{Key: "impression_count", Value: 1},
		{Key: "click_count", Value: 1},
		{Key: "purchase_request_count", Value: 1},
		{Key: "model_category", Value: 1},
		{Key: "model_attribute.gender", Value: 1},
		{Key: "model_attribute.silhouette", Value: 1},
		{Key: "model_attribute.height", Value: 1},
		{Key: "model_attribute.weight", Value: 1},
		{Key: "model_attribute.size", Value: 1},
		{Key: "images.key", Value: 1},
		{Key: "images.url", Value: 1},
		{Key: "favorite_count", Value: postRepository.GetFavoriteCountQuery()},
		{Key: "price", Value: 1},
		{Key: "category", Value: 1},
		{Key: "sale.sale_price", Value: 1},
		{Key: "sale.start_date", Value: 1},
		{Key: "sale.end_date", Value: 1},
	}}}
}

// getProjectQueryOtherCoordinatePost 自分以外のコーデ投稿情報を取得するプロジェクトのクエリ
func getProjectQueryOtherCoordinatePost(userID string) bson.D {
	if userID != "" {
		return bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "title", Value: 1},
			{Key: "shop_id", Value: 1},
			{Key: "shop.name", Value: 1},
			{Key: "shop.shop_setting_id", Value: 1},
			{Key: "detail", Value: 1},
			{Key: "url", Value: 1},
			{Key: "post_date", Value: 1},
			{Key: "post_user_id", Value: 1},
			{Key: "model_category", Value: 1},
			{Key: "model_attribute.gender", Value: 1},
			{Key: "model_attribute.silhouette", Value: 1},
			{Key: "model_attribute.height", Value: 1},
			{Key: "model_attribute.weight", Value: 1},
			{Key: "model_attribute.size", Value: 1},
			{Key: "images.key", Value: 1},
			{Key: "images.url", Value: 1},
			{Key: "favorite_count", Value: postRepository.GetFavoriteCountQuery()},
			{Key: "favorite_count_login_user", Value: postRepository.GetFavoriteCountQueryByUser(userID)},
			{Key: "price", Value: 1},
			{Key: "category", Value: 1},
			{Key: "sale.sale_price", Value: 1},
			{Key: "sale.start_date", Value: 1},
			{Key: "sale.end_date", Value: 1},
		}}}
	}
	return bson.D{{Key: "$project", Value: bson.D{
		{Key: "_id", Value: 1},
		{Key: "title", Value: 1},
		{Key: "shop_id", Value: 1},
		{Key: "shop.name", Value: 1},
		{Key: "shop.shop_setting_id", Value: 1},
		{Key: "detail", Value: 1},
		{Key: "url", Value: 1},
		{Key: "post_date", Value: 1},
		{Key: "post_user_id", Value: 1},
		{Key: "model_category", Value: 1},
		{Key: "model_attribute.gender", Value: 1},
		{Key: "model_attribute.silhouette", Value: 1},
		{Key: "model_attribute.height", Value: 1},
		{Key: "model_attribute.weight", Value: 1},
		{Key: "model_attribute.size", Value: 1},
		{Key: "images.key", Value: 1},
		{Key: "images.url", Value: 1},
		{Key: "favorite_count", Value: postRepository.GetFavoriteCountQuery()},
		{Key: "price", Value: 1},
		{Key: "category", Value: 1},
		{Key: "sale.sale_price", Value: 1},
		{Key: "sale.start_date", Value: 1},
		{Key: "sale.end_date", Value: 1},
	}}}
}

// getQueryResultToResponse クエリの結果をresponse用のstructに詰め替え
func getQueryResultToResponse(v primitive.M) (*coordinateModel.CoordinatePostInfo, error) {
	id, errGet := dproxy.New(v).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	title, _ := dproxy.New(v).M("title").String()
	detail, _ := dproxy.New(v).M("detail").String()
	shopID, errGet := dproxy.New(v).M("shop_id").String()
	if errGet != nil {
		return nil, errGet
	}
	shopSettingID, errGet := dproxy.New(v).M("shop").M("shop_setting_id").String()
	if errGet != nil {
		return nil, errGet
	}
	shopName, errGet := dproxy.New(v).M("shop").M("name").String()
	if errGet != nil {
		return nil, errGet
	}
	url, _ := dproxy.New(v).M("url").String()
	postDateInt := int64(v["post_date"].(primitive.DateTime))
	postUser, _ := dproxy.New(v).M("post_user_id").String()

	modelCategory, _ := dproxy.New(v).M("model_category").String()
	gender, _ := dproxy.New(v).M("model_attribute").M("gender").String()
	silhouette, _ := dproxy.New(v).M("model_attribute").M("silhouette").String()
	height, _ := dproxy.New(v).M("model_attribute").M("height").Int64()
	weight, _ := dproxy.New(v).M("model_attribute").M("weight").Int64()
	size, _ := dproxy.New(v).M("model_attribute").M("size").String()

	favoriteCount, _ := dproxy.New(v).M("favorite_count").Int64()
	favoriteCountLoginUser, _ := dproxy.New(v).M("favorite_count_login_user").Int64()
	status, _ := dproxy.New(v).M("status").String()
	impressionCount, _ := dproxy.New(v).M("impression_count").Int64()
	clickCount, _ := dproxy.New(v).M("click_count").Int64()
	purchaseRequestCount, _ := dproxy.New(v).M("purchase_request_count").Int64()

	price, _ := dproxy.New(v).M("price").Int64()
	category, _ := dproxy.New(v).M("category").String()

	var saleResponse *coordinateModel.SaleResponse
	salePrice, _ := dproxy.New(v).M("sale").M("sale_price").Int64()
	if salePrice > 0 {
		saleInfo := v["sale"].(primitive.M)
		saleStartDateInt := int64(saleInfo["start_date"].(primitive.DateTime))
		saleEndDateInt := int64(saleInfo["end_date"].(primitive.DateTime))

		saleStartDate := util.GetMongoDateIntToTime(saleStartDateInt)
		saleEndDate := util.GetMongoDateIntToTime(saleEndDateInt)

		saleResponse = &coordinateModel.SaleResponse{
			SalePrice: int(salePrice),
			StartDate: &saleStartDate,
			EndDate:   &saleEndDate,
		}
	}

	var images []coordinateModel.CoordinateImage
	if v["images"] != nil {
		imagesFromDB := v["images"].(primitive.A)
		for i := 0; i < reflect.ValueOf(imagesFromDB).Len(); i++ {
			imageKey, errGet := dproxy.New(imagesFromDB[i]).M("key").Int64()
			if errGet != nil {
				return nil, errGet
			}
			imageUrl, errGet := dproxy.New(imagesFromDB[i]).M("url").String()
			if errGet != nil {
				return nil, errGet
			}
			images = append(images, coordinateModel.CoordinateImage{
				Key: int(imageKey),
				URL: imageUrl,
			})
		}
	}

	return &coordinateModel.CoordinatePostInfo{
		ID:                   id,
		Title:                title,
		Detail:               detail,
		ShopID:               shopID,
		ShopSettingID:        shopSettingID,
		ShopName:             shopName,
		URL:                  url,
		PostUserID:           postUser,
		PostDate:             util.GetMongoDateIntToTime(postDateInt),
		ModelCategory:        modelCategory,
		FavoriteCount:        int(favoriteCount),
		FavoritedFlg:         favoriteCountLoginUser > 0,
		Status:               status,
		ImpressionCount:      int(impressionCount),
		ClickCount:           int(clickCount),
		PurchaseRequestCount: int(purchaseRequestCount),
		ModelAttribute: coordinateModel.ModelAttribute{
			Gender:     gender,
			Silhouette: silhouette,
			Height:     int(height),
			Weight:     int(weight),
			Size:       size,
		},
		Images:   images,
		Price:    int(price),
		Category: category,
		Sale:     saleResponse,
	}, nil
}
