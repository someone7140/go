package coordinate

import (
	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	coordinateModel "coda-api/src/model/coordinate"
	coordinateRepository "coda-api/src/repository/coordinate"
	"coda-api/src/util"
	"mime/multipart"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/thoas/go-funk"
)

// AddCoordinatePost コーデ投稿の追加
func AddCoordinatePost(
	request coordinateModel.CoordinatePostRequest,
	loginInfo *authModel.AuthUserResponse,
	files []multipart.File,
	fileHeaders []*multipart.FileHeader,
) error {
	if request.Title == "" || request.Status == "" {
		return errorConstants.ErrBadRequest
	}

	shop, _ := coordinateRepository.GetShopInfoByShopSettingID(request.ShopSettingID)
	if shop == nil {
		return errorConstants.ErrBadRequest
	}

	uuid, err := util.GenerateUUID()
	if err != nil {
		return err
	}

	images := uploadImages(uuid, files, fileHeaders)
	return coordinateRepository.AddCoordinateRepositoryPost(uuid, request, images, loginInfo.ID, shop.ID)
}

// UpdateCoordinatePost コーデ投稿の更新
func UpdateCoordinatePost(
	request coordinateModel.CoordinatePostUpdateRequest,
	loginInfo *authModel.AuthUserResponse,
	files []multipart.File,
	fileHeaders []*multipart.FileHeader,
) error {
	if request.ID == "" || request.Title == "" || request.Status == "" {
		return errorConstants.ErrBadRequest
	}

	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	registeredPost, _ := coordinateRepository.GetCoordinatePostByID(request.ID, myUserID, true)
	if registeredPost == nil {
		return errorConstants.ErrBadRequest
	}
	shop, err := coordinateRepository.GetShopInfoByShopSettingID(request.ShopSettingID)
	if shop == nil {
		return errorConstants.ErrBadRequest
	}

	// ファイルの削除
	deleteImages := funk.Filter(registeredPost.Images, func(i coordinateModel.CoordinateImage) bool {
		return funk.Contains(request.DeleteImageKeys, i.Key)
	}).([]coordinateModel.CoordinateImage)
	for _, deleteImage := range deleteImages {
		fileName := util.GetFileNameFromURL(deleteImage.URL)
		err = util.DeleteFileFromGcs("coordinatePost/" + fileName)
		if err != nil {
			return err
		}
	}
	// ファイルの追加
	addImages := uploadImages(request.ID, files, fileHeaders)
	// 追加されたファイルと変更がないファイルを結合
	images :=
		append(addImages, funk.Filter(registeredPost.Images, func(i coordinateModel.CoordinateImage) bool {
			return !funk.Contains(request.DeleteImageKeys, i.Key)
		}).([]coordinateModel.CoordinateImage)...)
	sort.Slice(images, func(i, j int) bool { return images[i].Key < images[j].Key })

	// DB更新
	return coordinateRepository.UpdateCoordinateRepositoryPost(request, images, loginInfo.ID, shop.ID)
}

//  DeleteCoordinatePost コーデ投稿の削除
func DeleteCoordinatePost(postID string, loginInfo *authModel.AuthUserResponse) error {
	if postID == "" {
		return errorConstants.ErrBadRequest
	}
	coordinatePostTarget, err := coordinateRepository.GetCoordinatePostByID(postID, loginInfo.ID, true)
	if err != nil || coordinatePostTarget == nil {
		return errorConstants.ErrBadRequest
	}
	if loginInfo.UserType != "admin" && coordinatePostTarget.PostUserID != loginInfo.ID {
		return errorConstants.ErrBadRequest
	}
	// imageURLが登録されている場合は削除
	if len(coordinatePostTarget.Images) > 0 {
		for _, image := range coordinatePostTarget.Images {
			fileName := util.GetFileNameFromURL(image.URL)
			err = util.DeleteFileFromGcs("coordinatePost/" + fileName)
			if err != nil {
				return err
			}
		}
	}
	return coordinateRepository.DeleteCoordinatePost(postID, loginInfo)
}

// GetRecentCoordinatePosts 最新のコーデ投稿の取得
func GetRecentCoordinatePosts(loginInfo *authModel.AuthUserResponse) ([]coordinateModel.CoordinatePostInfo, error) {
	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	return coordinateRepository.GetRecentCoordinatePosts(100, false, myUserID)
}

