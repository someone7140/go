package itemPost

import (
	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	itemPostModel "coda-api/src/model/item"
	itemPostRepository "coda-api/src/repository/itemPost"
	"coda-api/src/service/user"
	"coda-api/src/util"
	"mime/multipart"
	"strings"

	"github.com/dyatlov/go-opengraph/opengraph"
)

//  AddItemPost アイテム投稿を追加
func AddItemPost(
	itemPostRequest itemPostModel.ItemPostRequest,
	loginInfo *authModel.AuthUserResponse,
	file multipart.File,
	fileHeader *multipart.FileHeader) error {
	if itemPostRequest.Title == "" || itemPostRequest.ItemType == "" || itemPostRequest.Status == "" {
		return errorConstants.ErrBadRequest
	}
	uuid, err := util.GenerateUUID()
	if err != nil {
		return err
	}
	imageURL := ""
	// ファイルがある場合はアップロード処理
	if file != nil {
		filePath := util.GetFilePathForGcs(
			uuid, fileHeader, "itemPost",
		)
		imageURL, err = util.UploadNewFileToGcs(filePath, file)
		if err != nil {
			return err
		}
	}
	return itemPostRepository.AddItemPost(uuid, itemPostRequest, loginInfo.ID, imageURL)
}

//  UpdateItemPost アイテム投稿の更新
func UpdateItemPost(
	itemPostRequest itemPostModel.ItemPostUpdateRequest,
	loginInfo *authModel.AuthUserResponse,
	file multipart.File,
	fileHeader *multipart.FileHeader) error {
	if itemPostRequest.ID == "" || itemPostRequest.Title == "" || itemPostRequest.ItemType == "" || itemPostRequest.Status == "" {
		return errorConstants.ErrBadRequest
	}
	itempPostTarget, err := itemPostRepository.GetItemByPostID(itemPostRequest.ID, loginInfo.ID, true)
	if err != nil || itempPostTarget == nil {
		return errorConstants.ErrBadRequest
	}
	if loginInfo.UserType != "admin" && itempPostTarget.UserID != loginInfo.ID {
		return errorConstants.ErrBadRequest
	}
	imageURL := itempPostTarget.ImageURL
	// ファイルがある場合はアップロード処理
	if file != nil {
		// iconURLが登録されている場合は削除
		if imageURL != "" {
			fileName := util.GetFileNameFromURL(imageURL)
			err = util.DeleteFileFromGcs("itemPost/" + fileName)
			if err != nil {
				return err
			}
		}
		filePath := util.GetFilePathForGcs(
			itemPostRequest.ID, fileHeader, "itemPost",
		)
		imageURL, err = util.UploadNewFileToGcs(filePath, file)
		if err != nil {
			return err
		}
	}
	return itemPostRepository.UpdateItemPost(itemPostRequest, loginInfo, imageURL)
}

//  DeleteItemPost アイテム投稿の削除
func DeleteItemPost(postID string, loginInfo *authModel.AuthUserResponse) error {
	if postID == "" {
		return errorConstants.ErrBadRequest
	}
	itempPostTarget, err := itemPostRepository.GetItemByPostID(postID, loginInfo.ID, true)
	if err != nil || itempPostTarget == nil {
		return errorConstants.ErrBadRequest
	}
	if loginInfo.UserType != "admin" && itempPostTarget.UserID != loginInfo.ID {
		return errorConstants.ErrBadRequest
	}
	// imageURLが登録されている場合は削除
	if itempPostTarget.ImageURL != "" {
		fileName := util.GetFileNameFromURL(itempPostTarget.ImageURL)
		err = util.DeleteFileFromGcs("itemPost/" + fileName)
		if err != nil {
			return err
		}
	}
	return itemPostRepository.DeleteItemPost(postID, loginInfo)
}

// GetRecenstItemPosts 最新のアイテム投稿を取得
func GetRecenstItemPosts(limit int) ([]itemPostModel.ItemPostResponse, error) {
	return itemPostRepository.GetRecentItemPosts(limit)
}

// GetUserMatchingItemPosts ユーザにマッチしたアイテム投稿を取得
func GetUserMatchingItemPosts(limit int, loginUser *authModel.AuthUserResponse) ([]itemPostModel.ItemPostResponse, error) {
	userDetail, err := user.GetUserInfoDetail(loginUser.ID)
	if err != nil || userDetail == nil {
		return nil, errorConstants.ErrInternalServer
	}
	return itemPostRepository.GetCategoryMatchedItemPosts(
		limit,
		util.GetMatchItemCategories(loginUser.Categories, userDetail.Complexes),
		loginUser.ID,
	)
}

