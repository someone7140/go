package analytics

import (
	analyticsModel "coda-api/src/model/analytics"
	analyticsRepository "coda-api/src/repository/analytics"
)

// GerAccessAnalysis アクセス情報の収集サービス
func GerAccessAnalysis() ([]analyticsModel.AccessAnalytics, error) {
	return analyticsRepository.GetDateAccessAnalytics(200)
}
