package coordinate

import (
	errorConstants "coda-api/src/constants"
	shopModel "coda-api/src/model/coordinate"
	coordinate "coda-api/src/service/coordinate"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AddShop ショップの追加
func AddShop(c *gin.Context) {
	var request shopModel.ShopInfo
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err = coordinate.AddShop(request)
		setShopResponse(c, nil, err)
	}

}

// UpdateShop ショップの更新
func UpdateShop(c *gin.Context) {
	var request shopModel.ShopInfo
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err = coordinate.UpdateShop(request)
		setShopResponse(c, nil, err)
	}

}

// DeleteShop ショップの削除
func DeleteShop(c *gin.Context) {
	var request shopModel.ShopDeleteRequest
	err := c.BindJSON(&request)
	if err != nil {
		c.Status(http.StatusBadRequest)
	} else {
		err = coordinate.DeleteShop(request)
		setShopResponse(c, nil, err)
	}

}

// GetShopByShopSettingId ショップ情報の取得
func GetShopByShopSettingId(c *gin.Context) {
	shopInfo, err := coordinate.GetShopByShopSettingId(c.Query("shop_setting_id"))
	setShopResponse(c, shopInfo, err)
}

// GetShopList ショップのリストを取得
func GetShopList(c *gin.Context) {
	shopList, err := coordinate.GetShopList()
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, shopList)
	}
}

// setShopResponse ショップAPIのレスポンス
func setShopResponse(c *gin.Context, shopInfo *shopModel.ShopInfo, err error) {
	if err == errorConstants.ErrForbidden {
		c.Status(http.StatusForbidden)
	} else if err == errorConstants.ErrBadRequest {
		c.Status(http.StatusBadRequest)
	} else if err != nil {
		c.Status(http.StatusInternalServerError)
	} else {
		c.JSON(http.StatusOK, shopInfo)
	}
}
