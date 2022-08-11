package itemPostModel

// ItemPostRequest アイテム投稿登録のリクエスト
type ItemPostRequest struct {
	Title      string `form:"title" bson:"title"`
	Detail     string `form:"detail" bson:"detail"`
	ItemType   string `form:"item_type" bson:"item_type"`
	URL        string `form:"url" bson:"url"`
	Status     string `form:"status" bson:"status"`
	Gender     string `form:"gender" bson:"gender"`
	Silhouette string `form:"silhouette" bson:"silhouette"`
	Complex    string `form:"complex" bson:"complex"`
}

// ItemPostUpdateRequest アイテム投稿更新のリクエスト
type ItemPostUpdateRequest struct {
	ID         string `form:"_id" bson:"_id"`
	Title      string `form:"title" bson:"title"`
	Detail     string `form:"detail" bson:"detail"`
	ItemType   string `form:"item_type" bson:"item_type"`
	URL        string `form:"url" bson:"url"`
	Status     string `form:"status" bson:"status"`
	Gender     string `form:"gender" bson:"gender"`
	Silhouette string `form:"silhouette" bson:"silhouette"`
	Complex    string `form:"complex" bson:"complex"`
}

// ItemDeleteRequest アイテム削除のリクエスト
type ItemDeleteRequest struct {
	PostID string `json:"post_id" bson:"post_id"`
}
