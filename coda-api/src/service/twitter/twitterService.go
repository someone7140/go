package twitter

import (
	"coda-api/src/util"
	"encoding/json"
	"os"

	"github.com/koron/go-dproxy"
)

// GetTwitterFollowerCount ページビューの取得
func GetTwitterFollowerCount(twitterUserName string) (int64, error) {
	// TwitterのURL
	url := "https://api.twitter.com/2/users/by/username/" + twitterUserName + "?user.fields=public_metrics"
	resultStr, err := util.SendGetHTTPRequestWithBearerToken(url, os.Getenv("TWITTER_BEARER_TOKEN"))
	if err != nil {
		return 0, err
	}
	var responseInterface interface{}
	err = json.Unmarshal([]byte(resultStr), &responseInterface)
	if err != nil {
		return 0, err
	}
	result, err := dproxy.New(responseInterface).M("data").M("public_metrics").M("followers_count").Int64()
	if err != nil {
		return 0, err
	}
	return result, nil
}
