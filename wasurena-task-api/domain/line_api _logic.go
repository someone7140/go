package domain

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
	"golang.org/x/oauth2"
)

type LineUserInfo struct {
	UserID      string `json:"userId"`
	DisplayName string `json:"displayName"`
	PictureURL  string `json:"pictureUrl"`
}

// 認証コードからLINEのユーザ情報を取得
func GetLineUserInfoFromAuthCode(code string, redirectUrl string) (*LineUserInfo, error) {
	ctx := context.Background()
	oauth2Config := &oauth2.Config{
		ClientID:     os.Getenv("LINE_CLIENT_ID"),
		ClientSecret: os.Getenv("LINE_SECRET_ID"),
		Scopes:       []string{"profile", "openid"},
		RedirectURL:  redirectUrl,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://access.line.me/oauth2/v2.1/authorize",
			TokenURL: "https://api.line.me/oauth2/v2.1/token",
		},
	}
	token, err := oauth2Config.Exchange(ctx, code)
	if err != nil {
		return nil, err
	}

	// Call User Profile Endpoint
	client := oauth2Config.Client(context.Background(), token)
	resp, err := client.Get("https://api.line.me/v2/profile")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var userInfo LineUserInfo
	if err := json.Unmarshal(data, &userInfo); err != nil {
		return nil, err
	}

	return &userInfo, err
}

// トークンからLINEのユーザ情報を取得
func GetLineUserInfoFromToken(tokenInput string) (*LineUserInfo, error) {
	token, err := jwt.Parse(tokenInput, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}

	claims := token.Claims.(jwt.MapClaims)
	exp, err := claims.GetExpirationTime()
	if claims == nil || exp == nil || err != nil {
		return nil, errors.New("can not get token info")
	}
	if exp.Unix() < time.Now().Unix() {
		return nil, errors.New("expire token")
	}

	return &LineUserInfo{
		UserID:      claims["userId"].(string),
		DisplayName: claims["displayName"].(string),
		PictureURL:  claims["pictureUrl"].(string),
	}, nil
}

// LINEのユーザ情報をトークン化して返す
func (u LineUserInfo) GetLineUserToken() (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// 期限は3時間にする
	expTime := time.Now().UTC().Add(3 * time.Hour).Unix()
	token.Claims = jwt.MapClaims{
		"userId":      u.UserID,
		"displayName": u.DisplayName,
		"pictureUrl":  u.PictureURL,
		"exp":         expTime,
	}

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
