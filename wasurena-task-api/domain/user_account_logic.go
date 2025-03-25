package domain

import (
	"context"
	"errors"
	"os"
	"time"
	"wasurena-task-api/custom_middleware"
	"wasurena-task-api/db"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserAccount db.UserAccount

// トークンからユーザ情報を取得
func GetUserAccountInfoFromToken(ctx context.Context, tokenInput string) (*UserAccount, error) {
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

	// userIdをキーにクエリ発行
	userAccountDb, err := custom_middleware.GetDbQueries(ctx).SelectUserAccountById(ctx, claims["userId"].(string))
	if err != nil {
		return nil, err
	}
	userAccount := UserAccount(userAccountDb)

	return &userAccount, nil
}

// ユーザ情報をトークン化して返す
func (u UserAccount) GetAccountUserToken() (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// 期限は3ヶ月にする
	expTime := time.Now().UTC().AddDate(0, 3, 0).Unix()
	token.Claims = jwt.MapClaims{
		"userId": u.ID,
		"exp":    expTime,
	}

	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}
