package analytics

import (
	"coda-api/src/service/analytics"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ExecuteAccessAnalysis アクセス情報の収集Handler
func ExecuteAccessAnalysis(c *gin.Context) {
	err := analytics.AccessAnalysisBatchService()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}
