package analytics

import (
	analyticsModel "coda-api/src/model/analytics"
	coordinateRepository "coda-api/src/repository/coordinate"
)

// GetCoordinatePostAnalysis コーデ投稿の参照
func GetCoordinatePostAnalysis(analysisSpan string) ([]analyticsModel.CoordinatePostAnalytics, error) {
	return coordinateRepository.GetCoordinatePostAnalysis(analysisSpan)
}
