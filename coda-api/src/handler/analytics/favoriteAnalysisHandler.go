package analytics

import (
	"coda-api/src/service/analytics"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetFavoriteAnalysisData いいね分析データの取得
func GetFavoriteAnalysisData(c *gin.Context) {
	results, err := analytics.GerFavoriteAnalysis(c.Query("sort"))
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}
