package analytics

import (
	analyticsModel "coda-api/src/model/analytics"
	"coda-api/src/service/analytics"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetAccessAnalysisData アクセス分析データの取得
func GetAccessAnalysisData(c *gin.Context) {
	var results []analyticsModel.AccessAnalytics
	results, err := analytics.GerAccessAnalysis()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}
