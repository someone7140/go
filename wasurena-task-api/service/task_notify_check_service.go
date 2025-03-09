package service

import (
	"context"
	"errors"
	"os"
	"wasurena-task-api/db"
	"wasurena-task-api/middleware"
)

// 日次での通知チェック
func CheckDailyNotify(ctx context.Context, token string) (bool, error) {
	// tokenの確認
	if token != os.Getenv("BATCH_TOKEN") {
		return false, errors.New("unauthorized token")
	}

	notifies, err := middleware.GetDbQueries(ctx).SelectLatestTaskExecuteForNotify(ctx)
	if err != nil {
		return false, errors.New("can not  Db quey")
	}
	userNotifyMap := map[string][]db.DeadLineCheckEnum{}
	for _, notify := range notifies {
		if notify.DeadLineCheck != nil {
			// チェック処理
			checkResult := true
			if !checkResult {
				userNotifies, ok := userNotifyMap[notify.OwnerUserID]
				if ok {
					userNotifyMap[notify.OwnerUserID] = append(userNotifies, *notify.DeadLineCheck)
				} else {
					userNotifyMap[notify.OwnerUserID] = []db.DeadLineCheckEnum{*notify.DeadLineCheck}
				}
			}
		}
	}
	// 通知処理（TODO）

	return true, nil
}
