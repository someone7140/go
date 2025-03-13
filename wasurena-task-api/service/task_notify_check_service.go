package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"time"
	"wasurena-task-api/domain"
	"wasurena-task-api/middleware"
)

// contextのcloneのための定義
type customContext struct{ ctx context.Context }

func (customContext) Deadline() (time.Time, bool) { return time.Time{}, false }
func (customContext) Done() <-chan struct{}       { return nil }
func (customContext) Err() error                  { return nil }
func (x customContext) Value(key any) any         { return x.ctx.Value(key) }
func Clone(ctx context.Context) context.Context   { return customContext{ctx} }

// 日次での通知チェック
func CheckDailyNotify(ctx context.Context, token string) (bool, error) {
	// tokenの確認
	if token != os.Getenv("BATCH_TOKEN") {
		return false, errors.New("unauthorized token")
	}
	// contextをclone
	newCtx := Clone(ctx)
	// 非同期で処理実行
	go execCheckDailyNotify(newCtx)

	return true, nil
}

func execCheckDailyNotify(ctx context.Context) {
	// 現在時刻を取得
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		log.Fatal(err)
		return
	}
	now := time.Now().In(jst)
	var checkList []domain.TaskDeadLineCheckTarget

	// 0時の場合だけ全部の定義を対象にする予定（TODO）
	notifies, err := middleware.GetDbQueries(ctx).SelectLatestTaskExecuteForNotify(ctx)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, notify := range notifies {
		if notify.DeadLineCheck != nil {
			checkList = append(checkList, domain.TaskDeadLineCheckTarget{
				TaskID:                  notify.ID,
				TaskTitle:               notify.Title,
				DeadLineCheck:           *notify.DeadLineCheck,
				DeadLineCheckSubSetting: notify.DeadLineCheckSubSetting,
				OwnerUserID:             notify.OwnerUserID,
				ExecLatestDateTime:      notify.LatestDateTime,
			})
		}
	}

	deadLineCheck := domain.TaskDeadLineCheck{
		CheckTargetList: checkList,
		NowDateTime:     now,
	}
	// 通知対象ユーザの取得
	notificationUserMap, err := deadLineCheck.GetNotifyUserMap()
	if err != nil {
		log.Fatal(err)
		return
	}

	// 通知（TODO）
	fmt.Println(notificationUserMap)

}
