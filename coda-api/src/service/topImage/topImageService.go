package topImage

import (
	authModel "coda-api/src/model/auth"
	coordinateModel "coda-api/src/model/coordinate"
	itemPostModel "coda-api/src/model/item"
	topImageModel "coda-api/src/model/topImage"
	coordinateRepository "coda-api/src/repository/coordinate"
	itemPostRepository "coda-api/src/repository/itemPost"
	"sort"

	"github.com/thoas/go-funk"
)

//  GetTopRecentImages 最新順でトップ用の画像を取得
func GetTopRecentImages(loginUser *authModel.AuthUserResponse) (*topImageModel.TopImageResponse, error) {
	myUserID := ""
	if loginUser != nil {
		myUserID = loginUser.ID
	}

	// 画像登録があるアイテム投稿を取得
	itemPosts, err := itemPostRepository.GetRecentItemWithImagePosts(100, myUserID)
	if err != nil {
		return nil, err
	}
	// コーデ投稿を全部取得
	coordinatePosts, err := coordinateRepository.GetRecentCoordinateWithImagePosts(100, myUserID, false)
	if err != nil {
		return nil, err
	}

	// 画像の配列に追加
	var topImages []topImageModel.TopImage
	funk.ForEach(itemPosts, func(itemPost itemPostModel.ItemPostResponse) {
		topImages = append(topImages, topImageModel.TopImage{
			ID:                itemPost.ID,
			Category:          "item",
			Title:             itemPost.Title,
			ImageUrl:          itemPost.ImageURL,
			FavoritedFlg:      itemPost.FavoritedFlg,
			PostDate:          itemPost.PostDate,
			PostUserID:        itemPost.UserID,
			PostUserSettingID: itemPost.UserSettingID,
			PostUserName:      itemPost.UserName,
			ItemType:          itemPost.ItemType,
		})
	})
	funk.ForEach(coordinatePosts, func(coordinatePost coordinateModel.CoordinatePostInfo) {
		topImages = append(topImages, topImageModel.TopImage{
			ID:                 coordinatePost.ID,
			Category:           "coordinate",
			Title:              coordinatePost.Title,
			ImageUrl:           coordinatePost.Images[0].URL,
			FavoritedFlg:       coordinatePost.FavoritedFlg,
			PostDate:           coordinatePost.PostDate,
			PostUserID:         coordinatePost.PostUserID,
			ShopSettingID:      coordinatePost.ShopSettingID,
			ShopName:           coordinatePost.ShopName,
			CoordinateCategory: coordinatePost.Category,
			Price:              coordinatePost.Price,
			Sale:               coordinatePost.Sale,
		})
	})
	// ソート
	sort.Slice(topImages, func(i, j int) bool {
		return topImages[i].PostDate.Unix() > topImages[j].PostDate.Unix()
	})
	// 上位105件
	var recentImages []topImageModel.TopImage
	if len(topImages) > 105 {
		recentImages = topImages[0:105]
	} else {
		recentImages = topImages
	}

	// セール情報
	var saleImages []topImageModel.TopImage
	saleCoordinatePosts, err := coordinateRepository.GetRecentCoordinateWithImagePosts(100, myUserID, true)
	if err != nil {
		return nil, err
	}
	funk.ForEach(saleCoordinatePosts, func(coordinatePost coordinateModel.CoordinatePostInfo) {
		saleImages = append(saleImages, topImageModel.TopImage{
			ID:                 coordinatePost.ID,
			Category:           "coordinate",
			Title:              coordinatePost.Title,
			ImageUrl:           coordinatePost.Images[0].URL,
			FavoritedFlg:       coordinatePost.FavoritedFlg,
			PostDate:           coordinatePost.PostDate,
			PostUserID:         coordinatePost.PostUserID,
			ShopSettingID:      coordinatePost.ShopSettingID,
			ShopName:           coordinatePost.ShopName,
			CoordinateCategory: coordinatePost.Category,
			Price:              coordinatePost.Price,
			Sale:               coordinatePost.Sale,
		})
	})

	return &topImageModel.TopImageResponse{
		RecentPosts:   recentImages,
		SaleOnlyPosts: saleImages,
	}, nil

}
