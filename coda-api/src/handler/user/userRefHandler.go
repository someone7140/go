package user

import (
	"coda-api/src/service/auth"
	"coda-api/src/service/user"
	"net/http"

	errorConstants "coda-api/src/constants"

	"github.com/gin-gonic/gin"
)

// GetMyUserInfoDetail ユーザの詳細を取得
func GetMyUserInfoDetail(c *gin.Context) {
	loginInfo := auth.GetLoginInfoFromContext(c)
	response, err := user.GetUserInfoDetail(loginInfo.ID)
	if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, response)
	}
}
