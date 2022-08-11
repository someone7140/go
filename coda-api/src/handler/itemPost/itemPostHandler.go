package itemPost

import (
	errorConstants "coda-api/src/constants"
	itemPostModel "coda-api/src/model/item"
	"coda-api/src/service/auth"
	"net/http"

	itemPost "coda-api/src/service/itemPost"

	"github.com/gin-gonic/gin"
)

// AddItemPost アイテム投稿の追加
func AddItemPost(c *gin.Context) {
	var request itemPostModel.ItemPostRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	file, fileHeader, err := c.Request.FormFile("item_image")
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		err = itemPost.AddItemPost(request, loginInfo, file, fileHeader)
		file.Close()
		if err != nil {
			c.Status(http.StatusInternalServerError)
		} else {
			c.Status(http.StatusOK)
		}
	}
}

// UpdateItemPost アイテム投稿の編集
func UpdateItemPost(c *gin.Context) {
	var request itemPostModel.ItemPostUpdateRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	file, fileHeader, err := c.Request.FormFile("item_image")
	err = itemPost.UpdateItemPost(request, loginInfo, file, fileHeader)
	file.Close()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// DeleteItemPost アイテム投稿の削除
func DeleteItemPost(c *gin.Context) {
	var request itemPostModel.ItemDeleteRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	err = itemPost.DeleteItemPost(request.PostID, loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetItemsByUserID ユーザ指定のアイテム投稿取得
func GetItemsByUserID(c *gin.Context) {
	userSettingID := c.Query("user_setting_id")
	if userSettingID != "" {
		c.Status(http.StatusBadRequest)
	}
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	results, err := itemPost.GetItemsByUserID(userSettingID, loginInfo, 200)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

// GetItemByItemPostID アイテム投稿ID指定の取得
func GetItemByItemPostID(c *gin.Context) {
	itemPostID := c.Query("item_post_id")
	if itemPostID != "" {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)

	result, err := itemPost.GetItemByPostID(itemPostID, loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else if result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.Status(http.StatusNotFound)
	}
}

// GetRecommendItemPosts ユーザに提示するアイテム投稿一覧の取得
func GetRecommendItemPosts(c *gin.Context) {
	// 投稿の取得件数
	limit := 50
	var results []itemPostModel.ItemPostResponse
	var err error
	// ログイン済みユーザかどうかで判定
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	if loginInfo == nil {
		results, err = itemPost.GetRecenstItemPosts(limit)
	} else {
		results, err = itemPost.GetUserMatchingItemPosts(limit, loginInfo)
	}
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

// GetFavoritedItemPosts いいねしたアイテム投稿の一覧を取得
func GetFavoritedItemPosts(c *gin.Context) {
	// 投稿の取得件数
	limit := 200
	var err error
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	results, err := itemPost.GetFavoritedItemPosts(limit, loginInfo.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, results)
	}
}

// GetItemPostsSearch アイテム投稿の検索
func GetItemPostsSearch(c *gin.Context) {
	var request itemPostModel.ItemPostSearchRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
		results, err := itemPost.GetSearchItemPosts(50, request, loginInfo)
		if err != nil {
			c.Status(http.StatusInternalServerError)
		} else {
			c.JSON(http.StatusOK, results)
		}
	}
}

// UpdateFavoriteItemPost アイテム投稿に対するいいねの更新
func UpdateFavoriteItemPost(c *gin.Context) {
	var request itemPostModel.UpdateFavoriteItemPostRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		loginInfo := auth.GetLoginInfoFromContext(c)
		err = itemPost.UpdateFavoritePostItem(request.ItemPostID, loginInfo.ID)
		if err == errorConstants.ErrBadRequest {
			c.Status(http.StatusBadRequest)
		} else {
			c.Status(http.StatusInternalServerError)
		}
		c.Status(http.StatusOK)
	}

}

// UpdateImpressionPostItem アイテム投稿に対するインプレッションの更新
func UpdateImpressionPostItem(c *gin.Context) {
	var request itemPostModel.UpdateImpressionItemPostRequest
	err := c.BindJSON(&request)
	if err == nil {
		loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
		userId := ""
		if loginInfo != nil {
			userId = loginInfo.ID
		}
		itemPost.UpdateImpressionPostItem(request.ItemPostIDs, userId)
	}
	c.Status(http.StatusOK)
}

// UpdateClickPostItem アイテム投稿に対するクリックの更新
func UpdateClickPostItem(c *gin.Context) {
	var request itemPostModel.UpdateClickItemPostRequest
	err := c.BindJSON(&request)
	if err == nil {
		loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
		userId := ""
		if loginInfo != nil {
			userId = loginInfo.ID
		}
		itemPost.UpdateClickPostItem(request.ItemPostID, userId)
	}
	c.Status(http.StatusOK)
}

// GetOgpByURL URLに対してOGPの情報を返す
func GetOgpByURL(c *gin.Context) {
	var request itemPostModel.OgpRequestByUrl
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	result, err := itemPost.GetOgpInfoByURL(request.URL, "")
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, result)
}

// GetOgpInfosByPostIDAndURL アイテム投稿IDとURLに対してOGPの情報を返す
func GetOgpInfosByPostIDAndURL(c *gin.Context) {
	var requests []itemPostModel.OgpRequestByItemPostIdAndUrl
	err := c.BindJSON(&requests)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	result := itemPost.GetOgpInfosByPostIDAndURL(requests)
	c.JSON(http.StatusOK, result)
}
