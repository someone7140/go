package instagramBatch

import (
	"coda-api/src/service/instagramBatch"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GatherInstagramPost インスタの投稿収集Handler
func GatherInstagramPost(c *gin.Context) {
	err := instagramBatch.GatherInstagramPostService()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}
