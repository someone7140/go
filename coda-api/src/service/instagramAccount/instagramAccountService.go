package instagramAccount

import (
	errorConstants "coda-api/src/constants"
	instagramAccountModel "coda-api/src/model/instagramAccount"
	instagramAccountRepository "coda-api/src/repository/instagramAccount"
	"coda-api/src/util"
	"encoding/json"
	"os"

	"github.com/koron/go-dproxy"
)

// AddInstagramAccount インスタグラムアカウントの追加サービス
func AddInstagramAccount(
	instagramUserName string,
	status string,
	gender string,
	silhouette string,
	height string,
	genre string,
) error {
	if instagramUserName == "" ||
		status == "" ||
		gender == "" ||
		silhouette == "" ||
		height == "" ||
		genre == "" {
		return errorConstants.ErrBadRequest
	}
	// すでに登録されているアカウントかチェック
	instagramAccountInfoResponse, err := instagramAccountRepository.GetGatherInstagramAccountInfoByUserName(instagramUserName)
	if err != nil {
		return err
	}
	if instagramAccountInfoResponse != nil {
		return errorConstants.ErrForbidden
	}

	return instagramAccountRepository.AddInstagramAccount(
		instagramUserName, status, gender, silhouette, height, genre,
	)
}

// EditInstagramAccount インスタグラムアカウント情報の編集サービス
func EditInstagramAccount(
	id string,
	status string,
	gender string,
	silhouette string,
	height string,
	genre string,
) error {
	if id == "" ||
		status == "" ||
		gender == "" ||
		silhouette == "" ||
		height == "" ||
		genre == "" {
		return errorConstants.ErrBadRequest
	}
	err := instagramAccountRepository.EditInstagramAccount(
		id,
		status,
		gender,
		silhouette,
		height,
		genre,
	)
	if err != nil {
		return err
	}

	return nil
}

// GetGatherInstagramAccountInfoByUserName インスタグラムアカウント情報の取得サービス
func GetGatherInstagramAccountInfoByUserName(instagramUserName string) (*instagramAccountModel.InstagramAccountInfoResponse, error) {
	if instagramUserName == "" {
		return nil, errorConstants.ErrBadRequest
	}
	instagramAccountInfoResponse, err := instagramAccountRepository.GetGatherInstagramAccountInfoByUserName(instagramUserName)
	if err != nil {
		return nil, err
	}
	if instagramAccountInfoResponse == nil {
		return nil, errorConstants.ErrNotFound
	}

	return instagramAccountInfoResponse, nil
}

// GetGatherInstagramAccountWithPostsByUserName インスタグラムアカウントと投稿の取得サービス
func GetGatherInstagramAccountWithPostsByUserName(
	instagramUserName string,
	limit int,
	postStatus string,
) (*instagramAccountModel.InstagramAccountWithPosts, error) {
	if instagramUserName == "" || limit < 1 {
		return nil, errorConstants.ErrBadRequest
	}
	instagramAccountWithPosts, err := instagramAccountRepository.GetGatherInstagramAccountWithPostsByUserName(
		instagramUserName, limit, postStatus,
	)
	if err != nil {
		return nil, err
	}
	if instagramAccountWithPosts == nil {
		return nil, errorConstants.ErrNotFound
	}

	return instagramAccountWithPosts, nil
}

// GetInstagramFollowerCount インスタグラムのアカウントからフォロワー数を取得
func GetInstagramFollowerCount(instagramUserName string) (int64, error) {
	// インスタグラムのAPI
	url := os.Getenv("FACEBOOK_API_DOMAIN") +
		os.Getenv("INSTAGRAM_BUSINESS_ID") +
		"?fields=business_discovery.username(" +
		instagramUserName +
		"){followers_count}&access_token=" +
		os.Getenv("INSTAGRAM_ACCESS_TOKEN")
	resultStr, err := util.SendGetHTTPRequest(url)
	if err != nil {
		return 0, err
	}
	var responseInterface interface{}
	err = json.Unmarshal([]byte(resultStr), &responseInterface)
	if err != nil {
		return 0, err
	}
	result, err := dproxy.New(responseInterface).M("business_discovery").M("followers_count").Int64()
	if err != nil {
		return 0, err
	}
	return result, nil
}
