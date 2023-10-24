package repository

import (
	"context"
	modelDb "placeNote/src/model/db"
	"placeNote/src/placeNoteUtil"

	"github.com/bufbuild/connect-go"
	"go.mongodb.org/mongo-driver/bson"
)

const postCategoriesCollectionName = "post_categories"

// AddPostCategoryRepository 投稿カテゴリーの追加
func AddPostCategoryRepository(postCategoryEntity modelDb.PostCategoriesEntity) *connect.Error {
	col := placeNoteUtil.GetDbCollection(postCategoriesCollectionName)
	_, err := col.InsertOne(context.Background(), postCategoryEntity)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}

// UpdatePostCategoryRepository 投稿カテゴリーの更新
func UpdatePostCategoryRepository(updatePostCategoryEntity modelDb.PostCategoriesEntity) *connect.Error {
	col := placeNoteUtil.GetDbCollection(postCategoriesCollectionName)
	filter := bson.M{
		"_id":                    updatePostCategoryEntity.ID,
		"create_user_account_id": updatePostCategoryEntity.CreateUserAccountId,
	}
	updateSet := bson.M{"$set": bson.M{
		"name":               updatePostCategoryEntity.Name,
		"parent_category_id": updatePostCategoryEntity.ParentCategoryId,
		"memo":               updatePostCategoryEntity.Memo,
		"display_order":      updatePostCategoryEntity.DisplayOrder,
	}}

	_, err := col.UpdateOne(context.Background(), filter, updateSet)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}

// DeletePostCategoryRepository 投稿カテゴリーの削除
func DeletePostCategoryRepository(deleteId string, userAccountId string) *connect.Error {
	col := placeNoteUtil.GetDbCollection(postCategoriesCollectionName)

	// 削除カテゴリーを親として設定しているカテゴリーを更新
	filter := bson.M{
		"parent_category_id":     deleteId,
		"create_user_account_id": userAccountId,
	}
	updateSet := bson.M{"$set": bson.M{
		"parent_category_id": nil,
	}}
	_, err := col.UpdateMany(context.Background(), filter, updateSet)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}

	// 対象カテゴリーを削除
	deleteFilter := bson.M{
		"_id":                    deleteId,
		"create_user_account_id": userAccountId,
	}
	_, err = col.DeleteOne(context.Background(), deleteFilter)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}

	return nil
}

// GetPostCategoryListByUserAccountIdRepository ユーザが登録した投稿カテゴリー一覧を取得
func GetPostCategoryListByUserAccountIdRepository(userAccountId string) ([]modelDb.PostCategoriesEntity, *connect.Error) {
	col := placeNoteUtil.GetDbCollection(postCategoriesCollectionName)
	var docs []modelDb.PostCategoriesEntity

	matchStage := bson.M{"$match": bson.M{
		"create_user_account_id": userAccountId},
	}
	projectStage := bson.M{"$project": bson.M{
		"_id":                1,
		"name":               1,
		"parent_category_id": 1,
		"memo":               1,
		"display_order":      1,
	}}
	sortStage := bson.M{"$sort": bson.M{
		"display_order": 1},
	}

	cur, err := col.Aggregate(context.Background(),
		[]bson.M{
			matchStage, projectStage, sortStage,
		})
	if err != nil {
		return docs, connect.NewError(connect.CodeInternal, err)
	}

	for cur.Next(context.Background()) {
		var doc modelDb.PostCategoriesEntity
		if err = cur.Decode(&doc); err != nil {
			return nil, connect.NewError(connect.CodeInternal, err)
		}
		docs = append(docs, doc)
	}
	return docs, nil
}
