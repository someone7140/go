package service

import (
	"context"
	"follow-check/src/model"
	"time"

	"github.com/chromedp/chromedp"
)

// インスタグラムのフォローチェック
func CheckInstagramFollow(userId string, password string) (model.InstagramCheckResult, error) {
	ctx, cancel := chromedp.NewContext(
		context.Background(),
	)
	defer cancel()

	tasks := chromedp.Tasks{
		chromedp.Navigate("https://www.instagram.com/accounts/login/"), // インスタグラムのログイン画面
		chromedp.Sleep(3000 * time.Millisecond),
	}
	err := chromedp.Run(ctx, tasks)
	// err := commonUtil.GetScreenShot(ctx, tasks)

	return model.InstagramCheckResult{
		Text: userId,
	}, err
}
