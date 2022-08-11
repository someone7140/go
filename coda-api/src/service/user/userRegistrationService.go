package user

import (
	errorConstants "coda-api/src/constants"
	authModel "coda-api/src/model/auth"
	userModel "coda-api/src/model/user"
	coordinateRepository "coda-api/src/repository/coordinate"
	itemPostRepository "coda-api/src/repository/itemPost"
	postRepository "coda-api/src/repository/post"
	userRepository "coda-api/src/repository/user"
	"coda-api/src/service/auth"
	"coda-api/src/util"
	"mime/multipart"
)

// RegisterByGoogleLogin Googleログインユーザの新規登録
func RegisterByGoogleLogin(
	request userModel.RegisterUserByGoogleLoginRequest,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) (*authModel.AuthUserResponse, error) {
	if !ValidationNewUserRequest(request.UserInfo) {
		return nil, errorConstants.ErrBadRequest
	}
	if !userRepository.CheckUserSettingIDDuplicate(request.UserInfo.UserSettingID) {
		return nil, errorConstants.ErrForbidden
	}
	googleID := auth.GetGmailFromGoogleIDToken(request.GoogleIDToken)
	if googleID == "" {
		return nil, errorConstants.ErrUnauthorized
	}
	uuid, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	iconURL := ""
	// ファイルがある場合はアップロード処理
	if file != nil {
		filePath := util.GetFilePathForGcs(
			uuid, fileHeader, "icon",
		)
		iconURL, err = util.UploadNewFileToGcs(filePath, file)
		if err != nil {
			return nil, err
		}
	}
	err = userRepository.NewRegistrationUser(
		userModel.UserRegisterEntity{
			ID:            uuid,
			GoogleID:      googleID,
			UserSettingID: request.UserInfo.UserSettingID,
			Name:          request.UserInfo.Name,
			Gender:        request.UserInfo.Gender,
			BirthDate:     request.UserInfo.BirthDate,
			Silhouette:    request.UserInfo.Silhouette,
			Height:        request.UserInfo.Height,
			Weight:        request.UserInfo.Weight,
			Genres:        request.UserInfo.Genres,
			Complexes:     request.UserInfo.Complexes,
			Categories: util.GetCategoriesFromAttributeInput(
				request.UserInfo.Gender, request.UserInfo.Silhouette, request.UserInfo.Height, request.UserInfo.Genres,
			),
			Status:   "active",
			UserType: "normal",
			IconURL:  iconURL,
		})
	if err != nil {
		return nil, err
	}
	entity, err := userRepository.GetUserEntityByGoogleID(googleID)
	if err != nil {
		return nil, err
	}
	return auth.GetResponseFromEntity(entity), nil
}

// RegisterByFacdebookLogin Facebookログインユーザの新規登録
func RegisterByFacdebookLogin(
	request userModel.RegisterUserByFacebookLoginRequest,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) (*authModel.AuthUserResponse, error) {
	if !ValidationNewUserRequest(request.UserInfo) {
		return nil, errorConstants.ErrBadRequest
	}
	if !userRepository.CheckUserSettingIDDuplicate(request.UserInfo.UserSettingID) {
		return nil, errorConstants.ErrForbidden
	}
	facebookID := auth.GetIDFromFacebookAccessToken(request.FacebookAccessToken)
	if facebookID == "" {
		return nil, errorConstants.ErrUnauthorized
	}
	uuid, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	iconURL := ""
	// ファイルがある場合はアップロード処理
	if file != nil {
		filePath := util.GetFilePathForGcs(
			uuid, fileHeader, "icon",
		)
		iconURL, err = util.UploadNewFileToGcs(filePath, file)
		if err != nil {
			return nil, err
		}
	}
	err = userRepository.NewRegistrationUser(
		userModel.UserRegisterEntity{
			ID:            uuid,
			FacebookID:    facebookID,
			UserSettingID: request.UserInfo.UserSettingID,
			Name:          request.UserInfo.Name,
			Gender:        request.UserInfo.Gender,
			BirthDate:     request.UserInfo.BirthDate,
			Silhouette:    request.UserInfo.Silhouette,
			Height:        request.UserInfo.Height,
			Weight:        request.UserInfo.Weight,
			Genres:        request.UserInfo.Genres,
			Complexes:     request.UserInfo.Complexes,
			Categories: util.GetCategoriesFromAttributeInput(
				request.UserInfo.Gender, request.UserInfo.Silhouette, request.UserInfo.Height, request.UserInfo.Genres,
			),
			Status:   "active",
			UserType: "normal",
			IconURL:  iconURL,
		})
	if err != nil {
		return nil, err
	}
	entity, err := userRepository.GetUserEntityByFacebookID(facebookID)
	if err != nil {
		return nil, err
	}
	return auth.GetResponseFromEntity(entity), nil
}

