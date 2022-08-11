package analyticsModel

import "time"

// CoordinatePostAnalytics コーデ投稿の分析
type CoordinatePostAnalytics struct {
	ID                   string    `json:"_id" bson:"_id"`
	Title                string    `json:"title" bson:"title"`
	ShopID               string    `json:"shop_id" bson:"shop_id"`
	ShopName             string    `json:"shop_name" bson:"shop_name"`
	ShopSettingID        string    `json:"shop_setting_id" bson:"shop_setting_id"`
	PostDate             time.Time `json:"post_date" bson:"post_date"`
	ImpressionCount      int       `json:"impression_count" bson:"impression_count"`
	ClickCount           int       `json:"click_count" bson:"click_count"`
	PurchaseRequestCount int       `json:"purchase_request_count" bson:"purchase_request_count"`
}
