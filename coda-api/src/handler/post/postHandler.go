package post

import (
	"coda-api/src/service/auth"
	"coda-api/src/service/post"
	"net/http"
	"strconv"

	errorConstants "coda-api/src/constants"
	postModel "coda-api/src/model/post"

	"github.com/gin-gonic/gin"
)

// GetNotsetAllPosts ステータス未設定の投稿を取得
func GetNotsetAllPosts(c *gin.Context) {
	limitStr := c.Query("limit")
	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.Status(http.StatusBadRequest)
	}
	results, err := post.GetNotsetAllPosts(limit)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

// SetStatusPosts 投稿へのステータス一括設定
func SetStatusPosts(c *gin.Context) {
	var request []postModel.PostStatusUpdateRequest
	err := c.BindJSON(&request)
	if err != nil || len(request) < 0 {
		c.Status(http.StatusBadRequest)
	}
	err = post.SetStatusPosts(request)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetRecommendPosts ユーザに提示する投稿一覧の取得
func GetRecommendPosts(c *gin.Context) {
	// 投稿の取得件数
	limit := 50
	var results []postModel.PostResponseForRecommend
	var err error
	// ログイン済みユーザかどうかで判定
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	// ジャンルのパラメータ取得
	genre := c.Query("genre")
	if loginInfo == nil {
		results, err = post.GetRecenstPosts(limit, genre)
	} else {
		results, err = post.GetUserMatchingPosts(limit, loginInfo, genre)
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

// UpdateFavoritePost 投稿に対するいいねの更新
func UpdateFavoritePost(c *gin.Context) {
	var request postModel.UpdateFavoriteRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	err = post.UpdateFavoriteRequest(request.PostID, loginInfo.ID)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

// GetFavoritedPosts いいねした投稿一覧の取得
func GetFavoritedPosts(c *gin.Context) {
	// 投稿の取得件数
	limit := 50
	var results []postModel.PostResponseForRecommend
	loginInfo := auth.GetLoginInfoFromContext(c)
	results, err := post.GetFavoritedPosts(limit, loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}
