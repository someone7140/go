package coordinateModel

import "time"

// CoordinatePostInfo コーデ投稿のstruct
type CoordinatePostInfo struct {
	ID                   string            `json:"_id" bson:"_id"`
	Title                string            `json:"title" bson:"title"`
	Status               string            `json:"status" bson:"status"`
	Detail               string            `json:"detail" bson:"detail"`
	ShopID               string            `json:"shop_id" bson:"shop_id"`
	ShopName             string            `json:"shop_name" bson:"shop_name"`
	ShopSettingID        string            `json:"shop_setting_id" bson:"shop_setting_id"`
	URL                  string            `json:"url" bson:"url"`
	PostUserID           string            `json:"post_user_id" bson:"post_user_id"`
	PostDate             time.Time         `json:"post_date" bson:"post_date"`
	ModelCategory        string            `json:"model_category" bson:"model_category"`
	FavoriteCount        int               `json:"favorite_count" bson:"favorite_count"`
	FavoritedFlg         bool              `json:"favorited_flg" bson:"favorited_flg"`
	ImpressionCount      int               `json:"impression_count" bson:"impression_count"`
	ClickCount           int               `json:"click_count" bson:"click_count"`
	PurchaseRequestCount int               `json:"purchase_request_count" bson:"purchase_request_count"`
	Images               []CoordinateImage `json:"images" bson:"images"`
	ModelAttribute       ModelAttribute    `json:"model_attribute" bson:"model_attribute"`
	Price                int               `json:"price" bson:"price"`
	Category             string            `json:"category" bson:"category"`
	Sale                 *SaleResponse     `json:"sale" bson:"sale"`
}

// CoordinateImage コーデ画像
type CoordinateImage struct {
	Key int    `json:"key" bson:"key"`
	URL string `json:"url" bson:"url"`
}

// ModelAttribute モデルの属性
type ModelAttribute struct {
	Gender     string `json:"gender" bson:"gender"`
	Silhouette string `json:"silhouette" bson:"silhouette"`
	Height     int    `json:"height" bson:"height"`
	Weight     int    `json:"weight" bson:"weight"`
	Size       string `json:"size" bson:"size"`
}

// CoordinatePostSearchRequest コーデ検索のリクエスト
type CoordinatePostSearchRequest struct {
	Keyword       string `json:"keyword" bson:"keyword"`
	ShopSettingId string `json:"shop_setting_id" bson:"shop_setting_id"`
	Gender        string `json:"gender" bson:"gender"`
	Silhouette    string `json:"silhouette" bson:"silhouette"`
	MinHeight     int    `json:"min_height" bson:"min_height"`
	MaxHeight     int    `json:"max_height" bson:"max_height"`
	MinWeight     int    `json:"min_weight" bson:"min_weight"`
	MaxWeight     int    `json:"max_weight" bson:"max_weight"`
	MinPrice      int    `json:"min_price" bson:"min_price"`
	MaxPrice      int    `json:"max_price" bson:"max_price"`
	Category      string `json:"category" bson:"category"`
}

// SaleResponse セール情報のレスポンス
type SaleResponse struct {
	SalePrice int        `json:"sale_price" bson:"sale_price"`
	StartDate *time.Time `json:"start_date" bson:"start_date"`
	EndDate   *time.Time `json:"end_date" bson:"end_date"`
}
