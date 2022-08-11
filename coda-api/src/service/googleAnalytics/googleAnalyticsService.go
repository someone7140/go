package googleAnalytics

import (
	"context"
	"os"
	"strconv"

	"google.golang.org/api/analytics/v3"
	"google.golang.org/api/option"
)

// GetAccessUserCount アクセスユーザ数の取得
func GetAccessUserCount(yesterdayStr string) (int64, error) {
	ctx := context.Background()
	credentialFilePath := "./gcsKey/" + os.Getenv("GCS_KEY_FILE")
	client, err := analytics.NewService(ctx, option.WithCredentialsFile(credentialFilePath))
	if err != nil {
		return 0, err
	}
	resPVs, err := client.Data.Ga.Get(
		"ga:"+os.Getenv("GA_VIEW_ID"), yesterdayStr, yesterdayStr, "ga:users").Do()
	if err != nil {
		return 0, err
	}
	result, err := strconv.ParseInt(resPVs.TotalsForAllResults["ga:users"], 10, 64)
	if err != nil {
		return 0, err
	}
	return result, nil
}
