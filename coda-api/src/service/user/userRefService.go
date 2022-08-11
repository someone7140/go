package user

import (
	errorConstants "coda-api/src/constants"
	userModel "coda-api/src/model/user"
	userRepository "coda-api/src/repository/user"
)

// GetUserInfoDetail ユーザの詳細情報取得
func GetUserInfoDetail(userID string) (*userModel.UserDetailInfoResponse, error) {
	if userID == "" {
		return nil, errorConstants.ErrBadRequest
	}
	return userRepository.GetUserEntityByUserIDForDetail(userID)
}
