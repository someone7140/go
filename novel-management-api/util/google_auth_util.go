package util

import (
	"context"
	"os"

	"golang.org/x/oauth2/google"
	"google.golang.org/api/oauth2/v2"
	v2 "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
)

func GetGoogleUserProfileFromAuthCode(ctx context.Context, authCode string) (*v2.Userinfo, error) {
	credentialFilePath := "google_credential/" + os.Getenv("GOOGLE_CREDENTIAL_FILE")
	b, err := os.ReadFile(credentialFilePath)
	if err != nil {
		return nil, err
	}

	config, err := google.ConfigFromJSON(b, "email profile openid")
	if err != nil {
		return nil, err
	}
	config.RedirectURL = os.Getenv("FRONTEND_DOMAIN")
	token, err := config.Exchange(ctx, authCode)
	if token == nil || err != nil {
		return nil, err
	}

	srv, err := oauth2.NewService(ctx, option.WithTokenSource(config.TokenSource(ctx, token)))
	if err != nil {
		return nil, err
	}
	userInfo, err := srv.Userinfo.Get().Do()
	if err != nil {
		return nil, err
	}

	return userInfo, nil
}
