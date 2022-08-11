package coordinate

import (
	errorConstants "coda-api/src/constants"
	"coda-api/src/service/auth"
	"coda-api/src/util"
	"net/http"

	coordinateModel "coda-api/src/model/coordinate"
	coordinate "coda-api/src/service/coordinate"

	"github.com/gin-gonic/gin"
)

// AddCoordinatePost コーデ投稿の追加
func AddCoordinatePost(c *gin.Context) {
	var request coordinateModel.CoordinatePostRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	// 複数ファイルの取得
	files, fileHeaders, err := util.GetMultipartFiles(c, "coordinate_images")
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	err = coordinate.AddCoordinatePost(request, loginInfo, files, fileHeaders)
	util.CloseMultipartFiles(files)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// UpdateCoordinatePost コーデ投稿の更新
func UpdateCoordinatePost(c *gin.Context) {
	var request coordinateModel.CoordinatePostUpdateRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	// 複数ファイルの取得
	files, fileHeaders, err := util.GetMultipartFiles(c, "coordinate_images")
	if err != nil {
		c.Status(http.StatusInternalServerError)
	}
	err = coordinate.UpdateCoordinatePost(request, loginInfo, files, fileHeaders)
	util.CloseMultipartFiles(files)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// DeleteCoordinatePost コーデ投稿の削除
func DeleteCoordinatePost(c *gin.Context) {
	var request coordinateModel.CoordinateDeleteRequest
	err := c.Bind(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	err = coordinate.DeleteCoordinatePost(request.CoordinateID, loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.Status(http.StatusOK)
	}
}

// GetCoordinateByPostID コーデ投稿ID指定の取得
func GetCoordinateByPostID(c *gin.Context) {
	postID := c.Query("post_id")
	if postID != "" {
		c.Status(http.StatusBadRequest)
	}
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)

	result, err := coordinate.GetCoordinatePostByID(postID, loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else if result != nil {
		c.JSON(http.StatusOK, result)
	} else {
		c.Status(http.StatusNotFound)
	}
}

// GetRecentCoordinatePosts 最新のコーデ投稿の取得
func GetRecentCoordinatePosts(c *gin.Context) {
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	result, err := coordinate.GetRecentCoordinatePosts(loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetRecentCoordinatePostsForAdmin 最新のコーデ投稿の取得（管理者用）
func GetRecentCoordinatePostsForAdmin(c *gin.Context) {
	loginInfo := auth.GetLoginInfoFromContext(c)
	result, err := coordinate.GetRecentCoordinatePostsForAdmin(loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetSearchCoordinatePosts コーデ投稿の検索
func GetSearchCoordinatePosts(c *gin.Context) {
	var request coordinateModel.CoordinatePostSearchRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
	result, err := coordinate.GetSearchCoordinatePosts(request, loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetSearchCoordinatePostsForAdmin コーデ投稿の検索（管理者用）
func GetSearchCoordinatePostsForAdmin(c *gin.Context) {
	var request coordinateModel.CoordinatePostSearchRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	result, err := coordinate.GetSearchCoordinatePostsForAdmin(request, loginInfo)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// GetFavoritedCoordinatePosts いいねされたコーデ投稿の取得
func GetFavoritedCoordinatePosts(c *gin.Context) {
	loginInfo := auth.GetLoginInfoFromContext(c)
	result, err := coordinate.GetFavoritedCoordinatePosts(200, loginInfo.ID)
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, result)
	}
}

// UpdateFavoriteCoordinatePost コーデ投稿に対するいいねの更新
func UpdateFavoriteCoordinatePost(c *gin.Context) {
	var request coordinateModel.UpdateFavoriteCoordinatePostRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	loginInfo := auth.GetLoginInfoFromContext(c)
	err = coordinate.UpdateFavoritePostCoordinate(request.PostID, loginInfo.ID)
	if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else {
		c.Status(http.StatusInternalServerError)
	}
	c.Status(http.StatusOK)
}

// UpdateImpressionPostCoordinate コーデ投稿に対するインプレッションの更新
func UpdateImpressionPostCoordinate(c *gin.Context) {
	var request coordinateModel.UpdateImpressionCoordinatePostRequest
	err := c.BindJSON(&request)
	if err == nil {
		loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
		userId := ""
		if loginInfo != nil {
			userId = loginInfo.ID
		}
		coordinate.UpdateImpressionPostCoordinate(request.PostIDs, userId)
	}
	c.Status(http.StatusOK)
}

// UpdateClickPostCoordinate アイテム投稿に対するクリックの更新
func UpdateClickPostCoordinate(c *gin.Context) {
	var request coordinateModel.UpdateClickItemPostRequest
	err := c.BindJSON(&request)
	if err == nil {
		loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
		userId := ""
		if loginInfo != nil {
			userId = loginInfo.ID
		}
		coordinate.UpdateClickPostCoordinate(request.PostID, userId)
	}
	c.Status(http.StatusOK)
}

// UpdatePurchaseRequestCount アイテム投稿に対する購入申請の更新
func UpdatePurchaseRequestCount(c *gin.Context) {
	var request coordinateModel.UpdatePurchaseRequestCountCoordinatePost
	err := c.BindJSON(&request)
	if err == nil {
		loginInfo, _ := auth.GetLoginInfoInfoFromJwtToken(c)
		userId := ""
		if loginInfo != nil {
			userId = loginInfo.ID
		}
		coordinate.UpdatePurchaseRequestCount(request.PostID, userId)
	}
	c.Status(http.StatusOK)
}
