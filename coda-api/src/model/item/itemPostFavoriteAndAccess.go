package itemPostModel

// UpdateFavoriteItemPostRequest アイテムのいいね更新のリクエスト
type UpdateFavoriteItemPostRequest struct {
	ItemPostID string `json:"item_post_id" bson:"item_post_id"`
}

// UpdateImpressionItemPostRequest アイテムのインプレッション更新のリクエスト
type UpdateImpressionItemPostRequest struct {
	ItemPostIDs []string `json:"item_post_ids" bson:"item_post_ids"`
}

// UpdateClickItemPostRequest アイテムのインプレッション更新のリクエスト
type UpdateClickItemPostRequest struct {
	ItemPostID string `json:"item_post_id" bson:"item_post_id"`
}

// StatusUserFavoriteItemPost いいねしたかしてないかのステータス
type StatusUserFavoriteItemPost struct {
	PostID string `json:"post_id" bson:"post_id"`
	UserID string `json:"user_id" bson:"user_id"`
	Status string `json:"status" bson:"status"`
}
