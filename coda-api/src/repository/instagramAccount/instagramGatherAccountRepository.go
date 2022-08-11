package instagramAccountRepository

import (
	"context"
	"reflect"
	"time"

	instagramAccountModel "coda-api/src/model/instagramAccount"
	postModel "coda-api/src/model/post"
	"coda-api/src/util"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetGatherInstagramAccount 収集対象のアカウントを取得
func GetGatherInstagramAccount() ([]instagramAccountModel.InstagramAccuntID, error) {
	// UTC今日日付
	currentTime, err := time.Parse("20060102", time.Now().UTC().Format("20060102"))
	if err != nil {
		return nil, err
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		// 収集した日の昇順
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{Key: "gather_date", Value: 1}})
		filter := bson.D{{Key: "status", Value: "on"},
			{Key: "$or", Value: bson.A{
				bson.M{"gather_date": bson.D{{Key: "$lt", Value: currentTime}}},
				bson.M{"gather_date": nil},
			},
			}}
		cur, err := col.Find(context.Background(), filter, findOptions)
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
	queryResult, err := util.ExecuteDbQuery(f, "instagram_gather_account")
	if err != nil {
		return nil, err
	}
	var accounts []instagramAccountModel.InstagramAccuntID
	for _, v := range queryResult {
		id, errGet := dproxy.New(v).M("_id").String()
		instagramUserName, errGet2 := dproxy.New(v).M("instagram_user_name").String()
		if errGet != nil {
			return nil, errGet
		}
		if errGet2 != nil {
			return nil, errGet2
		}
		accounts = append(accounts, instagramAccountModel.InstagramAccuntID{ID: id, InstagramUserName: instagramUserName})
	}
	return accounts, nil
}

// UpdateGatherDate 収集日を更新
func UpdateGatherDate(accountID string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		// 収集した日を更新＆エラー日付を消去
		filter := bson.D{{Key: "_id", Value: accountID}}
		update := bson.D{{Key: "$set",
			Value: bson.D{{Key: "gather_date", Value: time.Now().UTC()}}},
			{Key: "$unset", Value: bson.D{{Key: "error_date", Value: ""}}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "instagram_gather_account")
	return err
}

// UpdateErrorDate エラー日付を更新
func UpdateErrorDate(accountID string) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		// 収集した日の昇順
		filter := bson.D{{Key: "_id", Value: accountID}}
		update := bson.D{{Key: "$set",
			Value: bson.D{{Key: "error_date", Value: time.Now().UTC()}}}}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "instagram_gather_account")
	return err
}

// GetGatherInstagramAccountInfoByUserName ユーザー名によるInstagramAccountの取得
func GetGatherInstagramAccountInfoByUserName(instagramUserName string) (*instagramAccountModel.InstagramAccountInfoResponse, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "instagram_user_name", Value: instagramUserName}}
		err := col.FindOne(context.Background(), filter).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "instagram_gather_account")
	if err != nil || queryResult == nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	v := queryResult[0]
	id, errGet := dproxy.New(v).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	status, errGet := dproxy.New(v).M("status").String()
	if errGet != nil {
		return nil, errGet
	}
	category, errGet := dproxy.New(v).M("category").String()
	if errGet != nil {
		return nil, errGet
	}
	return &instagramAccountModel.InstagramAccountInfoResponse{
		Id:                id,
		InstagramUserName: instagramUserName,
		Status:            status,
		Category:          category,
	}, nil
}

