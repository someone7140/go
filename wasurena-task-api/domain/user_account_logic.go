package domain

import (
	"errors"
	"os"
	"time"
	"wasurena-task-api/db"

	jwt "github.com/golang-jwt/jwt/v5"
)

type UserAccount db.UserAccount

// トークンからユーザIDを取得
func GetUserAccountIdFromToken(t string) (*string, error) {
	token, err := jwt.Parse(string(t), func(token *jwt.Token) (any, error) {
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

	userAccountId := claims["userId"].(string)
	return &userAccountId, nil
}

// ユーザIDをトークン化して返す
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
