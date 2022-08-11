package analytics

import (
	"coda-api/src/service/analytics"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCoordinatePostAccessAnalysisData ショップ投稿の分析データの取得
func GetCoordinatePostAccessAnalysisData(c *gin.Context) {
	analysisSpan := c.Query("analysis_span")
	results, err := analytics.GetCoordinatePostAnalysis(analysisSpan)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}
