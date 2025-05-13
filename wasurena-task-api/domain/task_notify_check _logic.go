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
	LatestExecDateTime      time.Time
}

type TaskDeadLineCheck struct {
	CheckTargetList []TaskDeadLineCheckTarget
	NowDateTime     time.Time
}

type CheckSubSettingDailyHour struct {
	HourInterval int
}

// タスクの期限が超過しているユーザのリストを取得
func (deadLineCheck TaskDeadLineCheck) GetNotifyUserMap() map[string][]TaskDeadLineCheckTarget {

	userNotifyMap := map[string][]TaskDeadLineCheckTarget{}
	for _, check := range deadLineCheck.CheckTargetList {
		// チェック処理
		checkResult := checkTaskDeadLine(check, deadLineCheck.NowDateTime)
		// 期限が超過している場合はユーザーごとにmapに追加
		if !checkResult {
			userNotifies, ok := userNotifyMap[check.OwnerUserID]
			if ok {
				userNotifyMap[check.OwnerUserID] = append(userNotifies, check)
			} else {
				userNotifyMap[check.OwnerUserID] = []TaskDeadLineCheckTarget{check}
			}
		}
	}
	return userNotifyMap
}

// タスクの期限が超過しているリストを生成
func (deadLineCheck TaskDeadLineCheck) GetExceededTaskList() []TaskDeadLineCheckTarget {

	exceededTaskList := []TaskDeadLineCheckTarget{}
	for _, check := range deadLineCheck.CheckTargetList {
		// チェック処理
		checkResult := checkTaskDeadLine(check, deadLineCheck.NowDateTime)
		// 期限が超過している場合はリストに追加
		if !checkResult {
			exceededTaskList = append(exceededTaskList, check)
		}
	}
	return exceededTaskList
}

// タスクの期限超過チェック
func checkTaskDeadLine(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumDailyOnce {
		return checkDailyOnce(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumDailyHour {
		return checkDailyHour(checkTarget, now)
	}
	return false
}

// hour単位のタスクの期限が超過していないか
func checkDailyHour(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	hourInterval, ok := checkTarget.DeadLineCheckSubSetting["hourInterval"].(float64)
	if !ok {
		return false
	}
	diffHour := now.Sub(checkTarget.LatestExecDateTime).Hours()

	return diffHour < hourInterval
}

// hour単位のタスクの期限が超過していないか
func checkDailyOnce(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	diffHour := now.Sub(checkTarget.LatestExecDateTime).Hours()
	return diffHour <= 24
}
