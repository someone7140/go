package itemPostRepository

import (
	"context"
	"reflect"

	authModel "coda-api/src/model/auth"
	itemPostModel "coda-api/src/model/item"
	postRepository "coda-api/src/repository/post"
	userRepository "coda-api/src/repository/user"
	"coda-api/src/util"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// GetRecentItemPosts 最新のアイテム投稿を取得
func GetRecentItemPosts(limit int) ([]itemPostModel.ItemPostResponse, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
		projectStage := getProjectQueryOtherItemPost("")
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
	queryResult, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return nil, err
	}
	var itemPosts = []itemPostModel.ItemPostResponse{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			itemPosts = append(itemPosts, *response)
		}
	}
	return itemPosts, err
}

// GetRecentItemWithImagePosts 最新の画像があるアイテム投稿を取得
func GetRecentItemWithImagePosts(limit int, userID string) ([]itemPostModel.ItemPostResponse, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "$and", Value: bson.A{
				bson.M{"status": "open"},
				bson.M{"image_url": bson.D{{Key: "$ne", Value: nil}}},
				bson.M{"image_url": bson.D{{Key: "$ne", Value: ""}}},
			}},
		}},
		}
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		projectStage := getProjectQueryOtherItemPost(userID)
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
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
	queryResult, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return nil, err
	}
	var itemPosts = []itemPostModel.ItemPostResponse{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			itemPosts = append(itemPosts, *response)
		}
	}
	return itemPosts, err
}

// GetCategoryMatchedItemPosts カテゴリーにマッチしたアイテム投稿を取得
func GetCategoryMatchedItemPosts(
	limit int,
	targetCategories map[string]int,
	userID string) ([]itemPostModel.ItemPostResponse, error) {
	// カテゴリー名のリストを抽出
	categoryNameBsonA := bson.A{}
	for categoryKey := range targetCategories {
		categoryNameBsonA = append(categoryNameBsonA, categoryKey)
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage1 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		projectStage := getProjectQueryOtherItemPost(userID)
		matchCategoryStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "category", Value: bson.D{{Key: "$in", Value: categoryNameBsonA}}}}}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		// limitの2倍をまずは取得
		limitStage := bson.D{{Key: "$limit", Value: limit * 2}}
		var cur *mongo.Cursor
		var err error
		cur, err = col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage1, lookUpStage, unwindStage, projectStage, matchCategoryStage, sortStage, limitStage,
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
	if err != nil {
		return nil, err
	}
	var itemPosts = []itemPostModel.ItemPostResponse{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			itemPosts = append(itemPosts, *response)
		}
	}
	return util.GetItemPostResponseSortedByPostDate(itemPosts, targetCategories, limit), nil
}

// GetItemsByUserID ユーザ指定の投稿一覧取得
func GetItemsByUserID(userSettingID string, loginInfo *authModel.AuthUserResponse, limit int) ([]itemPostModel.ItemPostResponse, error) {
	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	adminFlag := loginInfo != nil && loginInfo.UserType == "admin"
	refUser, err := userRepository.GetUserEntityByUserSettingID(userSettingID, "")
	if refUser == nil || err != nil {
		return nil, err
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "post_user_id", Value: refUser.ID}},
		}}
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		var cur *mongo.Cursor
		var err error
		projectStage := map[bool]bson.D{
			true:  getProjectQueryMyItemPost(),
			false: getProjectQueryOtherItemPost(myUserID),
		}[refUser.ID == myUserID || adminFlag]
		if refUser.ID == myUserID || adminFlag {
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStage, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
				})
		} else {
			matchStatusStage := bson.D{{Key: "$match", Value: bson.D{
				{Key: "status", Value: "open"}},
			}}
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStage, matchStatusStage, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
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
	queryResult, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return nil, err
	}
	var itemPosts = []itemPostModel.ItemPostResponse{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			itemPosts = append(itemPosts, *response)
		}
	}
	return itemPosts, err
}

