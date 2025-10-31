package service

import (
	"context"
	"main/db/db_model"
	"main/graph/graphql_model"
	"main/repository"
	"main/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func GetUserAccountRegisterToken(authCode string) (*string, error) {
	ctx := context.Background()
	googleUserInfo, err := util.GetGoogleUserProfileFromAuthCode(ctx, authCode)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Extensions: map[string]any{
				"code": 401,
			}}
	}
	// 該当のgmailアカウントで既に登録があるかチェック
	userAccount, err := repository.GetUserAccountByGmail(ctx, googleUserInfo.Email)
	if err != nil {
		return nil, err
	}
	if userAccount != nil {
		return nil, &gqlerror.Error{
			Message: "Already registered gmail user",
			Extensions: map[string]any{
				"code": 403,
			}}
	}

	// ユーザー登録用のjwtトークンを生成
	claims := jwt.MapClaims{
		"gmail":    googleUserInfo.Email,
		"imageUrl": googleUserInfo.Picture,
		"exp":      time.Now().UTC().Add(3 * time.Hour).Unix(),
	}
	return util.GetJwtToken(claims)
}

func AddUserAccountByGoogleAuth(registerToken string, userSettingID string, name string) (*graphql_model.UserAccountResponse, error) {
	ctx := context.Background()
	// 登録用のトークンからgoogleの認証情報を取得
	claims, err := util.DecodeJwtToken(registerToken)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Extensions: map[string]any{
				"code": 401,
			}}
	}
	gmail := (*claims)["gmail"].(string)
	imageUrl := (*claims)["imageUrl"].(string)

	// 該当のgmailアカウントで既に登録があるかチェック
	userAccount, err := repository.GetUserAccountByGmail(ctx, gmail)
	if err != nil {
		return nil, err
	}
	if userAccount != nil {
		return nil, &gqlerror.Error{
			Message: "Already registered gmail user",
			Extensions: map[string]any{
				"code": 403,
			}}
	}

	// DBに登録
	uid, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}
	userAccountDbModel := db_model.UserAccount{
		ID:            uid.String(),
		UserSettingID: userSettingID,
		Name:          name,
		Gmail:         gmail,
		ImageURL:      &imageUrl,
		CreatedAt:     time.Now(),
	}
	err = repository.AddUserAccount(ctx, userAccountDbModel)
	if err != nil {
		return nil, err
	}
	authToken, err := generateAuthToken(uid.String())
	if err != nil {
		return nil, err
	}

	return &graphql_model.UserAccountResponse{
		Token:         *authToken,
		UserSettingID: userSettingID,
		Name:          name,
		ImageURL:      imageUrl,
	}, nil
}

func generateAuthToken(userID string) (*string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().UTC().AddDate(0, 3, 0).Unix(),
	}

	return util.GetJwtToken(claims)
}
