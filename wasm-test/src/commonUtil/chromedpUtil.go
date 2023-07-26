package commonUtil

import (
	"context"
	"os"

	"github.com/chromedp/chromedp"
)

// スクリーンショットを撮る
func GetScreenShot(ctx context.Context, tasks chromedp.Tasks) error {
	var buf []byte
	screenShotTasks := append(tasks, chromedp.FullScreenshot(&buf, 90))
	err := chromedp.Run(ctx, screenShotTasks)
	err = os.WriteFile("fullScreenshot.png", buf, 0o644)

	return err
}
