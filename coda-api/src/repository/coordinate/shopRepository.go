package coordinateRepository

import (
	"context"
	"time"

	"coda-api/src/util"

	"github.com/koron/go-dproxy"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	shopModel "coda-api/src/model/coordinate"
)

var selectProjection = bson.M{
	"_id":                         1,
	"name":                        1,
	"shop_setting_id":             1,
	"shop_url":                    1,
	"detail":                      1,
	"create_date":                 1,
	"shop_target_user.gender":     1,
	"shop_target_user.silhouette": 1,
	"shop_target_user.min_height": 1,
	"shop_target_user.max_height": 1,
	"shop_target_user.min_weight": 1,
	"shop_target_user.max_weight": 1,
}

// GetShopInfoByShopSettingID shopSettingIDでのユーザ取得
func GetShopInfoByShopSettingID(shopSettingID string) (*shopModel.ShopInfo, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		var result bson.M
		var results []bson.M
		filter := bson.D{{Key: "shop_setting_id", Value: shopSettingID}}
		err := col.FindOne(context.Background(), filter, options.FindOne().SetProjection(selectProjection)).Decode(&result)
		if err == mongo.ErrNoDocuments {
			return nil, nil
		} else if err != nil {
			return nil, err
		}
		return append(results, result), nil
	}
	queryResult, err := util.ExecuteDbQuery(f, "shop")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	return decodeShopInfo(queryResult[0])
}

// GetShopList shopのリストを取得
func GetShopList() ([]shopModel.ShopInfo, error) {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		// 作成日の降順
		findOptions := options.Find()
		findOptions.SetSort(bson.D{{Key: "create_date", Value: -1}})
		cur, err := col.Find(context.Background(), bson.D{}, findOptions)
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
	queryResult, err := util.ExecuteDbQuery(f, "shop")
	if err != nil {
		return nil, err
	}
	if queryResult == nil {
		return nil, nil
	}
	var shopList []shopModel.ShopInfo
	for _, v := range queryResult {
		shop, err := decodeShopInfo(v)
		if err != nil {
			return nil, err
		}
		shopList = append(shopList, *shop)
	}
	return shopList, nil
}

// AddShop shopの追加
func AddShop(uuid string, req shopModel.ShopInfo) error {
	now := time.Now().UTC()
	registerBson := bson.M{
		"_id":             uuid,
		"name":            req.Name,
		"shop_setting_id": req.ShopSettingID,
		"shop_url":        map[bool]*string{true: &req.ShopURL, false: nil}[req.ShopURL != ""],
		"detail":          map[bool]*string{true: &req.Detail, false: nil}[req.Detail != ""],
		"create_date":     now,
		"shop_target_user": bson.M{
			"gender":     req.ShopTaegetUser.Gender,
			"silhouette": req.ShopTaegetUser.Silhouette,
			"min_height": map[bool]*int{true: &req.ShopTaegetUser.MinHeight, false: nil}[req.ShopTaegetUser.MinHeight > 0],
			"max_height": map[bool]*int{true: &req.ShopTaegetUser.MaxHeight, false: nil}[req.ShopTaegetUser.MaxHeight > 0],
			"min_weight": map[bool]*int{true: &req.ShopTaegetUser.MinWeight, false: nil}[req.ShopTaegetUser.MinWeight > 0],
			"max_weight": map[bool]*int{true: &req.ShopTaegetUser.MaxWeight, false: nil}[req.ShopTaegetUser.MaxWeight > 0],
		},
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.InsertOne(context.Background(), registerBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "shop")
	if err != nil {
		return err
	}
	return nil
}

// UpdateShop shopの更新
func UpdateShop(req shopModel.ShopInfo) error {
	f := func(col *mongo.Collection) ([]bson.M, error) {
		filter := bson.D{{Key: "_id", Value: req.ID}}
		update := bson.D{{Key: "$set",
			Value: bson.D{
				{Key: "name", Value: req.Name},
				{Key: "shop_setting_id", Value: req.ShopSettingID},
				{Key: "shop_url", Value: req.ShopURL},
				{Key: "detail", Value: req.Detail},
				{Key: "shop_target_user", Value: bson.D{
					{Key: "gender", Value: req.ShopTaegetUser.Gender},
					{Key: "silhouette", Value: req.ShopTaegetUser.Silhouette},
					{Key: "min_height", Value: req.ShopTaegetUser.MinHeight},
					{Key: "max_height", Value: req.ShopTaegetUser.MaxHeight},
					{Key: "min_weight", Value: req.ShopTaegetUser.MinWeight},
					{Key: "max_weight", Value: req.ShopTaegetUser.MaxWeight},
				},
				},
			}},
		}
		_, err := col.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "shop")
	if err != nil {
		return err
	}
	return nil
}

// DeleteShop ショップの削除
func DeleteShop(shopID string) error {
	filter := bson.D{{Key: "_id", Value: shopID}}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.DeleteOne(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "shop")
	if err != nil {
		return err
	}
	return nil
}

// decodeShopInfo bson.Mからユーザ詳細の構造体にdecode
func decodeShopInfo(input primitive.M) (*shopModel.ShopInfo, error) {
	id, errGet := dproxy.New(input).M("_id").String()
	if errGet != nil {
		return nil, errGet
	}
	shopSettingID, errGet := dproxy.New(input).M("shop_setting_id").String()
	if errGet != nil {
		return nil, errGet
	}
	name, errGet := dproxy.New(input).M("name").String()
	if errGet != nil {
		return nil, errGet
	}
	shopURL, _ := dproxy.New(input).M("shop_url").String()
	detail, _ := dproxy.New(input).M("detail").String()
	createDate := int64(input["create_date"].(primitive.DateTime))
	gender, _ := dproxy.New(input).M("shop_target_user").M("gender").String()
	silhouette, _ := dproxy.New(input).M("shop_target_user").M("silhouette").String()
	minHeight, _ := dproxy.New(input).M("shop_target_user").M("min_height").Int64()
	maxHeight, _ := dproxy.New(input).M("shop_target_user").M("max_height").Int64()
	minWeight, _ := dproxy.New(input).M("shop_target_user").M("min_weight").Int64()
	maxWeight, _ := dproxy.New(input).M("shop_target_user").M("max_weight").Int64()

	return &shopModel.ShopInfo{
		ID:            id,
		Name:          name,
		ShopSettingID: shopSettingID,
		ShopURL:       shopURL,
		Detail:        detail,
		CreateDate:    util.GetMongoDateIntToTime(createDate),
		ShopTaegetUser: shopModel.ShopTaegetUser{
			Gender:     gender,
			Silhouette: silhouette,
			MinHeight:  int(minHeight),
			MaxHeight:  int(maxHeight),
			MinWeight:  int(minWeight),
			MaxWeight:  int(maxWeight),
		},
	}, nil
}
