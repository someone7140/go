package handler

import (
	"net/http"

	"../config"
	"../util"
	"github.com/gin-gonic/gin"
)

// YahooMapJs YahooMapのJsを取得
func YahooMapJs(c *gin.Context) {
	yahooMapJs := util.SendGetHTTPRequest(
		"https://map.yahooapis.jp/js/V1/jsapi?appid=" + config.YahooAPIKey)
	if yahooMapJs == "" {
		c.Status(http.StatusNoContent)
	} else {
		c.JSON(http.StatusOK, yahooMapJs)
	}
}
