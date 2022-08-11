package coordinateRepository

import (
	"context"
	"time"

	authModel "coda-api/src/model/auth"
	coordinateModel "coda-api/src/model/coordinate"
	"coda-api/src/util"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// AddCoordinateRepositoryPost コーデ投稿データをDB登録
func AddCoordinateRepositoryPost(
	uuid string,
	request coordinateModel.CoordinatePostRequest,
	images []coordinateModel.CoordinateImage,
	userID string,
	shopID string,
) error {
	now := time.Now().UTC()
	var emptyArray []string
	registerBson := bson.M{
		"_id":              uuid,
		"title":            request.Title,
		"url":              request.URL,
		"detail":           request.Detail,
		"status":           request.Status,
		"shop_id":          shopID,
		"post_user_id":     userID,
		"post_date":        now,
		"impression_count": 0,
		"click_count":      0,
		"model_category":   util.GetCategoriesFromAttributeInput(request.Gender, request.Silhouette, request.Height, emptyArray),
		"model_attribute": bson.M{
			"gender":     request.Gender,
			"silhouette": request.Silhouette,
			"height":     request.Height,
			"weight":     request.Weight,
			"size":       request.Size,
		},
		"images":   getImagesOfBsonA(images),
		"price":    request.Price,
		"category": request.Category,
		"sale":     getRegisterSaleInfo(request.Sale),
	}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.InsertOne(context.Background(), registerBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "coordinate_post")
	return err
}

// UpdateCoordinateRepositoryPost コーデ投稿データをDB更新
func UpdateCoordinateRepositoryPost(
	request coordinateModel.CoordinatePostUpdateRequest,
	images []coordinateModel.CoordinateImage,
	userID string,
	shopID string) error {
	filter := bson.D{
		{Key: "_id", Value: request.ID},
		{Key: "post_user_id", Value: userID},
	}
	var emptyArray []string
	updateBson := bson.D{{Key: "$set",
		Value: bson.M{
			"title":          request.Title,
			"url":            request.URL,
			"detail":         request.Detail,
			"status":         request.Status,
			"shop_id":        shopID,
			"model_category": util.GetCategoriesFromAttributeInput(request.Gender, request.Silhouette, request.Height, emptyArray),
			"model_attribute": bson.M{
				"gender":     request.Gender,
				"silhouette": request.Silhouette,
				"height":     request.Height,
				"weight":     request.Weight,
				"size":       request.Size,
			},
			"images":   getImagesOfBsonA(images),
			"price":    request.Price,
			"category": request.Category,
			"sale":     getRegisterSaleInfo(request.Sale),
		}}}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.UpdateOne(context.Background(), filter, updateBson)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "coordinate_post")
	return err
}

// DeleteCoordinatePost コーデ投稿の削除
func DeleteCoordinatePost(postID string, user *authModel.AuthUserResponse) error {
	filter := map[bool]bson.D{
		true: {
			{Key: "_id", Value: postID},
		},
		false: {
			{Key: "_id", Value: postID},
			{Key: "post_user_id", Value: user.ID},
		},
	}[user.UserType == "admin"]
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.DeleteOne(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "coordinate_post")
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllCoordinatePostsByUser ユーザの全コーデ投稿を削除
func DeleteAllCoordinatePostsByUser(userID string) error {
	filter := bson.D{{Key: "post_user_id", Value: userID}}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.DeleteMany(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "coordinate_post")
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllCoordinatePostsByShop ショップの全コーデ投稿を削除
func DeleteAllCoordinatePostsByShop(shopId string) error {
	filter := bson.D{{Key: "shop_id", Value: shopId}}
	f := func(col *mongo.Collection) ([]bson.M, error) {
		_, err := col.DeleteMany(context.Background(), filter)
		if err != nil {
			return nil, err
		}
		return nil, nil
	}
	_, err := util.ExecuteDbQuery(f, "coordinate_post")
	if err != nil {
		return err
	}
	return nil
}

// imageのstructをbson.Aに変換
func getImagesOfBsonA(images []coordinateModel.CoordinateImage) bson.A {
	registerImages := bson.A{}
	for _, image := range images {
		registerImages = append(registerImages, bson.M{
			"key": image.Key,
			"url": image.URL,
		})
	}
	return registerImages
}

// セール情報の登録内容を取得
func getRegisterSaleInfo(saleRequest *coordinateModel.SaleRequest) *bson.M {
	if saleRequest != nil &&
		saleRequest.SalePrice > 0 &&
		saleRequest.StartDate != nil &&
		saleRequest.EndDate != nil {
		return &bson.M{
			"sale_price": saleRequest.SalePrice,
			"start_date": saleRequest.StartDate,
			"end_date":   saleRequest.EndDate,
		}
	} else {
		return nil
	}
}