// GetGatherInstagramAccountWithPostsByUserName ユーザー名によるInstagramAccountの取得
func GetGatherInstagramAccountWithPostsByUserName(
	instagramUserName string,
	limit int,
	postStatus string,
) (*instagramAccountModel.InstagramAccountWithPosts, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		matchStage1 := bson.D{{Key: "$match", Value: bson.D{
			{Key: "instagram_user_name", Value: instagramUserName}},
		}}
		pipleLineLookup := bson.A{}
		pipleLineLookup = append(pipleLineLookup, bson.D{{Key: "$sort", Value: bson.D{
			{Key: "post_date", Value: -1},
		}}})
		var eqBsonA bson.A
		eqBsonA = append(eqBsonA, "$$account_id")
		eqBsonA = append(eqBsonA, "$post_instagram_account_id")
		pipleLineLookup = append(pipleLineLookup, bson.D{{Key: "$match", Value: bson.D{
			{Key: "$expr", Value: bson.D{{Key: "$eq", Value: eqBsonA}}}}}})
		if postStatus == "open" || postStatus == "close" || postStatus == "notset" {
			pipleLineLookup = append(pipleLineLookup, bson.D{{Key: "$match", Value: bson.D{
				{Key: "status", Value: postStatus}},
			}})
		}
		pipleLineLookup = append(pipleLineLookup, bson.D{{Key: "$limit", Value: limit}})
		lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "post"},
			{Key: "let", Value: bson.D{{Key: "account_id", Value: "$_id"}}},
			{Key: "pipeline", Value: pipleLineLookup},
			{Key: "as", Value: "posts"},
		}}}
		projectStage := bson.D{{Key: "$project", Value: bson.D{
			{Key: "_id", Value: 1},
			{Key: "status", Value: 1},
			{Key: "category", Value: 1},
			{Key: "gather_date", Value: 1},
			{Key: "posts._id", Value: 1},
			{Key: "posts.status", Value: 1},
			{Key: "posts.content_url", Value: 1},
			{Key: "posts.genre", Value: 1},
			{Key: "posts.post_date", Value: 1},
		}}}

		pipeLine := mongo.Pipeline{matchStage1, lookupStage, projectStage}

		cur, err := col.Aggregate(context.Background(), pipeLine)
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
	queryResult, err := util.ExecuteDbQuery(f, "instagram_gather_account")
	if err != nil || queryResult == nil {
		return nil, err
	}
	v := queryResult[0]
	id, errGet := dproxy.New(v).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	accountStatus, errGet := dproxy.New(v).M("status").String()
	if errGet != nil {
		return nil, errGet
	}
	category, errGet := dproxy.New(v).M("category").String()
	if errGet != nil {
		return nil, errGet
	}
	gatherDateInt := int64(v["gather_date"].(primitive.DateTime))

	var posts []postModel.PostWithAccountResponseForAdmin
	getPosts := v["posts"].(primitive.A)
	for i := 0; i < reflect.ValueOf(getPosts).Len(); i++ {
		post := getPosts[i]
		postID, errGet := dproxy.New(post).M("_id").String()
		if errGet != nil {
			return nil, errGet
		}
		postStatus, errGet := dproxy.New(post).M("status").String()
		if errGet != nil {
			return nil, errGet
		}
		contentURL, errGet := dproxy.New(post).M("content_url").String()
		if errGet != nil {
			return nil, errGet
		}
		postDateInt := int64(post.(primitive.M)["post_date"].(primitive.DateTime))
		genre, _ := dproxy.New(post).M("genre").String()
		posts = append(posts, postModel.PostWithAccountResponseForAdmin{
			ID:         postID,
			Status:     postStatus,
			ContentURL: contentURL,
			PostDate:   util.GetMongoDateIntToTime(postDateInt),
			PostGenre:  genre,
		})
	}
	return &instagramAccountModel.InstagramAccountWithPosts{
		Id:                id,
		InstagramUserName: instagramUserName,
		Status:            accountStatus,
		Category:          category,
		GatherDate:        util.GetMongoDateIntToTime(gatherDateInt),
		Posts:             posts,
	}, nil
}

// AddInstagramAccount インスタグラムアカウントの追加
func AddInstagramAccount(
	instagramUserName string,
	status string,
	gender string,
	silhouette string,
	height string,
	genre string,
) error {
	uuid, err := util.GenerateUUID()
	if err != nil {
		return err
	}
	initialGatherTime, err := time.Parse("20060102", "19000101")
	if err != nil {
		return err
	}
	registerBson := bson.M{
		"_id":                 uuid,
		"instagram_user_name": instagramUserName,
		"status":              status,
		"category":            gender + "-" + silhouette + "-" + height + "-" + genre,
		"gather_date":         initialGatherTime,
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.InsertOne(context.Background(), registerBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err = util.ExecuteDbQuery(f, "instagram_gather_account")
	if err != nil {
		return err
	}
	return nil
}

// EditInstagramAccount インスタグラムアカウントの編集
func EditInstagramAccount(
	id string,
	status string,
	gender string,
	silhouette string,
	height string,
	genre string,
) error {
	category := gender + "-" + silhouette + "-" + height + "-" + genre
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: id}}
		update := bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "status", Value: status},
				{Key: "category", Value: category},
			},
		}}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "instagram_gather_account")
	if err != nil {
		return err
	}
	return nil
}
