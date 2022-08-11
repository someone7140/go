package instagramAccount

import (
	"net/http"
	"strconv"

	errorConstants "coda-api/src/constants"
	instagramAccountModel "coda-api/src/model/instagramAccount"
	"coda-api/src/service/instagramAccount"

	"github.com/gin-gonic/gin"
)

// AddInstagramAccount インスタアカウントの追加
func AddInstagramAccount(c *gin.Context) {
	var request instagramAccountModel.InstagramAccountAddRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	err = instagramAccount.AddInstagramAccount(
		request.InstagramUserName,
		request.Status,
		request.Gender,
		request.Silhouette,
		request.Height,
		request.Genre,
	)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// EditInstagramAccount インスタアカウントの編集
func EditInstagramAccount(c *gin.Context) {
	var request instagramAccountModel.InstagramAccountEditRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	err = instagramAccount.EditInstagramAccount(
		request.ID,
		request.Status,
		request.Gender,
		request.Silhouette,
		request.Height,
		request.Genre,
	)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetInstagramAccountInfo インスタアカウントの情報取得
func GetInstagramAccountInfo(c *gin.Context) {
	instagramUserName := c.Query("instagram_user_name")
	if instagramUserName == "" {
		c.Status(http.StatusBadRequest)
	}
	instagramAccountInfoResponse, err := instagramAccount.GetGatherInstagramAccountInfoByUserName(instagramUserName)
	if err == errorConstants.ErrNotFound {
		c.Status(http.StatusNotFound)
	} else if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, instagramAccountInfoResponse)
	}
}

// GetInstagramAccountWithPosts インスタアカウントを投稿付きで取得
func GetInstagramAccountWithPosts(c *gin.Context) {
	instagramUserName := c.Query("instagram_user_name")
	limit, err := strconv.Atoi(c.Query("limit"))
	postStatus := c.Query("post_status")
	if err != nil || instagramUserName == "" {
		c.Status(http.StatusBadRequest)
	}
	instagramAccountInfoResponse, err := instagramAccount.GetGatherInstagramAccountWithPostsByUserName(
		instagramUserName, limit, postStatus,
	)
	if err == errorConstants.ErrNotFound {
		c.Status(http.StatusNotFound)
	} else if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, instagramAccountInfoResponse)
	}
}