// RegisterByEmailAuth メール認証ユーザの新規登録
func RegisterByEmailAuth(
	request userModel.RegisterUserByEmailAuthRequest,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) (*authModel.AuthUserResponse, error) {
	if !ValidationNewUserRequest(request.UserInfo) {
		return nil, errorConstants.ErrBadRequest
	}
	if !userRepository.CheckUserSettingIDDuplicate(request.UserInfo.UserSettingID) {
		return nil, errorConstants.ErrForbidden
	}
	temporaryUser, err := userRepository.GetUserEntityByUserID(request.UserID, "email")
	if err != nil {
		return nil, err
	}

	if temporaryUser == nil {
		return nil, errorConstants.ErrNotFound
	}

	if temporaryUser.Status != "confirming" || temporaryUser.EmailAuth.Token != request.Token {
		return nil, errorConstants.ErrUnauthorized
	}
	// 新規uuidの発行
	uuid, err := util.GenerateUUID()
	if err != nil {
		return nil, err
	}
	iconURL := ""
	// ファイルがある場合はアップロード処理
	if file != nil {
		filePath := util.GetFilePathForGcs(
			uuid, fileHeader, "icon",
		)
		iconURL, err = util.UploadNewFileToGcs(filePath, file)
		if err != nil {
			return nil, err
		}
	}
	// temporaryUserの削除
	err = userRepository.DeleteUser(request.UserID)
	if err != nil {
		return nil, err
	}

	err = userRepository.NewRegistrationUser(
		userModel.UserRegisterEntity{
			ID:            uuid,
			Email:         temporaryUser.Email,
			Password:      temporaryUser.Password,
			UserSettingID: request.UserInfo.UserSettingID,
			Name:          request.UserInfo.Name,
			Gender:        request.UserInfo.Gender,
			BirthDate:     request.UserInfo.BirthDate,
			Silhouette:    request.UserInfo.Silhouette,
			Height:        request.UserInfo.Height,
			Weight:        request.UserInfo.Weight,
			Genres:        request.UserInfo.Genres,
			Complexes:     request.UserInfo.Complexes,
			Categories: util.GetCategoriesFromAttributeInput(
				request.UserInfo.Gender, request.UserInfo.Silhouette, request.UserInfo.Height, request.UserInfo.Genres,
			),
			Status:   "active",
			UserType: "normal",
			IconURL:  iconURL,
		})
	if err != nil {
		return nil, err
	}
	entity, err := userRepository.GetUserEntityByUserID(uuid, "email")
	if err != nil {
		return nil, err
	}
	return auth.GetResponseFromEntity(entity), nil
}

// UpdateUserInfo ユーザ情報の更新
func UpdateUserInfo(
	loginInfo *authModel.AuthUserResponse,
	request userModel.UpdateUserInfoRequest,
	file multipart.File,
	fileHeader *multipart.FileHeader,
) (*authModel.AuthUserResponse, error) {
	if !ValidationNewUserRequest(request.UserInfo) {
		return nil, errorConstants.ErrBadRequest
	}

	if request.UserInfo.UserSettingID != loginInfo.UserSettingID &&
		!userRepository.CheckUserSettingIDDuplicate(request.UserInfo.UserSettingID) {
		return nil, errorConstants.ErrForbidden
	}
	var err error
	iconURL := loginInfo.IconURL
	// ファイルがある場合はアップロード処理
	if file != nil {
		// iconURLが登録されている場合は削除
		if iconURL != "" {
			fileName := util.GetFileNameFromURL(iconURL)
			err = util.DeleteFileFromGcs("icon/" + fileName)
			if err != nil {
				return nil, err
			}
		}
		filePath := util.GetFilePathForGcs(
			loginInfo.ID, fileHeader, "icon",
		)
		iconURL, err = util.UploadNewFileToGcs(filePath, file)
		if err != nil {
			return nil, err
		}
	}
	err = userRepository.UpdateUser(
		loginInfo.ID,
		userModel.UserRegisterEntity{
			UserSettingID: request.UserInfo.UserSettingID,
			Name:          request.UserInfo.Name,
			Gender:        request.UserInfo.Gender,
			BirthDate:     request.UserInfo.BirthDate,
			Silhouette:    request.UserInfo.Silhouette,
			Height:        request.UserInfo.Height,
			Weight:        request.UserInfo.Weight,
			Genres:        request.UserInfo.Genres,
			Complexes:     request.UserInfo.Complexes,
			IconURL:       iconURL,
			Categories: util.GetCategoriesFromAttributeInput(
				request.UserInfo.Gender, request.UserInfo.Silhouette, request.UserInfo.Height, request.UserInfo.Genres,
			),
		},
	)
	if err != nil {
		return nil, err
	}
	entity, err := userRepository.GetUserEntityByUserID(loginInfo.ID, loginInfo.AuthMethod)
	if err != nil {
		return nil, err
	}
	return auth.GetResponseFromEntity(entity), nil
}

// ValidationNewUserRequest ユーザ登録リクエストのチェック
func ValidationNewUserRequest(req userModel.UserInfoRequest) bool {
	if req.UserSettingID == "" || req.Name == "" || req.Gender == "" || req.Silhouette == "" {
		return false
	}
	return true
}

// DeleteUser ユーザの削除
func DeleteUser(loginInfo *authModel.AuthUserResponse) error {
	// インスタ投稿に対するいいねのレコードを削除
	err := postRepository.DeleteFavoriteUseAll(loginInfo.ID)
	if err != nil {
		return err
	}
	// アイテム投稿の画像削除
	imageURLs, err := itemPostRepository.GetImageURLsByUserID(loginInfo.ID)
	for _, url := range imageURLs {
		fileName := util.GetFileNameFromURL(url)
		util.DeleteFileFromGcs("itemPost/" + fileName)
	}

	// アイテム投稿の削除
	err = itemPostRepository.DeleteAllItemPostsByUser(loginInfo.ID)
	if err != nil {
		return err
	}
	// アイテム投稿に対するいいねのレコードを削除
	err = itemPostRepository.DeleteFavoriteItemPostUserAll(loginInfo.ID)
	if err != nil {
		return err
	}
	// iconURLが登録されている場合は削除
	if loginInfo.IconURL != "" {
		fileName := util.GetFileNameFromURL(loginInfo.IconURL)
		util.DeleteFileFromGcs("icon/" + fileName)
	}
	// コーデ投稿に対するいいねの削除
	err = coordinateRepository.DeleteFavoriteCoordinatePostUserAll(loginInfo.ID)
	if err != nil {
		return err
	}
	// ユーザの削除
	err = userRepository.DeleteUser(loginInfo.ID)
	if err != nil {
		return err
	}
	return nil
}