// GetItemByPostID アイテム投稿IDを指定して取得
func GetItemByPostID(itemPostID string, myUserID string, myPostFlag bool) (*itemPostModel.ItemPostResponse, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage := bson.D{{Key: "$match", Value: bson.D{
			{Key: "_id", Value: itemPostID}},
		}}
		matchStageStatus := map[bool]bson.D{
			true: {{Key: "$match", Value: bson.M{"status": bson.D{{Key: "$ne", Value: nil}}}}},
			false: {{Key: "$match", Value: bson.D{
				{Key: "status", Value: "open"}},
			}}}[myPostFlag]
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		var cur *mongo.Cursor
		var err error
		projectStage := getProjectQueryOtherItemPost(myUserID)
		cur, err = col.Aggregate(context.Background(),
			mongo.Pipeline{
				matchStage, matchStageStatus, lookUpStage, unwindStage, projectStage,
			})
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
	return getQueryResultToResponse(queryResult[0])
}

// GetFavoritedItemPosts いいねをしたアイテム投稿を取得
func GetFavoritedItemPosts(limit int, userID string) ([]itemPostModel.ItemPostResponse, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage1 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "status", Value: "open"}},
		}}
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "title", Value: 1},
			{Key: "item_type", Value: 1},
			{Key: "detail", Value: 1},
			{Key: "user._id", Value: 1},
			{Key: "user.user_setting_id", Value: 1},
			{Key: "user.name", Value: 1},
			{Key: "user.icon_url", Value: 1},
			{Key: "url", Value: 1},
			{Key: "post_date", Value: 1},
			{Key: "category", Value: 1},
			{Key: "attribute.gender", Value: 1},
			{Key: "attribute.silhouette", Value: 1},
			{Key: "attribute.complex", Value: 1},
			{Key: "image_url", Value: 1},
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
	queryResult, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return nil, err
	}
	var itemPosts = []itemPostModel.ItemPostResponse{}

	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			itemPosts = append(itemPosts, *response)
		}
	}
	return itemPosts, err
}

