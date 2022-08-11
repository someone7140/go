package postModel

import "time"

// PostResponseForAdmin 投稿の管理用Response用struct
type PostResponseForAdmin struct {
	ID                     string    `json:"_id" bson:"_id"`
	Status                 string    `json:"status" bson:"status"`
	PostInstagramAccountId string    `json:"post_instagram_account_id" bson:"post_instagram_account_id"`
	PostInstagramUserName  string    `json:"instagram_user_name" bson:"instagram_user_name"`
	Category               string    `json:"category" bson:"category"`
	PostGenre              string    `json:"post_genre" bson:"post_genre"`
	ContentURL             string    `json:"content_url" bson:"content_url"`
	PostDate               time.Time `json:"post_date" bson:"post_date"`
}

// PostWithAccountResponseForAdmin アカウントに紐づく投稿の管理用Response用struct
type PostWithAccountResponseForAdmin struct {
	ID         string    `json:"_id" bson:"_id"`
	Status     string    `json:"status" bson:"status"`
	ContentURL string    `json:"content_url" bson:"content_url"`
	PostDate   time.Time `json:"post_date" bson:"post_date"`
	PostGenre  string    `json:"post_genre" bson:"post_genre"`
}

// PostInfoWithCategory カテゴリー付きの投稿情報
type PostInfoWithCategory struct {
	ID            string    `json:"_id" bson:"_id"`
	ContentURL    string    `json:"content_url" bson:"content_url"`
	PostDate      time.Time `json:"post_date" bson:"post_date"`
	Category      string    `json:"category" bson:"category"`
	FavoriteCount int       `json:"favorite_count" bson:"favorite_count"`
	FavoritedFlg  bool      `json:"favorited_flg" bson:"favorited_flg"`
}

// PostResponseForRecommend ユーザ提示用のレスポンス
type PostResponseForRecommend struct {
	ID            string    `json:"_id" bson:"_id"`
	ContentURL    string    `json:"content_url" bson:"content_url"`
	Genre         string    `json:"genre" bson:"genre"`
	PostDate      time.Time `json:"post_date" bson:"post_date"`
	FavoriteCount int       `json:"favorite_count" bson:"favorite_count"`
	FavoritedFlg  bool      `json:"favorited_flg" bson:"favorited_flg"`
}

// PostStatusUpdateRequest 投稿のステータス更新リクエスト
type PostStatusUpdateRequest struct {
	PostID string `json:"post_id" bson:"post_id"`
	Status string `json:"status" bson:"status"`
	Genre  string `json:"genre" bson:"genre"`
}

// UpdateFavoriteRequest いいね更新のリクエスト
type UpdateFavoriteRequest struct {
	PostID string `json:"post_id" bson:"post_id"`
}

// StatusUserFavorite いいねしたかしてないかのステータス
type StatusUserFavorite struct {
	PostID string `json:"post_id" bson:"post_id"`
	Status string `json:"status" bson:"status"`
}
