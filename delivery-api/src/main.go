package main

import (
	"./handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "index.html", gin.H{})
	})
	router.POST("/searchStore", handler.StoreSearch)
	router.GET("/yahooMapJs", handler.YahooMapJs)

	router.Run()
}
