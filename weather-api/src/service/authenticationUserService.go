package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"
	"weather-api/src/pb"

	"weather-api/src/db/repository"

	jwt "github.com/golang-jwt/jwt/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	v2Oauth "google.golang.org/api/oauth2/v2"
	"google.golang.org/api/option"
	"xorm.io/xorm"
)

type AuthenticationUserService struct {
	dbEngine *xorm.Engine
}

const USER_ID_CONTEXT_KEY string = "userId"

func (s *AuthenticationUserService) VerifyGoogleAuthCode(ctx context.Context, r *pb.VerifyGoogleAuthCodeRequest) (*pb.UserResponse, error) {
	if ctx.Err() == context.Canceled {
		return &pb.UserResponse{}, fmt.Errorf("client cancelled: abandoning")
	}

	// 認証コードからユーザ情報取得してDB登録
	userId, userName := getUserInforomGoogleAuthCode(r.GoogleAuthCode)
	if userId == "" {
		return nil, errors.New("Can not get user")
	}
	err := repository.RegsiterUser(s.dbEngine, userId, userName)
	if err != nil {
		return nil, err
	}

	// 新しくjwtトークンを作成
	authToken, err := getNewJwtToken(userId)
	if err != nil {
		return nil, err
	}
	return &pb.UserResponse{
		Id:        userId,
		Name:      userName,
		AuthToken: authToken,
	}, nil
}

// getUserInforomGoogleAuthCode コードが認証できたらユーザidとユーザ名を返す
func getUserInforomGoogleAuthCode(authCode string) (string, string) {
	ctx := context.Background()

	conf := oauth2.Config{
		ClientID:     os.Getenv("GCP_AUTH_CLIENT_ID"),
		ClientSecret: os.Getenv("GCP_AUTH_CLIENT_SECRET"),
		Scopes:       []string{"profile", "email"},
		Endpoint:     google.Endpoint,
		RedirectURL:  os.Getenv("VIEW_DOMAIN"),
	}
	token, err := conf.Exchange(ctx, authCode)
	if err != nil {
		return "", ""
	}
	oauth2Service, err := v2Oauth.NewService(ctx, option.WithTokenSource(conf.TokenSource(ctx, token)))
	userInfo, err := oauth2Service.Userinfo.Get().Do()
	if err != nil {
		return "", ""
	}
	return userInfo.Id, userInfo.Name
}

func (s *AuthenticationUserService) VerifyAuthToken(ctx context.Context, r *pb.VerifyAuthTokenRequest) (*pb.AuthTokenResponse, error) {
	if ctx.Err() == context.Canceled {
		return &pb.AuthTokenResponse{}, fmt.Errorf("client cancelled: abandoning")
	}

	// contextからユーザID取得
	userId := GetUserIdFromContext(ctx)
	if userId == "" {
		return nil, errors.New("Can not get user")
	}

	// ユーザ情報取得
	user, err := repository.SelectUser(s.dbEngine, userId)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("Can not get user from DB")
	}

	return &pb.AuthTokenResponse{
		Id:   user.Id,
		Name: user.Name,
	}, nil
}

// getNewJwtToken 新しいjwtトークンを取得
func getNewJwtToken(userId string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		"id":  userId,
		"exp": getAuthExpire(),
	}
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

// getAuthExpire 認証期限の取得
func getAuthExpire() int64 {
	return time.Now().UTC().AddDate(0, 3, 0).Unix()
}

// GetUserIdInfoFromJwtToken jwtトークンからユーザIDを取得
func GetUserIdInfoFromJwtToken(jwtToken string) (string, error) {
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		b := []byte(os.Getenv("JWT_SECRET"))
		return b, nil
	})
	if token == nil || err != nil {
		return "", errors.New("Can not get userId")
	}
	claims := token.Claims.(jwt.MapClaims)
	// expの判定
	expTime := time.Unix(int64(claims["exp"].(float64)), 0).UTC()
	if time.Now().UTC().After(expTime) {
		return "", errors.New("expTime")
	}
	return claims["id"].(string), nil
}

// GetUserIdSetContext ユーザIDをContextに設定
func GetUserIdSetContext(ctx context.Context, userId string) context.Context {
	return context.WithValue(ctx, USER_ID_CONTEXT_KEY, userId)
}

// GetUserIdFromContext ContextからユーザIDを取得
func GetUserIdFromContext(ctx context.Context) string {
	return ctx.Value(USER_ID_CONTEXT_KEY).(string)
}
