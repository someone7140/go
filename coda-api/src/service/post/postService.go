package post

import (
	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	postModel "coda-api/src/model/post"
	postRepository "coda-api/src/repository/post"
	"coda-api/src/util"
)

//  GetNotsetAllPosts ステータス未設定の投稿を取得
func GetNotsetAllPosts(limit int) ([]postModel.PostResponseForAdmin, error) {
	return postRepository.GetNotsetAllPosts(limit)
}

//  SetStatusPosts 投稿に対するステータス変更
func SetStatusPosts(request []postModel.PostStatusUpdateRequest) error {
	return postRepository.SetStatusPosts(request)
}

// 最新の投稿を取得
func GetRecenstPosts(limit int, genre string) ([]postModel.PostResponseForRecommend, error) {
	return postRepository.GetRecentPosts(limit, genre)
}

// ユーザにマッチした投稿を取得
func GetUserMatchingPosts(limit int, loginUser *authModel.AuthUserResponse, genre string) ([]postModel.PostResponseForRecommend, error) {
	return postRepository.GetCategoryMatchedPosts(limit, util.GetMatchCategories(loginUser.Categories), loginUser.ID, genre)
}

// いいねした投稿を取得
func GetFavoritedPosts(limit int, loginUser *authModel.AuthUserResponse) ([]postModel.PostResponseForRecommend, error) {
	return postRepository.GetFavoritedPosts(limit, loginUser.ID)
}

// 投稿に対するいいねの更新
func UpdateFavoriteRequest(postID string, userID string) error {
	if postID == "" || userID == "" {
		return errorConstants.ErrBadRequest
	}
	// 自ユーザの投稿いいね状況
	statusFavorite, err := postRepository.GetStatusFavorite(postID, userID)
	if err != nil {
		return errorConstants.ErrInternalServer
	}
	if statusFavorite == nil {
		return errorConstants.ErrBadRequest
	}
	// いいねが付いているときは削除
	if statusFavorite.Status == "registered" {
		return postRepository.DeleteFavorite(postID, userID)
	} else {
		// いいねが付いていないときは追加
		return postRepository.AddFavorite(postID, userID)
	}
}
