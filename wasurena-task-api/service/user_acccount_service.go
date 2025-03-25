package service

import (
	"context"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/domain"
	"wasurena-task-api/graph/model"

	"database/sql"

	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CreateUserAccount(ctx context.Context, input model.NewUserAccount) (*model.UserAccountResponse, error) {
	// UserSettingIDが重複してるかチェック
	_, err := custom_middleware.GetDbQueries(ctx).SelectUserAccountByUserSettingId(ctx, input.UserSettingID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		// ユーザが取得できていたら重複エラー
		return nil, &gqlerror.Error{
			Message: "Dupilicate userSettingId",
			Extensions: map[string]interface{}{
				"code": "400",
			}}
	}

	// トークンからLINEのユーザを取得
	lineUser, err := domain.GetLineUserInfoFromToken(input.AuthToken)
	if err != nil {
		return nil, err
	}
	// LINE_IDの重複チェック
	_, err = custom_middleware.GetDbQueries(ctx).SelectUserAccountByLineId(ctx, lineUser.UserID)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	} else {
		// ユーザが取得できていたら重複エラー
		return nil, &gqlerror.Error{
			Message: "Dupilicate lineId",
			Extensions: map[string]interface{}{
				"code": "403",
			}}
	}

	// DBに登録
	createData := db.CreateUserAccountParams{
		ID:            xid.New().String(),
		UserSettingID: input.UserSettingID,
		LineID:        lineUser.UserID,
		UserName:      input.UserName,
		ImageUrl:      &lineUser.PictureURL,
	}
	createdDbUser, err := custom_middleware.GetDbQueries(ctx).CreateUserAccount(ctx, createData)
	if err != nil {
		return nil, err
	}

	// レスポンス構築
	createdUser := domain.UserAccount(createdDbUser)
	userToken, err := createdUser.GetAccountUserToken()
	if err != nil {
		return nil, err
	}

	return &model.UserAccountResponse{
		Token:           userToken,
		UserSettingID:   createdUser.UserSettingID,
		UserName:        createdUser.UserName,
		ImageURL:        createdUser.ImageUrl,
		IsLineBotFollow: createdUser.IsLineBotFollow,
	}, err
}

func GetUserRegisterToken(ctx context.Context, lineAuthCode string) (*model.CreateUserRegisterTokenResponse, error) {
	// 認証コードからLINEのユーザ情報を取得しトークン化する
	lineInfo, err := domain.GetLineUserInfoFromAuthCode(lineAuthCode)
	if err != nil {
		return nil, err
	}
	token, err := lineInfo.GetLineUserToken()
	if err != nil {
		return nil, err
	}

	return &model.CreateUserRegisterTokenResponse{
		Token:    token,
		LineName: lineInfo.DisplayName,
	}, err
}
