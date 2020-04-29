package handler

import (
	"net/http"

	"../model"
	"../service"
	"../util"
	"github.com/gin-gonic/gin"
)

// StoreSearch ぐるなびに店検索のAPIを投げる関数
func StoreSearch(c *gin.Context) {
	var request model.StoreInfoRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	}
	baseURL := service.GetStoreSearchURL(request)
	if baseURL == "" {
		c.Status(http.StatusNoContent)
	} else {
		// デリバリー可の一覧
		delivelyStoreList := service.ResponseStoreJSONConvert(
			util.SendGetHTTPRequest(baseURL+"&deliverly=1"),
			"deliverly")
		// テイクアウト可の一覧
		takeoutStoreList := service.ResponseStoreJSONConvert(
			util.SendGetHTTPRequest(baseURL+"&takeout=1"),
			"takeout")
		c.JSON(http.StatusOK, service.MergeStoreList(delivelyStoreList, takeoutStoreList))
	}
}

// GetCategoryL ぐるなびに大カテゴリーのAPIを投げる関数
func GetCategoryL(c *gin.Context) {
	url := service.GetCategoryLURL()
	if url == "" {
		c.Status(http.StatusNoContent)
	} else {
		// カテゴリーの一覧
		c.JSON(http.StatusOK, service.ResponseCategoryLJSONConvert(util.SendGetHTTPRequest(url)))
	}
}
