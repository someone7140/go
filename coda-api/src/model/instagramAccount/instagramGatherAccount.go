package instagramAccountModel

import (
	postModel "coda-api/src/model/post"
	"time"
)

// InstagramAccountAddRequest インスタグラムユーザの追加リクエスト
type InstagramAccountAddRequest struct {
	InstagramUserName string `json:"instagram_user_name" bson:"instagram_user_name"`
	Status            string `json:"status" bson:"status"`
	Gender            string `json:"gender" bson:"gender"`
	Silhouette        string `json:"silhouette" bson:"silhouette"`
	Height            string `json:"height" bson:"height"`
	Genre             string `json:"genre" bson:"genre"`
}

// InstagramAccountEditRequest インスタグラムユーザの編集リクエスト
type InstagramAccountEditRequest struct {
	ID         string `json:"id" bson:"id"`
	Status     string `json:"status" bson:"status"`
	Gender     string `json:"gender" bson:"gender"`
	Silhouette string `json:"silhouette" bson:"silhouette"`
	Height     string `json:"height" bson:"height"`
	Genre      string `json:"genre" bson:"genre"`
}

// InstagramAccountInfoResponse インスタグラムユーザのアカウント情報レスポンス
type InstagramAccountInfoResponse struct {
	Id                string `json:"_id" bson:"_id"`
	InstagramUserName string `json:"instagram_user_name" bson:"instagram_user_name"`
	Status            string `json:"status" bson:"status"`
	Category          string `json:"category" bson:"category"`
}

// InstagramAccountWithPosts インスタグラムユーザのアカウントと投稿情報レスポンス
type InstagramAccountWithPosts struct {
	Id                string                                      `json:"_id" bson:"_id"`
	InstagramUserName string                                      `json:"instagram_user_name" bson:"instagram_user_name"`
	Status            string                                      `json:"status" bson:"status"`
	Category          string                                      `json:"category" bson:"category"`
	GatherDate        time.Time                                   `json:"gather_date" bson:"gather_date"`
	Posts             []postModel.PostWithAccountResponseForAdmin `json:"posts" bson:"posts"`
}

// InstagramAccuntID インスタグラムユーザのIDとユーザネームを保持
type InstagramAccuntID struct {
	ID                string
	InstagramUserName string
}

// InstagramPost インスタグラムの投稿を保持
type InstagramPost struct {
	ID        string
	Permalink string
	PostDate  time.Time
}
