package service

import (
	"context"
	"fmt"
	"os"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"
	"wasurena-task-api/domain"
	"wasurena-task-api/graph/model"

	"github.com/jackc/pgx/v5"
	"github.com/rs/xid"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

// ユーザーのアカウントを追加する
func CreateUserAccount(ctx context.Context, input model.NewUserAccount) (*model.UserAccountResponse, error) {
	// UserSettingIDが重複してるかチェック
	_, err := custom_middleware.GetDbQueries(ctx).SelectUserAccountByUserSettingId(ctx, input.UserSettingID)
	if err != nil {
		if err != pgx.ErrNoRows {
			return nil, err
		}
	} else {
		// ユーザが取得できていたら重複エラー
		return nil, &gqlerror.Error{
			Message: "Dupilicate userSettingId",
			Extensions: map[string]any{
				"code": 400,
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
		if err != pgx.ErrNoRows {
			return nil, err
		}
	} else {
		// ユーザが取得できていたら重複エラー
		return nil, &gqlerror.Error{
			Message: "Dupilicate lineId",
			Extensions: map[string]any{
				"code": 403,
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

// LINEの認証コードから登録用トークンを返す
func GetUserRegisterTokenFromLineAuthCode(ctx context.Context, lineAuthCode string) (*model.CreateUserRegisterTokenResponse, error) {
	// 認証コードからLINEのユーザ情報を取得しトークン化する
	lineInfo, err := domain.GetLineUserInfoFromAuthCode(
		lineAuthCode,
		fmt.Sprintf("%s%s", os.Getenv("FRONTEND_DOMAIN"), os.Getenv("LINE_AUTH_REGISTER_REDIRECT_PATH")))
	if err != nil {
		return nil, err
	}

	_, err = custom_middleware.GetDbQueries(ctx).SelectUserAccountByLineId(ctx, lineInfo.UserID)
	if err != nil {
		if err != pgx.ErrNoRows {
			return nil, err
		}
	} else {
		// ユーザが取得できていたら重複エラー
		return nil, &gqlerror.Error{
			Message: "Dupilicate lineId",
			Extensions: map[string]any{
				"code": 403,
			}}
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

// Contextの認証情報からユーザ情報を返す
func GetUserAccountFromContext(ctx context.Context) (*model.UserAccountResponse, error) {
	userAccountId := custom_middleware.GeUserAccountId(ctx)
	userAccount, err := custom_middleware.GetDbQueries(ctx).SelectUserAccountById(ctx, *userAccountId)
	if err != nil {
		return nil, err
	}

	return &model.UserAccountResponse{
		Token:           "", // tokenは新たには発行しない
		UserSettingID:   userAccount.UserSettingID,
		UserName:        userAccount.UserName,
		ImageURL:        userAccount.ImageUrl,
		IsLineBotFollow: userAccount.IsLineBotFollow,
	}, nil
}

// LINEの認証コードからユーザ情報を取得
func GetUserAccountFromLineAuthCode(ctx context.Context, lineAuthCode string) (*model.UserAccountResponse, error) {
	// 認証コードからLINEのユーザ情報を取得
	lineInfo, err := domain.GetLineUserInfoFromAuthCode(
		lineAuthCode,
		fmt.Sprintf("%s%s", os.Getenv("FRONTEND_DOMAIN"), os.Getenv("LINE_AUTH_LOGIN_REDIRECT_PATH")))
	if err != nil {
		return nil, err
	}

	// DBからユーザ情報取得
	userAccountDb, err := custom_middleware.GetDbQueries(ctx).SelectUserAccountByLineId(ctx, lineInfo.UserID)
	if err != nil {
		if err != pgx.ErrNoRows {
			return nil, &gqlerror.Error{
				Message: "Can not find userAccount",
				Extensions: map[string]any{
					"code": 404,
				}}
		} else {
			return nil, err
		}
	}

	// イメージ画像のURLを更新
	_, err = custom_middleware.GetDbQueries(ctx).UpdateUserImageUrl(ctx, db.UpdateUserImageUrlParams{
		ID:       userAccountDb.ID,
		ImageUrl: &lineInfo.PictureURL,
	})
	if err != nil {
		return nil, err
	}
	userAccountDb.ImageUrl = &lineInfo.PictureURL

	// レスポンス構築
	userAccount := domain.UserAccount(userAccountDb)
	userToken, err := userAccount.GetAccountUserToken()
	if err != nil {
		return nil, err
	}

	return &model.UserAccountResponse{
		Token:           userToken,
		UserSettingID:   userAccount.UserSettingID,
		UserName:        userAccount.UserName,
		ImageURL:        userAccount.ImageUrl,
		IsLineBotFollow: userAccount.IsLineBotFollow,
	}, nil
}
