package coordinate

import (
	errorConstants "coda-api/src/constants"
	coordinateModel "coda-api/src/model/coordinate"
	coordinateRepository "coda-api/src/repository/coordinate"
	"coda-api/src/util"
)

// AddShop ショップの追加
func AddShop(shopInfoRequest coordinateModel.ShopInfo) error {
	if shopInfoRequest.ShopSettingID == "" || shopInfoRequest.Name == "" {
		return errorConstants.ErrBadRequest
	}
	// IDの重複チェック
	shopRegsitered, err := coordinateRepository.GetShopInfoByShopSettingID(shopInfoRequest.ShopSettingID)
	if err != nil || shopRegsitered != nil {
		return errorConstants.ErrBadRequest
	}

	uuid, err := util.GenerateUUID()
	if err != nil {
		return err
	}
	return coordinateRepository.AddShop(uuid, shopInfoRequest)
}

// UpdateShop ショップの更新
func UpdateShop(shopInfoRequest coordinateModel.ShopInfo) error {
	if shopInfoRequest.ID == "" || shopInfoRequest.ShopSettingID == "" || shopInfoRequest.Name == "" {
		return errorConstants.ErrBadRequest
	}
	// IDチェック
	shopRegsitered, err := coordinateRepository.GetShopInfoByShopSettingID(shopInfoRequest.ShopSettingID)
	if err != nil || (shopRegsitered != nil && shopRegsitered.ID != shopInfoRequest.ID) {
		return errorConstants.ErrBadRequest
	}

	return coordinateRepository.UpdateShop(shopInfoRequest)
}

// GetShopByShopSettingId ショップ情報の取得
func GetShopByShopSettingId(shopSettingId string) (*coordinateModel.ShopInfo, error) {
	if shopSettingId == "" {
		return nil, errorConstants.ErrBadRequest
	}
	return coordinateRepository.GetShopInfoByShopSettingID(shopSettingId)

}

// GetShopList ショップ情報のリスト
func GetShopList() ([]coordinateModel.ShopInfo, error) {
	return coordinateRepository.GetShopList()
}

// DeleteShop ショップ情報の削除
func DeleteShop(shopDeleteRequest coordinateModel.ShopDeleteRequest) error {
	if shopDeleteRequest.ShopSettingID == "" {
		return errorConstants.ErrBadRequest
	}
	// ショップの取得
	shopInfo, err := coordinateRepository.GetShopInfoByShopSettingID(shopDeleteRequest.ShopSettingID)
	if err != nil {
		return err
	}
	if shopInfo == nil {
		return nil
	}
	// コーデ投稿の取得
	coordinatePosts, err := coordinateRepository.GetSearchCoordinatePosts(
		-1,
		true,
		"",
		coordinateModel.CoordinatePostSearchRequest{
			ShopSettingId: shopDeleteRequest.ShopSettingID,
		},
	)
	if err != nil {
		return err
	}
	if len(coordinatePosts) > 0 {
		// 画像の削除
		var deleteImageUrls []string
		for _, post := range coordinatePosts {
			if len(post.Images) > 0 {
				for _, image := range post.Images {
					deleteImageUrls = append(deleteImageUrls, image.URL)
				}
			}
		}
		for _, imageUrl := range deleteImageUrls {
			fileName := util.GetFileNameFromURL(imageUrl)
			err = util.DeleteFileFromGcs("coordinatePost/" + fileName)
			if err != nil {
				return err
			}
		}
		// 投稿の削除
		coordinateRepository.DeleteAllCoordinatePostsByShop(shopInfo.ID)
	}
	// ショップの削除
	return coordinateRepository.DeleteShop(shopInfo.ID)
}
