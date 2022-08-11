package instagramBatch

import (
	instagramAccountModel "coda-api/src/model/instagramAccount"
	instagramAccountRepository "coda-api/src/repository/instagramAccount"
	postRepository "coda-api/src/repository/post"
	"coda-api/src/util"
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"time"

	"github.com/koron/go-dproxy"
	"github.com/thoas/go-funk"
)

// GatherInstagramPostService インスタの投稿収集サービス
func GatherInstagramPostService() error {
	// 収集アカウントのID取得
	accounts, err := instagramAccountRepository.GetGatherInstagramAccount()
	if err != nil {
		return err
	}
	for _, account := range accounts {
		// インスタグラムのAPIで投稿を取得
		posts, err := getInstagramPost(account.InstagramUserName)
		// エラーが発生したらそこで終了
		if err != nil {
			// エラー日付を更新
			instagramAccountRepository.UpdateErrorDate(account.ID)
			break
		}
		// 投稿を登録
		if len(posts) > 0 {
			err = registerInstagramPosts(posts, account.ID)
			if err != nil {
				return err
			}
			// 収集日を更新
			instagramAccountRepository.UpdateGatherDate(account.ID)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

// getInstagramPost インスタグラムのAPIから最新の投稿を取得する
func getInstagramPost(instagramUserName string) ([]instagramAccountModel.InstagramPost, error) {
	// インスタグラムのAPI
	url := os.Getenv("FACEBOOK_API_DOMAIN") +
		os.Getenv("INSTAGRAM_BUSINESS_ID") +
		"?fields=business_discovery.username(" +
		instagramUserName +
		"){media{caption,permalink,timestamp}}&access_token=" +
		os.Getenv("INSTAGRAM_ACCESS_TOKEN")
	resultStr, err := util.SendGetHTTPRequest(url)
	if err != nil {
		return nil, err
	}
	var responseInterface interface{}
	err = json.Unmarshal([]byte(resultStr), &responseInterface)
	if err != nil {
		return nil, err
	}
	datas, err := dproxy.New(responseInterface).M("business_discovery").M("media").M("data").ProxySet().MapArray()
	if err != nil {
		return nil, err
	}
	length := reflect.ValueOf(datas).Len()
	maxArrayIndex := map[bool]int{true: 10, false: length}[length > 10]
	var posts []instagramAccountModel.InstagramPost

	for i := 0; i < maxArrayIndex; i++ {
		data := dproxy.New(responseInterface).M("business_discovery").M("media").M("data").A(i)
		id, err := data.M("id").String()
		if err != nil {
			break
		}
		permalink, err := data.M("permalink").String()
		if err != nil {
			break
		}
		timestamp, err := data.M("timestamp").String()
		if err != nil {
			break
		}
		postDate, err := time.Parse("2006-01-02T15:04:05.000000Z", strings.Replace(timestamp, "+0000", "", -1)+".000000Z")
		posts = append(posts, instagramAccountModel.InstagramPost{
			ID:        id,
			Permalink: permalink,
			PostDate:  postDate,
		})
	}
	return posts, nil
}

// registerInstagramPosts インスタグラムの投稿を登録
func registerInstagramPosts(posts []instagramAccountModel.InstagramPost, userID string) error {
	// 登録済みの投稿ID
	registeredIds, err := postRepository.GetPostsByIds(
		funk.Map(posts, func(p instagramAccountModel.InstagramPost) string {
			return p.ID
		}).([]string))
	if err != nil {
		return err
	}
	// 登録する投稿
	newRegisterPosts := funk.Filter(posts, func(p instagramAccountModel.InstagramPost) bool {
		return !funk.ContainsString(registeredIds, p.ID)
	}).([]instagramAccountModel.InstagramPost)
	if len(newRegisterPosts) > 0 {
		err = postRepository.RegisterPosts(newRegisterPosts, userID)
		if err != nil {
			return err
		}
	}
	return nil
}
