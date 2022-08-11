package coordinateModel

import "time"

// ShopInfo ショップ情報
type ShopInfo struct {
	ID             string         `json:"_id" bson:"_id"`
	Name           string         `json:"name" bson:"name"`
	ShopSettingID  string         `json:"shop_setting_id" bson:"shop_setting_id"`
	ShopURL        string         `json:"shop_url" bson:"shop_url"`
	Detail         string         `json:"detail" bson:"detail"`
	CreateDate     time.Time      `json:"create_date" bson:"create_date"`
	ShopTaegetUser ShopTaegetUser `json:"shop_target_user" bson:"shop_target_user"`
}

// ShopTaegetUser ショップのユーザ層
type ShopTaegetUser struct {
	Gender     string `json:"gender" bson:"gender"`
	Silhouette string `json:"silhouette" bson:"silhouette"`
	MinHeight  int    `json:"min_height" bson:"min_height"`
	MaxHeight  int    `json:"max_height" bson:"max_height"`
	MinWeight  int    `json:"min_weight" bson:"min_weight"`
	MaxWeight  int    `json:"max_weight" bson:"max_weight"`
}

// ShopDeleteRequest ショップ削除のリクエスト
type ShopDeleteRequest struct {
	ShopSettingID string `json:"shop_setting_id" bson:"shop_setting_id"`
}
