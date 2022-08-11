package coordinateModel

// UpdateFavoriteCoordinatePostRequest コーデ投稿のいいね更新のリクエスト
type UpdateFavoriteCoordinatePostRequest struct {
	PostID string `json:"post_id" bson:"post_id"`
}

// UpdateImpressionCoordinatePostRequest コーデのインプレッション更新のリクエスト
type UpdateImpressionCoordinatePostRequest struct {
	PostIDs []string `json:"post_ids" bson:"post_ids"`
}

// UpdateClickItemPostRequest アイテムのクリック更新のリクエスト
type UpdateClickItemPostRequest struct {
	PostID string `json:"post_id" bson:"post_id"`
}

// UpdatePurchaseRequestCountCoordinatePost コーデの購入申請のリクエスト
type UpdatePurchaseRequestCountCoordinatePost struct {
	PostID string `json:"post_id" bson:"post_id"`
}
