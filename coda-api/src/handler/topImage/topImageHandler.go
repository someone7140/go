package topImage

import (
	"coda-api/src/service/auth"
	"coda-api/src/service/topImage"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetTopRecentImages 最新順でトップ用の画像を取得
func GetTopRecentImages(c *gin.Context) {
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	results, err := topImage.GetTopRecentImages(loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}