// GetRecentCoordinatePostsForAdmin 最新のコーデ投稿の取得（管理者用）
func GetRecentCoordinatePostsForAdmin(loginInfo *authModel.AuthUserResponse) ([]coordinateModel.CoordinatePostInfo, error) {
	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	return coordinateRepository.GetRecentCoordinatePosts(200, true, myUserID)
}

// GetSearchCoordinatePosts コーデ投稿の検索
func GetSearchCoordinatePosts(
	searchRequest coordinateModel.CoordinatePostSearchRequest,
	loginInfo *authModel.AuthUserResponse,
) ([]coordinateModel.CoordinatePostInfo, error) {
	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	return coordinateRepository.GetSearchCoordinatePosts(100, false, myUserID, searchRequest)
}

// GetSearchCoordinatePostsForAdmin コーデ投稿の検索（管理者用）
func GetSearchCoordinatePostsForAdmin(
	searchRequest coordinateModel.CoordinatePostSearchRequest,
	loginInfo *authModel.AuthUserResponse,
) ([]coordinateModel.CoordinatePostInfo, error) {
	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	return coordinateRepository.GetSearchCoordinatePosts(200, true, myUserID, searchRequest)
}

// GetCoordinatePostByID ID指定でコーデ投稿の取得
func GetCoordinatePostByID(postID string, loginInfo *authModel.AuthUserResponse) (*coordinateModel.CoordinatePostInfo, error) {
	if postID == "" {
		return nil, errorConstants.ErrBadRequest
	}
	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	return coordinateRepository.GetCoordinatePostByID(postID, myUserID, false)
}

// GetFavoritedCoordinatePosts いいねしたコーデ投稿の一覧を取得
func GetFavoritedCoordinatePosts(limit int, userID string) ([]coordinateModel.CoordinatePostInfo, error) {
	return coordinateRepository.GetFavoritedCoordinatePosts(limit, userID)
}

// UpdateFavoritePostCoordinate コーデ投稿に対するいいねの更新
func UpdateFavoritePostCoordinate(coordinatePostID string, userID string) error {
	if coordinatePostID == "" || userID == "" {
		return errorConstants.ErrBadRequest
	}
	// 自ユーザの投稿いいね状況
	statusFavorite, err := coordinateRepository.GetStatusCoordinatePostFavorite(coordinatePostID, userID)
	if err != nil {
		return errorConstants.ErrInternalServer
	}
	if statusFavorite == nil || statusFavorite.UserID == userID {
		return errorConstants.ErrBadRequest
	}
	// いいねが付いているときは削除
	if statusFavorite.Status == "registered" {
		return coordinateRepository.DeleteFavoriteCoordinatePost(coordinatePostID, userID)
	} else {
		// いいねが付いていないときは追加
		return coordinateRepository.AddFavoriteCoordinatePost(coordinatePostID, userID)
	}
}

// UpdateImpressionPostCoordinate コーデ投稿に対するインプレッションの更新
func UpdateImpressionPostCoordinate(coordinatePostIDs []string, userID string) {
	if coordinatePostIDs != nil {
		coordinateRepository.AddImpressionCoordinatePosts(coordinatePostIDs, userID)
	}
}

// UpdateClickPostCoordinate コーデ投稿に対するクリックの更新
func UpdateClickPostCoordinate(coordinatePostID string, userID string) {
	if coordinatePostID != "" {
		coordinateRepository.AddClickCoordinatePost(coordinatePostID, userID)
	}
}

// UpdatePurchaseRequestCount コーデ投稿に対するクリックの更新
func UpdatePurchaseRequestCount(coordinatePostID string, userID string) {
	if coordinatePostID != "" {
		coordinateRepository.AddPurchaseRequestCount(coordinatePostID, userID)
	}
}

func uploadImages(
	uuid string,
	files []multipart.File,
	fileHeaders []*multipart.FileHeader) []coordinateModel.CoordinateImage {
	var images []coordinateModel.CoordinateImage
	for i, f := range files {
		fileHeader := fileHeaders[i]
		path := util.GetFilePathForGcs(
			uuid+"_"+fileHeader.Filename, nil, "coordinatePost",
		)
		imageURL, err := util.UploadNewFileToGcs(path, f)
		if err == nil {
			// ファイル名からkeyを取得
			key, err := strconv.Atoi(strings.TrimSuffix(fileHeader.Filename, filepath.Ext(fileHeader.Filename)))
			if err == nil {
				images = append(images, coordinateModel.CoordinateImage{
					Key: key,
					URL: imageURL,
				})
			}
		}
	}
	return images
}
