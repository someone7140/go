package userUseCase

import (
	"errors"
	"os"
	"time"

	"github.com/bufbuild/connect-go"
	jwt "github.com/golang-jwt/jwt/v5"
)

// GetJwtTokenAuth 認証用のJwtトークンを生成
func GetJwtTokenAuth(propertyName string, propertyValue string, inputExp int64) (string, *connect.Error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = jwt.MapClaims{
		propertyName: propertyValue,
		"exp":        inputExp,
	}
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", connect.NewError(connect.CodeInternal, err)
	}
	return tokenStr, nil
}

// GetJwtAuthMonthExpire 月単位の認証期限取得
func GetJwtAuthMonthExpire(month int) int64 {
	return time.Now().UTC().AddDate(0, month, 0).Unix()
}

// GetJwtAuthDayExpire 日単位の認証期限取得
func GetJwtAuthDayExpire(day int) int64 {
	return time.Now().UTC().AddDate(0, 0, day).Unix()
}

// GetStrInfoFromToken トークンから設定した文字列を複合
func GetStrInfoFromToken(propertyName string, tokenSrt string) (string, *connect.Error) {
	token, err := jwt.Parse(tokenSrt, func(token *jwt.Token) (interface{}, error) {
		return []byte("SECRET_KEY"), nil
	})
	if err != nil {
		return "", connect.NewError(connect.CodeUnauthenticated, errors.New("can not get token info"))
	}

	claims := token.Claims.(jwt.MapClaims)
	exp, err := claims.GetExpirationTime()
	if claims == nil || exp == nil || err != nil {
		return "", connect.NewError(connect.CodeUnauthenticated, errors.New("can not get token info"))
	}
	if exp.Unix() < time.Now().Unix() {
		return "", connect.NewError(connect.CodeUnauthenticated, errors.New("can not get token info"))
	}

	return claims[propertyName].(string), nil
}
