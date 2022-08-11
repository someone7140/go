package analytics

import (
	analyticsRepository "coda-api/src/repository/analytics"
	userRepository "coda-api/src/repository/user"
	"coda-api/src/service/googleAnalytics"
	"coda-api/src/service/instagramAccount"
	"coda-api/src/service/twitter"
	"os"
	"time"
)

// AccessAnalysisBatchService アクセス情報の収集サービス
func AccessAnalysisBatchService() error {
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	// JST今日日付
	currentTimeJst, err := time.ParseInLocation("20060102", time.Now().In(jst).Format("20060102"), jst)
	if err != nil {
		return err
	}
	// JST昨日日付
	yesterDayTimeJst := currentTimeJst.AddDate(0, 0, -1)
	// UTC時間での日本の今日日付
	currentTimeUtc := currentTimeJst.UTC()
	// UTC時間での日本の昨日日付
	yesterDayTimeUtc := yesterDayTimeJst.UTC()

	// 日付のキー
	yesterDayKey := yesterDayTimeJst.Format("20060102")

	// 本日日付までに登録された会員数
	totalUserCount, _ := userRepository.GetRegistrationUserCount(currentTimeUtc)
	// 昨日から本日にログインした会員数
	loginUserCount, _ := userRepository.GetLoginUserCount(yesterDayTimeUtc, currentTimeUtc)
	// GoogleAnalyticsから昨日のアクセスユーザ数を取得
	accessUserCount, _ := googleAnalytics.GetAccessUserCount(yesterDayTimeJst.Format("2006-01-02"))
	// instagramのフォロワー数を取得
	instagramFollowerCount, err := instagramAccount.GetInstagramFollowerCount(os.Getenv("CODA_INSTAGRAM_ID"))
	// twitterのフォロワー数を取得
	twitterFollowerCount, err := twitter.GetTwitterFollowerCount(os.Getenv("CODA_TWITTER_ID"))
	// 登録
	return analyticsRepository.AddDateAccessAnalytics(
		yesterDayKey,
		totalUserCount,
		loginUserCount,
		accessUserCount,
		instagramFollowerCount,
		twitterFollowerCount,
	)
}