// GetItemsByUserID ユーザID指定の投稿取得
func GetItemsByUserID(userSettingID string, loginInfo *authModel.AuthUserResponse, limit int) ([]itemPostModel.ItemPostResponse, error) {
	return itemPostRepository.GetItemsByUserID(userSettingID, loginInfo, 200)
}

// GetItemByPostID アイテム投稿ID指定の投稿取得
func GetItemByPostID(itemPostID string, loginInfo *authModel.AuthUserResponse) (*itemPostModel.ItemPostResponse, error) {
	myUserID := ""
	if loginInfo != nil {
		myUserID = loginInfo.ID
	}
	return itemPostRepository.GetItemByPostID(itemPostID, myUserID, false)
}

// GetSearchItemPosts アイテム投稿の検索
func GetSearchItemPosts(limit int, searchRequest itemPostModel.ItemPostSearchRequest, loginInfo *authModel.AuthUserResponse) ([]itemPostModel.ItemPostResponse, error) {
	return itemPostRepository.GetSearchItemPosts(limit, searchRequest, loginInfo)
}

// GetFavoritedItemPosts いいねしたアイテム投稿の一覧を取得
func GetFavoritedItemPosts(limit int, userID string) ([]itemPostModel.ItemPostResponse, error) {
	return itemPostRepository.GetFavoritedItemPosts(limit, userID)
}

// UpdateFavoritePostItem アイテム投稿に対するいいねの更新
func UpdateFavoritePostItem(itemPostID string, userID string) error {
	if itemPostID == "" || userID == "" {
		return errorConstants.ErrBadRequest
	}
	// 自ユーザの投稿いいね状況
	statusFavorite, err := itemPostRepository.GetStatusItemPostFavorite(itemPostID, userID)
	if err != nil {
		return errorConstants.ErrInternalServer
	}
	if statusFavorite == nil || statusFavorite.UserID == userID {
		return errorConstants.ErrBadRequest
	}
	// いいねが付いているときは削除
	if statusFavorite.Status == "registered" {
		return itemPostRepository.DeleteFavoriteItemPost(itemPostID, userID)
	} else {
		// いいねが付いていないときは追加
		return itemPostRepository.AddFavoriteItemPost(itemPostID, userID)
	}
}

// UpdateImpressionPostItem アイテム投稿に対するインプレッションの更新
func UpdateImpressionPostItem(itemPostIDs []string, userID string) {
	if itemPostIDs != nil {
		itemPostRepository.AddImpressionItemPosts(itemPostIDs, userID)
	}
}

// UpdateClickPostItem アイテム投稿に対するクリックの更新
func UpdateClickPostItem(itemPostID string, userID string) {
	if itemPostID != "" {
		itemPostRepository.AddClickItemPost(itemPostID, userID)
	}
}

// OGPの取得（URL単体）
func GetOgpInfoByURL(inputUrl string, postID string) (*itemPostModel.OgpResponse, error) {
	responseHtml, err := util.SendGetHTTPRequestWithUserAgent(inputUrl, util.GetOgpUserAgentBotFormURL(inputUrl))
	if err != nil || responseHtml == "" {
		return &itemPostModel.OgpResponse{
			PostID: postID}, err
	}
	og := opengraph.NewOpenGraph()
	err = og.ProcessHTML(strings.NewReader(responseHtml))
	if err != nil {
		return &itemPostModel.OgpResponse{
			PostID: postID}, err
	}
	imageURL := ""
	if len(og.Images) > 0 {
		imageURL = og.Images[0].URL
	}

	/*
		domain := ""
		if imageURL != "" {
			domain, _ = util.GetDomainFromUrl(imageURL)
		}
		// インスタグラムの場合はbase64で取得
		if strings.HasSuffix(domain, "instagram.com") {
			base64Data, _ := util.SendGetHTTPRequestForBase64Image(imageURL, "bot")
			if base64Data != "" {
				imageURL = "data:image/jpeg;base64," + base64Data
			}
		}
	*/

	return &itemPostModel.OgpResponse{
		PostID:      postID,
		Title:       og.Title,
		Description: og.Description,
		ImageURL:    imageURL,
	}, nil
}

// OGPの取得（アイテム投稿とURLのセット）
func GetOgpInfosByPostIDAndURL(requests []itemPostModel.OgpRequestByItemPostIdAndUrl) []itemPostModel.OgpResponse {
	var responses []itemPostModel.OgpResponse
	for _, r := range requests {
		response, err := GetOgpInfoByURL(r.URL, r.PostId)
		if err == nil {
			responses = append(responses, *response)
		}
	}

	return responses
}
