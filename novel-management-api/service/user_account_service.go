package service

import (
	"context"
	"errors"
	"main/db/db_model"
	"main/graph/graphql_model"
	"main/repository"
	"main/util"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

func GetUserAccountRegisterTokenByGoogleAuthCode(authCode string) (*string, error) {
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
	_, err = repository.GetUserAccountByGmail(ctx, googleUserInfo.Email)
	err = checkErrorUserAccountAlreadyRegistered(err)
	if err != nil {
		return nil, err
	}

	// ユーザー登録用のjwtトークンを生成
	claims := jwt.MapClaims{
		"gmail":    googleUserInfo.Email,
		"imageUrl": googleUserInfo.Picture,
		"exp":      time.Now().UTC().Add(3 * time.Hour).Unix(),
	}
	return util.GetJwtToken(claims)
}

func GetUserAccountByGoogleAuthCode(authCode string) (*graphql_model.UserAccountResponse, error) {
	ctx := context.Background()
	googleUserInfo, err := util.GetGoogleUserProfileFromAuthCode(ctx, authCode)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Extensions: map[string]any{
				"code": 401,
			}}
	}

	userAccount, err := repository.GetUserAccountByGmail(ctx, googleUserInfo.Email)
	err = checkErrorUserAccountNotFound(err)
	if err != nil {
		return nil, err
	}
	authToken, err := GenerateAuthToken(userAccount.ID)
	if err != nil {
		return nil, err
	}

	return &graphql_model.UserAccountResponse{
		Token:         *authToken,
		UserSettingID: userAccount.UserSettingID,
		Name:          userAccount.Name,
		ImageURL:      *userAccount.ImageURL,
	}, nil
}

func GetUserAccountByUserAccountID(userAccountID string) (*graphql_model.UserAccountResponse, error) {
	ctx := context.Background()
	userAccount, err := repository.GetUserAccountByID(ctx, userAccountID)
	err = checkErrorUserAccountNotFound(err)
	if err != nil {
		return nil, err
	}
	authToken, err := GenerateAuthToken(userAccount.ID)
	if err != nil {
		return nil, err
	}

	return &graphql_model.UserAccountResponse{
		Token:         *authToken,
		UserSettingID: userAccount.UserSettingID,
		Name:          userAccount.Name,
		ImageURL:      *userAccount.ImageURL,
	}, nil
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

	// 該当のgmailアカウントかuserSettingIDで既に登録があるかチェック
	_, err = repository.GetUserAccountByGmail(ctx, gmail)
	err = checkErrorUserAccountAlreadyRegistered(err)
	if err != nil {
		return nil, err
	}
	_, err = repository.GetUserAccountByUserSettingID(ctx, userSettingID)
	err = checkErrorUserAccountAlreadyRegistered(err)
	if err != nil {
		return nil, err
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
	authToken, err := GenerateAuthToken(uid.String())
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

func GenerateAuthToken(userID string) (*string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().UTC().AddDate(0, 3, 0).Unix(),
	}

	return util.GetJwtToken(claims)
}

func DecodeAuthToken(token string) (*string, error) {
	claims, err := util.DecodeJwtToken(token)
	if err != nil {
		return nil, &gqlerror.Error{
			Message: err.Error(),
			Extensions: map[string]any{
				"code": 401,
			}}
	}
	userID := (*claims)["userID"].(string)
	return &userID, nil
}

func checkErrorUserAccountAlreadyRegistered(err error) error {
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return err
		}
	} else {
		return &gqlerror.Error{
			Message: "Already registered user account",
			Extensions: map[string]any{
				"code": 403,
			}}
	}
	return nil
}

func checkErrorUserAccountNotFound(err error) error {
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &gqlerror.Error{
				Message: "User account not found",
				Extensions: map[string]any{
					"code": 404,
				}}
		}
		return err
	}
	return nil
}
