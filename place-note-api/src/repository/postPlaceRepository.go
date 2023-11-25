package repository

import (
	"context"
	modelDb "placeNote/src/model/db"
	"placeNote/src/placeNoteUtil"

	"github.com/bufbuild/connect-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const postPlacesCollectionName = "post_places"

// AddPostPlaceRepository 場所の追加
func AddPostPlaceRepository(postPlaceEntity modelDb.PostPlacesEntity) *connect.Error {
	col := placeNoteUtil.GetDbCollection(postPlacesCollectionName)
	_, err := col.InsertOne(context.Background(), postPlaceEntity)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}

// UpdatePostPlaceRepository 場所の更新
func UpdatePostPlaceRepository(updatePostPlaceEntity modelDb.PostPlacesEntity) *connect.Error {
	col := placeNoteUtil.GetDbCollection(postPlacesCollectionName)
	filter := bson.M{
		"_id":                    updatePostPlaceEntity.ID,
		"create_user_account_id": updatePostPlaceEntity.CreateUserAccountId,
	}
	updateSet := bson.M{"$set": bson.M{
		"name":             updatePostPlaceEntity.Name,
		"lon_lat":          updatePostPlaceEntity.LonLat,
		"prefecture_code":  updatePostPlaceEntity.PrefectureCode,
		"category_id_list": updatePostPlaceEntity.CategoryIdList,
		"detail":           updatePostPlaceEntity.Detail,
		"url_list":         updatePostPlaceEntity.UrlList,
	}}

	_, err := col.UpdateOne(context.Background(), filter, updateSet)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}

// DeletePostPlaceRepository 場所の削除
func DeletePostPlaceRepository(deleteId string, userAccountId string) *connect.Error {
	col := placeNoteUtil.GetDbCollection(postPlacesCollectionName)

	deleteFilter := bson.M{
		"_id":                    deleteId,
		"create_user_account_id": userAccountId,
	}
	_, err := col.DeleteOne(context.Background(), deleteFilter)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}

	return nil
}

// GetUserPostPlaceListByUserAccountIdRepository ユーザが登録した場所一覧を取得
func GetUserPostPlaceListByUserAccountIdRepository(userAccountId string) ([]modelDb.PostPlacesEntity, *connect.Error) {
	col := placeNoteUtil.GetDbCollection(postPlacesCollectionName)
	var docs []modelDb.PostPlacesEntity

	matchStage := bson.M{"$match": bson.M{
		"create_user_account_id": userAccountId},
	}
	projectStage := bson.M{"$project": bson.M{
		"_id":              1,
		"name":             1,
		"address":          1,
		"lon_lat":          1,
		"prefecture_code":  1,
		"category_id_list": 1,
		"detail":           1,
		"url_list":         1,
	}}

	cur, err := col.Aggregate(context.Background(),
		[]bson.M{
			matchStage, projectStage,
		})
	if err != nil {
		return docs, connect.NewError(connect.CodeInternal, err)
	}

	for cur.Next(context.Background()) {
		var doc modelDb.PostPlacesEntity
		if err = cur.Decode(&doc); err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

// GetUserPostPlaceByIdRepository idによる場所取得
func GetUserPostPlaceByIdRepository(userAccountId string, placeId string) (*modelDb.PostPlacesEntity, *connect.Error) {
	col := placeNoteUtil.GetDbCollection(postPlacesCollectionName)

	var result modelDb.PostPlacesEntity
	filter := bson.M{
		"_id":                    placeId,
		"create_user_account_id": userAccountId,
	}
	err := col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &result, nil
}

// DeleteCategoryFromPlaceRepository 場所からカテゴリーを削除
func DeleteCategoryFromPlaceRepository(deleteCategoryId string, userAccountId string) *connect.Error {
	col := placeNoteUtil.GetDbCollection(postPlacesCollectionName)
	filter := bson.M{
		"create_user_account_id": userAccountId,
		"category_id_list":       bson.M{"$in": bson.A{deleteCategoryId}},
	}
	updateSet := bson.M{"$pull": bson.M{
		"category_id_list": deleteCategoryId,
	}}

	_, err := col.UpdateMany(context.Background(), filter, updateSet)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}
