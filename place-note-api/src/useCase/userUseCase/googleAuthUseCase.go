package userUseCase

import (
	"context"
	"errors"
	"os"

	placeNote "placeNote/src/gen/proto"
	"placeNote/src/repository"

	"github.com/bufbuild/connect-go"
	"golang.org/x/oauth2/google"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

const GmailJwtPropertyName = "gmail"

func AuthGoogleUserAccount(authCode string) (*connect.Response[placeNote.AuthGoogleAccountResponse], *connect.Error) {
	userInfo, err := GetGoogleUserProfileFromAuthCode(authCode)
	if err != nil {
		return nil, err
	}
	// gmailの重複チェック
	gmail := userInfo.Email
	_, err = repository.GetUserAccountByGmailRepository(gmail)
	// ユーザが取得できている（すでにgmailが登録済み）の場合はエラー
	if err == nil {
		return nil, connect.NewError(connect.CodeAlreadyExists, errors.New("GmailAlreadyExists"))
	} else if err.Code() != connect.CodeNotFound {
		return nil, err
	}

	// トークンの生成
	token, err := GetGmailAuthToken(gmail)
	if err != nil {
		return nil, err
	}
	return connect.NewResponse(&placeNote.AuthGoogleAccountResponse{
		Token: token,
	}), nil
}

func GetGoogleUserProfileFromAuthCode(authCode string) (*v2.Tokeninfo, *connect.Error) {
	cxt := context.Background()
	credentialFilePath := "../googleCredential/" + os.Getenv("GOOGLE_CREDENTIAL_FILE")
	b, err := os.ReadFile(credentialFilePath)
	if err != nil {
		return nil, connect.NewError(connect.CodeInternal, errors.New("credentialファイル読み込みエラー"))
	}

	config, err := google.ConfigFromJSON(b, "email profile openid")
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("Google認証エラー"))
	}
	token, _ := config.Exchange(cxt, authCode)
	if token == nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("Google認証エラー"))
	}

	service, err := v2.NewService(cxt, option.WithTokenSource(config.TokenSource(cxt, token)))
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("Google認証エラー"))
	}

	userInfo, err := service.Tokeninfo().AccessToken(token.AccessToken).Context(cxt).Do()
	if err != nil {
		return nil, connect.NewError(connect.CodeUnauthenticated, errors.New("Google認証エラー"))
	}

	return userInfo, nil
}

// GetGmailAuthToken Google認証後のgmailトークンを生成
func GetGmailAuthToken(gmail string) (string, *connect.Error) {
	return GetJwtTokenAuth(GmailJwtPropertyName, gmail, GetJwtAuthDayExpire(1))
}

// GetGmailFromToken Google認証後のgmailトークン複合
func GetGmailFromToken(token string) (string, *connect.Error) {
	return GetStrInfoFromToken(GmailJwtPropertyName, token)
}
