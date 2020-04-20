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
		delivelyStoreList := service.ResponseJSONConvert(
			util.SendGetHTTPRequest(baseURL+"&deliverly=1"),
			"deliverly")
		// テイクアウト可の一覧
		takeoutStoreList := service.ResponseJSONConvert(
			util.SendGetHTTPRequest(baseURL+"&takeout=1"),
			"takeout")
		c.JSON(http.StatusOK, service.MergeStoreList(delivelyStoreList, takeoutStoreList))
	}
}
