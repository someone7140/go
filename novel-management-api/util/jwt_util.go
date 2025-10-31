package util

import (
	"errors"
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GetJwtToken(claims jwt.MapClaims) (*string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	token.Claims = claims
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return nil, err
	}

	return &tokenStr, nil
}

func DecodeJwtToken(t string) (*jwt.MapClaims, error) {
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

	return &claims, nil
}
