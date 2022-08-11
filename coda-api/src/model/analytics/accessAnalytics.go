package analyticsModel

// AccessAnalytics アクセス分析データ
type AccessAnalytics struct {
	ID                     string `json:"_id" bson:"_id"`
	TotalUserCount         int64  `json:"total_user_count" bson:"total_user_count"`
	LoginUserCount         int64  `json:"login_user_count" bson:"login_user_count"`
	AccessUserCount        int64  `json:"access_user_count" bson:"access_user_count"`
	InstagramFollowerCount int64  `json:"instagram_follower_count" bson:"instagram_follower_count"`
	TwitterFollowerCount   int64  `json:"twitter_follower_count" bson:"twitter_follower_count"`
}