// GetSearchItemPosts アイテム投稿の検索
func GetSearchItemPosts(limit int, searchRequest itemPostModel.ItemPostSearchRequest, loginInfo *authModel.AuthUserResponse) ([]itemPostModel.ItemPostResponse, error) {
	var refUserID = ""
	if searchRequest.UserSettingID != "" {
		refUser, err := userRepository.GetUserEntityByUserSettingID(searchRequest.UserSettingID, "")
		if refUser == nil || err != nil {
			return nil, err
		}
		refUserID = refUser.ID
	}
	var userID = ""
	if loginInfo != nil {
		userID = loginInfo.ID
	}

	myPostRefFlag := refUserID != "" && loginInfo != nil && (refUserID == loginInfo.ID || loginInfo.UserType == "admin")

	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStageStatus := map[bool]bson.D{
			true: {{Key: "$match", Value: bson.M{"status": bson.D{{Key: "$ne", Value: nil}}}}},
			false: {{Key: "$match", Value: bson.D{
				{Key: "status", Value: "open"}},
			}}}[myPostRefFlag]
		lookUpStage := GetLookupUserQuery()
		unwindStage := bson.D{{Key: "$unwind", Value: "$user"}}
		limitStage := bson.D{{Key: "$limit", Value: limit}}
		sortStage := bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}}
		projectStage := map[bool]bson.D{
			true:  getProjectQueryMyItemPost(),
			false: getProjectQueryOtherItemPost(userID),
		}[myPostRefFlag]
		// リクエストから検索条件の設定
		searchBsonA := bson.A{}
		if searchRequest.Keyword != "" {
			searchBsonA = append(searchBsonA, bson.M{"$or": bson.A{
				bson.M{"title": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: searchRequest.Keyword, Options: "i"}}}},
				bson.M{"detail": bson.D{{Key: "$regex", Value: primitive.Regex{Pattern: searchRequest.Keyword, Options: "i"}}}},
			}})
		}
		if searchRequest.ItemType != "" {
			searchBsonA = append(searchBsonA, bson.M{"item_type": searchRequest.ItemType})
		}
		if searchRequest.Gender != "" {
			searchBsonA = append(searchBsonA, bson.M{"gender": searchRequest.Gender})
		}
		if searchRequest.Silhouette != "" {
			searchBsonA = append(searchBsonA, bson.M{"silhouette": searchRequest.Silhouette})
		}
		if searchRequest.Complex != "" {
			searchBsonA = append(searchBsonA, bson.M{"complex": searchRequest.Complex})
		}
		if refUserID != "" {
			searchBsonA = append(searchBsonA, bson.M{"post_user_id": refUserID})
		}

		var cur *mongo.Cursor
		var err error
		if reflect.ValueOf(searchBsonA).Len() > 0 {
			matchStageSearch := bson.D{{Key: "$match", Value: bson.D{{Key: "$and", Value: searchBsonA}}}}
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStageStatus, matchStageSearch, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
				})
		} else {
			cur, err = col.Aggregate(context.Background(),
				mongo.Pipeline{
					matchStageStatus, lookUpStage, unwindStage, projectStage, sortStage, limitStage,
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
	queryResult, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return nil, err
	}
	var itemPosts = []itemPostModel.ItemPostResponse{}
	for _, v := range queryResult {
		response, err := getQueryResultToResponse(v)
		if err == nil || response != nil {
			itemPosts = append(itemPosts, *response)
		}
	}
	return itemPosts, err
}

// GetImageURLsByUserID userIDを指定してimageURLの一覧を取得
func GetImageURLsByUserID(userID string) ([]string, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "$and", Value: bson.A{
			bson.M{"post_user_id": userID},
			bson.M{"image_url": bson.D{{Key: "$ne", Value: nil}}},
			bson.M{"image_url": bson.D{{Key: "$ne", Value: ""}}},
		}}}
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
	queryResult, err := util.ExecuteDbQuery(f, "item_post")
	if err != nil {
		return nil, err
	}
	var getImageURLs = []string{}
	for _, v := range queryResult {
		imageURL, errGet := dproxy.New(v).M("image_url").String()
		if errGet != nil {
			return nil, errGet
		}
		getImageURLs = append(getImageURLs, imageURL)
	}
	return getImageURLs, nil
}

// GetLookupUserQuery userと結合するためのクエリ
func GetLookupUserQuery() bson.D {
	return bson.D{{Key: "$lookup", Value: bson.D{
		{Key: "from", Value: "user"},
		{Key: "localField", Value: "post_user_id"},
		{Key: "foreignField", Value: "_id"},
		{Key: "as", Value: "user"},
	}}}
}

// getProjectQueryMyItemPost 自分のアイテム投稿情報を取得するプロジェクトのクエリ
func getProjectQueryMyItemPost() bson.D {
	return bson.D{{Key: "$project", Value: bson.D{
		{Key: "_id", Value: 1},
		{Key: "title", Value: 1},
		{Key: "item_type", Value: 1},
		{Key: "detail", Value: 1},
		{Key: "user._id", Value: 1},
		{Key: "user.user_setting_id", Value: 1},
		{Key: "user.name", Value: 1},
		{Key: "user.icon_url", Value: 1},
		{Key: "url", Value: 1},
		{Key: "post_date", Value: 1},
		{Key: "category", Value: 1},
		{Key: "attribute.gender", Value: 1},
		{Key: "attribute.silhouette", Value: 1},
		{Key: "attribute.complex", Value: 1},
		{Key: "image_url", Value: 1},
		{Key: "favorite_count", Value: postRepository.GetFavoriteCountQuery()},
		{Key: "status", Value: 1},
		{Key: "impression_count", Value: 1},
		{Key: "click_count", Value: 1},
	}}}
}

