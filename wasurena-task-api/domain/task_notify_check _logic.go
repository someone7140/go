package domain

import (
	"time"
	"wasurena-task-api/db"
	db_type "wasurena-task-api/db/type"
)

type TaskDeadLineCheckTarget struct {
	TaskID                  string
	TaskTitle               string
	DeadLineCheck           db.DeadLineCheckEnum
	DeadLineCheckSubSetting db_type.Jsonb
	OwnerUserID             string
	ExecLatestDateTime      time.Time
}

type TaskDeadLineCheck struct {
	CheckTargetList []TaskDeadLineCheckTarget
	NowDateTime     time.Time
}

type CheckSubSettingDailyHour struct {
	HourInterval int
}

// タスクの期限が超過しているユーザのリストを取得
func (deadLineCheck TaskDeadLineCheck) GetNotifyUserMap() (map[string][]db.DeadLineCheckEnum, error) {

	userNotifyMap := map[string][]db.DeadLineCheckEnum{}
	for _, check := range deadLineCheck.CheckTargetList {
		// チェック処理
		checkResult := false
		if check.DeadLineCheck == db.DeadLineCheckEnumDailyOnce {
			checkResult = checkDailyOnce(check, deadLineCheck.NowDateTime)
		} else if check.DeadLineCheck == db.DeadLineCheckEnumDailyHour {
			checkResult = checkDailyHour(check, deadLineCheck.NowDateTime)
		}

		if !checkResult {
			userNotifies, ok := userNotifyMap[check.OwnerUserID]
			if ok {
				userNotifyMap[check.OwnerUserID] = append(userNotifies, check.DeadLineCheck)
			} else {
				userNotifyMap[check.OwnerUserID] = []db.DeadLineCheckEnum{check.DeadLineCheck}
			}
		}
	}
	return userNotifyMap, nil
}

// hour単位のタスクの期限が超過していないか
func checkDailyHour(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	hourInterval, ok := checkTarget.DeadLineCheckSubSetting["hourInterval"].(float64)
	if !ok {
		return false
	}
	diffHour := now.Sub(checkTarget.ExecLatestDateTime).Hours()

	return diffHour < hourInterval
}

// hour単位のタスクの期限が超過していないか
func checkDailyOnce(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	diffHour := now.Sub(checkTarget.ExecLatestDateTime).Hours()
	return diffHour <= 24
}
