package service

import (
	"context"
	"wasurena-task-api/domain"
	"wasurena-task-api/graph/model"
	"wasurena-task-api/middleware"

	"database/sql"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func CreateUserAccount(ctx context.Context, input model.NewUserAccount) (bool, error) {
	// 重複してるかチェック
	_, err := middleware.GetDbQueries(ctx).SelectUserAccountByUserSettingId(ctx, input.UserSettingID)
	if err != nil {
		if err != sql.ErrNoRows {
			return false, err
		}
	} else {
		// ユーザが取得できていたら重複エラー
		return false, &gqlerror.Error{
			Message: "Dupilicate userSettingId",
			Extensions: map[string]interface{}{
				"code": "400",
			}}
	}

	// トークンからLINEのユーザを取得
	return true, err
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
