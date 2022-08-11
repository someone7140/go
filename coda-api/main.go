package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"coda-api/src/handler/analytics"
	"coda-api/src/handler/auth"
	"coda-api/src/handler/common"
	"coda-api/src/handler/coordinate"
	"coda-api/src/handler/cron"
	"coda-api/src/handler/instagramAccount"
	"coda-api/src/handler/instagramBatch"
	"coda-api/src/handler/itemPost"
	"coda-api/src/handler/post"
	"coda-api/src/handler/topImage"
	user "coda-api/src/handler/user"
	"coda-api/src/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// envLoad 環境毎の設定を読み込み
func envLoad() {
	// 環境変数GO_ENV
	if os.Getenv("GO_ENV") == "" {
		os.Setenv("GO_ENV", "local")
	}

	err := godotenv.Load(fmt.Sprintf("./env/%s.env", os.Getenv("GO_ENV")))

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	envLoad()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		// 許可したいHTTPメソッドの一覧
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
		// 許可したいHTTPリクエストヘッダの一覧
		AllowHeaders: []string{
			"Access-Control-Allow-Headers",
			"Content-Type",
			"Content-Length",
			"Accept-Encoding",
			"X-CSRF-Token",
			"Authorization",
		},
		// 許可したいアクセス元の一覧
		AllowOrigins: []string{
			os.Getenv("VIEW_DOMAIN"),
		},
		MaxAge:           24 * time.Hour,
		AllowCredentials: true,
	}))
	// 共通系
	router.GET("/getMaster", common.GetMaster)
	// 認証系
	router.POST("/authByGoogle", auth.LoginByGoogle)
	router.POST("/authByFacebook", auth.LoginByFacebook)
	router.POST("/authByEmail", auth.LoginByEmail)
	router.POST("/authEmailRegister", auth.AuthEmailRegister)
	router.POST("/authEmailToken", auth.AuthEmailToken)
	router.POST("/registerPasswordReset", auth.RegisterPasswordReset)
	router.POST("/passwordResetUpdate", auth.PasswordResetUpdate)
	// ユーザ系
	router.GET("/checkRegisterByGoogleLogin", user.CheckRegisterByGoogleLogin)
	router.GET("/checkRegisterByFacebookLogin", user.CheckRegisterByFacebookLogin)
	router.POST("/registerByGoogleLogin", user.RegisterByGoogleLogin)
	router.POST("/registerByFacebookLogin", user.RegisterByFacebookLogin)
	router.POST("/registerByEmailAuth", user.RegisterByEmailAuth)
	// インスタ投稿系
	router.GET("/getRecommendPosts", post.GetRecommendPosts)
	// アイテム投稿系
	router.GET("/getItemsByUserID", itemPost.GetItemsByUserID)
	router.GET("/getItemByItemPostID", itemPost.GetItemByItemPostID)
	router.GET("/getRecommendItemPosts", itemPost.GetRecommendItemPosts)
	router.POST("/getItemPostsSearch", itemPost.GetItemPostsSearch)
	router.POST("/updateImpressionPostItem", itemPost.UpdateImpressionPostItem)
	router.POST("/updateClickPostItem", itemPost.UpdateClickPostItem)
	router.POST("/getOgpByURL", itemPost.GetOgpByURL)
	router.POST("/getOgpsByItempostIDAndURL", itemPost.GetOgpInfosByPostIDAndURL)
	// コーデ関連
	router.GET("/getCoordinateByPostID", coordinate.GetCoordinateByPostID)
	router.GET("/getRecentCoordinatePosts", coordinate.GetRecentCoordinatePosts)
	router.POST("/getCoordinatePostsSearch", coordinate.GetSearchCoordinatePosts)
	router.POST("/updateImpressionPostCoordinate", coordinate.UpdateImpressionPostCoordinate)
	router.POST("/updateClickPostCoordinate", coordinate.UpdateClickPostCoordinate)
	router.POST("/updatePurchaseRequestCount", coordinate.UpdatePurchaseRequestCount)
	// ショップ関連
	router.GET("/getShopList", coordinate.GetShopList)
	// トップ画像
	router.GET("/getTopRecentImages", topImage.GetTopRecentImages)
	// クーロン
	router.GET("/cronStartPostGather", cron.StartPostGatherHandler)
	router.GET("/cronStartAccessAnalysis", cron.StartAccessAnalysisHandler)
	router.POST("/executePostGatherBatch", cron.ExecutePostGatherBatch)
	router.POST("/executeAccessAnalysisBatch", cron.ExecuteAccessAnalysisBatch)
	// ログインユーザがアクセス可能
	loggedInUserGroup := router.Group("/loggedIn")
	loggedInUserGroup.Use(middleware.LoginCheckMiddleware())
	{
		// インスタ投稿系
		loggedInUserGroup.POST("/updateFavoritePost", post.UpdateFavoritePost)
		loggedInUserGroup.GET("/getFavoritedPosts", post.GetFavoritedPosts)
		// 認証系
		loggedInUserGroup.POST("/authCheck", auth.AuthCheck)
		loggedInUserGroup.POST("/changePassword", auth.ChangePassword)
		// ユーザ系
		loggedInUserGroup.GET("/getMyUserInfoDetail", user.GetMyUserInfoDetail)
		loggedInUserGroup.POST("/updateUserInfo", user.UpdateUserInfo)
		loggedInUserGroup.POST("/deleteUser", user.DeleteUser)
		// アイテム投稿系
		loggedInUserGroup.POST("/addItemPost", itemPost.AddItemPost)
		loggedInUserGroup.POST("/updateItemPost", itemPost.UpdateItemPost)
		loggedInUserGroup.POST("/deleteItemPost", itemPost.DeleteItemPost)
		loggedInUserGroup.GET("/getFavoritedItemPosts", itemPost.GetFavoritedItemPosts)
		loggedInUserGroup.POST("/updateFavoriteItemPost", itemPost.UpdateFavoriteItemPost)
		// コーデ関連
		loggedInUserGroup.GET("/getFavoritedCoordinatePosts", coordinate.GetFavoritedCoordinatePosts)
		loggedInUserGroup.POST("/updateFavoriteCoordinatePost", coordinate.UpdateFavoriteCoordinatePost)
	}
	// adminのみアクセス可能
	adminUserGroup := router.Group("/admin")
	adminUserGroup.Use(middleware.LoginCheckAdminMiddleware())
	{
		// 投稿関連
		adminUserGroup.GET("/getNotsetAllPosts", post.GetNotsetAllPosts)
		adminUserGroup.POST("/setStatusPosts", post.SetStatusPosts)
		// インスタアカウント関連
		adminUserGroup.GET("/getInstagramAccountInfo", instagramAccount.GetInstagramAccountInfo)
		adminUserGroup.GET("/getInstagramAccountWithPosts", instagramAccount.GetInstagramAccountWithPosts)
		adminUserGroup.POST("/addInstagramAccount", instagramAccount.AddInstagramAccount)
		adminUserGroup.POST("/editInstagramAccount", instagramAccount.EditInstagramAccount)
		// コーデ関連
		adminUserGroup.POST("/addCoordinatePost", coordinate.AddCoordinatePost)
		adminUserGroup.POST("/updateCoordinatePost", coordinate.UpdateCoordinatePost)
		adminUserGroup.POST("/deleteCoordinatePost", coordinate.DeleteCoordinatePost)
		adminUserGroup.GET("/getRecentCoordinatePostsForAdmin", coordinate.GetRecentCoordinatePostsForAdmin)
		adminUserGroup.POST("/getCoordinatePostsSearchForAdmin", coordinate.GetSearchCoordinatePostsForAdmin)
		// ショップ関連
		adminUserGroup.POST("/addShop", coordinate.AddShop)
		adminUserGroup.POST("/updateShop", coordinate.UpdateShop)
		adminUserGroup.POST("/deleteShop", coordinate.DeleteShop)
		adminUserGroup.GET("/getShopByShopSettingId", coordinate.GetShopByShopSettingId)
		// 分析関連
		adminUserGroup.GET("/getAccessAnalysis", analytics.GetAccessAnalysisData)
		adminUserGroup.GET("/getFavoriteAnalysis", analytics.GetFavoriteAnalysisData)
		adminUserGroup.GET("/getCoordinatePostAnalysis", analytics.GetCoordinatePostAccessAnalysisData)
		// バッチ関連
		adminUserGroup.GET("/gatherInstagranPost", instagramBatch.GatherInstagramPost)
		adminUserGroup.GET("/executeAccessAnalysis", analytics.ExecuteAccessAnalysis)
	}
	router.Run()
}
