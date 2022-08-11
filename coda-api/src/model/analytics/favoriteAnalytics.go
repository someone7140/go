package analyticsModel

import "time"

// FavoriteAnalytics いいね分析データ
type FavoriteAnalytics struct {
	PostID                 string                    `json:"post_id" bson:"post_id"`
	Status                 string                    `json:"status" bson:"status"`
	PostInstagramAccountId string                    `json:"post_instagram_account_id" bson:"post_instagram_account_id"`
	PostInstagramUserName  string                    `json:"instagram_user_name" bson:"instagram_user_name"`
	Category               string                    `json:"category" bson:"category"`
	ContentURL             string                    `json:"content_url" bson:"content_url"`
	PostDate               time.Time                 `json:"post_date" bson:"post_date"`
	FavoriteTotalCount     int64                     `json:"favorite_total_count" bson:"favorite_total_count"`
	FavoriteDetails        []FavoriteAnalyticsDetail `json:"favorite_analytics_details" bson:"favorite_analytics_details"`
}

// FavoriteAnalyticsDetail いいね分析の詳細データ
type FavoriteAnalyticsDetail struct {
	UserCategory  string `json:"user_category" bson:"user_category"`
	FavoriteCount int    `json:"favorite_count" bson:"favorite_count"`
}