// getProjectQueryOtherItemPost 自分以外のアイテム投稿情報を取得するプロジェクトのクエリ
func getProjectQueryOtherItemPost(userID string) bson.D {
	if userID != "" {
		return bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "title", Value: 1},
			{Key: "item_type", Value: 1},
			{Key: "detail", Value: 1},
			{Key: "user._id", Value: 1},
			{Key: "user.user_setting_id", Value: 1},
			{Key: "user.name", Value: 1},
			{Key: "user.icon_url", Value: 1},
			{Key: "url", Value: 1},
			{Key: "post_date", Value: 1},
			{Key: "category", Value: 1},
			{Key: "attribute.gender", Value: 1},
			{Key: "attribute.silhouette", Value: 1},
			{Key: "attribute.complex", Value: 1},
			{Key: "image_url", Value: 1},
			{Key: "favorite_count", Value: postRepository.GetFavoriteCountQuery()},
			{Key: "favorite_count_login_user", Value: postRepository.GetFavoriteCountQueryByUser(userID)},
		}}}
	}
	return bson.D{{Key: "$project", Value: bson.D{
		{Key: "_id", Value: 1},
		{Key: "title", Value: 1},
		{Key: "item_type", Value: 1},
		{Key: "detail", Value: 1},
		{Key: "user._id", Value: 1},
		{Key: "user.user_setting_id", Value: 1},
		{Key: "user.name", Value: 1},
		{Key: "user.icon_url", Value: 1},
		{Key: "url", Value: 1},
		{Key: "post_date", Value: 1},
		{Key: "category", Value: 1},
		{Key: "attribute.gender", Value: 1},
		{Key: "attribute.silhouette", Value: 1},
		{Key: "attribute.complex", Value: 1},
		{Key: "image_url", Value: 1},
		{Key: "favorite_count", Value: postRepository.GetFavoriteCountQuery()},
	}}}
}

// getQueryResultToResponse クエリの結果をresponse用のstructに詰め替え
func getQueryResultToResponse(v primitive.M) (*itemPostModel.ItemPostResponse, error) {
	id, errGet := dproxy.New(v).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	title, _ := dproxy.New(v).M("title").String()
	itemType, errGet := dproxy.New(v).M("item_type").String()
	if errGet != nil {
		return nil, errGet
	}
	detail, _ := dproxy.New(v).M("detail").String()
	userID, errGet := dproxy.New(v).M("user").M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	userSettingID, errGet := dproxy.New(v).M("user").M("user_setting_id").String()
	if errGet != nil {
		return nil, errGet
	}
	userName, errGet := dproxy.New(v).M("user").M("name").String()
	if errGet != nil {
		return nil, errGet
	}
	userIconURL, _ := dproxy.New(v).M("user").M("icon_url").String()
	url, _ := dproxy.New(v).M("url").String()
	postDateInt := int64(v["post_date"].(primitive.DateTime))
	category, _ := dproxy.New(v).M("category").String()
	gender, _ := dproxy.New(v).M("attribute").M("gender").String()
	silhouette, _ := dproxy.New(v).M("attribute").M("silhouette").String()
	complex, _ := dproxy.New(v).M("attribute").M("complex").String()
	imageURL, _ := dproxy.New(v).M("image_url").String()

	favoriteCount, _ := dproxy.New(v).M("favorite_count").Int64()
	favoriteCountLoginUser, _ := dproxy.New(v).M("favorite_count_login_user").Int64()
	status, _ := dproxy.New(v).M("status").String()
	impressionCount, _ := dproxy.New(v).M("impression_count").Int64()
	clickCount, _ := dproxy.New(v).M("click_count").Int64()

	return &itemPostModel.ItemPostResponse{
		ID:              id,
		Title:           title,
		ItemType:        itemType,
		Detail:          detail,
		UserID:          userID,
		UserSettingID:   userSettingID,
		UserName:        userName,
		UserIconURL:     userIconURL,
		URL:             url,
		PostDate:        util.GetMongoDateIntToTime(postDateInt),
		Category:        category,
		Gender:          gender,
		Silhouette:      silhouette,
		Complex:         complex,
		ImageURL:        imageURL,
		FavoriteCount:   int(favoriteCount),
		FavoritedFlg:    favoriteCountLoginUser > 0,
		Status:          status,
		ImpressionCount: int(impressionCount),
		ClickCount:      int(clickCount),
	}, nil
}
