package itemPostModel

import "time"

// ItemPostResponse アイテム投稿のresponse用struct
type ItemPostResponse struct {
	ID              string    `json:"_id" bson:"_id"`
	Title           string    `json:"title" bson:"title"`
	Detail          string    `json:"detail" bson:"detail"`
	ItemType        string    `json:"item_type" bson:"item_type"`
	URL             string    `json:"url" bson:"url"`
	PostDate        time.Time `json:"post_date" bson:"post_date"`
	UserID          string    `json:"user_id" bson:"user_id"`
	UserSettingID   string    `json:"user_setting_id" bson:"user_setting_id"`
	UserName        string    `json:"user_name" bson:"user_name"`
	UserIconURL     string    `json:"user_icon_url" bson:"user_icon_url"`
	Category        string    `json:"category" bson:"category"`
	Gender          string    `json:"gender" bson:"gender"`
	Silhouette      string    `json:"silhouette" bson:"silhouette"`
	Complex         string    `json:"complex" bson:"complex"`
	ImageURL        string    `json:"image_url" bson:"image_url"`
	FavoriteCount   int       `json:"favorite_count" bson:"favorite_count"`
	FavoritedFlg    bool      `json:"favorited_flg" bson:"favorited_flg"`
	ImpressionCount int       `json:"impression_count" bson:"impression_count"`
	ClickCount      int       `json:"click_count" bson:"click_count"`
	Status          string    `json:"status" bson:"status"`
}

// ItemPostSearchRequest アイテム検索のリクエスト
type ItemPostSearchRequest struct {
	Keyword       string `json:"keyword" bson:"keyword"`
	ItemType      string `json:"item_type" bson:"item_type"`
	Gender        string `json:"gender" bson:"gender"`
	Silhouette    string `json:"silhouette" bson:"silhouette"`
	Complex       string `json:"complex" bson:"complex"`
	UserSettingID string `json:"user_setting_id" bson:"user_setting_id"`
}

// OgpRequestByItemPostIdAndUrl アイテムの投稿ID・URLでのOGP取得リクエスト
type OgpRequestByItemPostIdAndUrl struct {
	PostId string `json:"post_id" bson:"post_id"`
	URL    string `json:"url" bson:"url"`
}

// OgpRequestByUrl URLでのOGP取得リクエスト
type OgpRequestByUrl struct {
	URL string `form:"url" bson:"url"`
}

// OgpResponse OGPの情報
type OgpResponse struct {
	PostID      string `json:"post_id" bson:"post_id"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	ImageURL    string `json:"image_url" bson:"image_url"`
}
