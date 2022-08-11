package analytics

import (
	analyticsModel "coda-api/src/model/analytics"
	postRepository "coda-api/src/repository/post"
)

// GerFavoriteAnalysis いいね情報の参照
func GerFavoriteAnalysis(sort string) ([]analyticsModel.FavoriteAnalytics, error) {
	return postRepository.GetFavoriteAnalytics(sort, 200)
}
