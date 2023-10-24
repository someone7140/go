package repository

import (
	"context"
	modelDb "placeNote/src/model/db"
	"placeNote/src/placeNoteUtil"

	"github.com/bufbuild/connect-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const userAccountsCollectionName = "user_accounts"

// GetUserAccountByUserAccountIdRepository userAccountのIDによるユーザ取得
func GetUserAccountByUserAccountIdRepository(userAccountId string) (*modelDb.UserAccountsEntity, *connect.Error) {
	col := placeNoteUtil.GetDbCollection(userAccountsCollectionName)

	var result modelDb.UserAccountsEntity
	filter := bson.M{"_id": userAccountId}
	err := col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &result, nil
}

// GetUserAccountByUserSettingIdRepository userSettingIdによるユーザ取得
func GetUserAccountByUserSettingIdRepository(userSettingId string) (*modelDb.UserAccountsEntity, *connect.Error) {
	col := placeNoteUtil.GetDbCollection(userAccountsCollectionName)

	var result modelDb.UserAccountsEntity
	filter := bson.M{"user_setting_id": userSettingId}
	err := col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &result, nil
}

// GetUserAccountByGmailRepository gmailによるユーザ取得
func GetUserAccountByGmailRepository(gmail string) (*modelDb.UserAccountsEntity, *connect.Error) {
	col := placeNoteUtil.GetDbCollection(userAccountsCollectionName)

	var result modelDb.UserAccountsEntity
	filter := bson.M{"gmail": gmail}
	err := col.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, connect.NewError(connect.CodeNotFound, err)
		}
		return nil, connect.NewError(connect.CodeInternal, err)
	}
	return &result, nil
}

// NewRegistrationUserAccountRepository ユーザの新規登録
func NewRegistrationUserAccountRepository(userEntity modelDb.UserAccountsEntity) *connect.Error {
	col := placeNoteUtil.GetDbCollection(userAccountsCollectionName)
	_, err := col.InsertOne(context.Background(), userEntity)
	if err != nil {
		return connect.NewError(connect.CodeInternal, err)
	}
	return nil
}
