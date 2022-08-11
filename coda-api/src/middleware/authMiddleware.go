package middleware

import (
	"coda-api/src/service/auth"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LoginCheckMiddleware ログインチェック
func LoginCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginInfo, err := auth.GetLoginInfoInfoFromJwtToken(c)
		if loginInfo == nil || err != nil {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			// コンテキストにログイン情報のJsonをセット
			err = auth.SetLoginInfoJSONToContext(c, loginInfo)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				c.Abort()
			}
			c.Next()
		}
	}
}

// LoginCheckAdminMiddleware ADMIN用のログインチェック
func LoginCheckAdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		loginInfo, err := auth.GetLoginInfoInfoFromJwtToken(c)
		if loginInfo == nil || err != nil || loginInfo.UserType != "admin" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		} else {
			// コンテキストにログイン情報のJsonをセット
			err = auth.SetLoginInfoJSONToContext(c, loginInfo)
			if err != nil {
				c.Status(http.StatusInternalServerError)
				c.Abort()
			}
			c.Next()
		}
	}
}
