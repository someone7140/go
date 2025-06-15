package domain

import (
	"strconv"
	"strings"
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
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumWeeklyDay {
		return checkWeeklyDay(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumWeeklyDayInterval {
		return checkWeeklyDayInterval(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumMonthOnce {
		return checkMonthOnce(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumMonthDate {
		return checkMonthDate(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumYearOnceDate {
		return checkYearDate(checkTarget, now)
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

// 週の曜日単位のタスクの期限が超過していないか
func checkWeeklyDay(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	// 現在の曜日
	nowWeekDay := int(now.Weekday())
	// 指定された曜日からチェック期限の日を取得
	weeklyDay, ok := checkTarget.DeadLineCheckSubSetting["weeklyDay"].(int)
	if !ok {
		return false
	}
	var targetWeekDate time.Time
	if nowWeekDay < weeklyDay {
		targetWeekDate = now.AddDate(0, 0, -(7 - (weeklyDay - nowWeekDay)))
	} else {
		targetWeekDate = now.AddDate(0, 0, -(nowWeekDay - weeklyDay))
	}

	diffHour := targetWeekDate.Sub(checkTarget.LatestExecDateTime).Hours()
	// 指定した日付から7日以内に実行しているか
	return diffHour <= 7*24
}

// 指定された週間隔のタスクの期限が超過していないか
func checkWeeklyDayInterval(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	// 指定指定された週間隔
	weekInterval, ok := checkTarget.DeadLineCheckSubSetting["weekInterval"].(float64)
	if !ok {
		return false
	}

	diffHour := now.Sub(checkTarget.LatestExecDateTime).Hours()
	// 指定した週以内に実行しているか
	return diffHour <= 7*24*weekInterval
}

// 月に1度実行されているか（先月以降に実行されているか）
func checkMonthOnce(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	// 現在日時から先月に戻る
	lastMonth := now.AddDate(0, -1, 0)
	// 先月の1日を取得
	firstOfLastMonth := time.Date(lastMonth.Year(), lastMonth.Month(), 1, 0, 0, 0, 0, time.Local)

	// 先月の1日以降に実行されているか
	return checkTarget.LatestExecDateTime.Sub(firstOfLastMonth).Hours() >= 0
}

// 月の指定した日に実行されているか
func checkMonthDate(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	// 指定指定された月の日
	targetMonthlyDay, ok := checkTarget.DeadLineCheckSubSetting["monthlyDay"].(int)
	if !ok {
		return false
	}
	// 現在日時の日より対象日が小さい場合
	if now.Day() > targetMonthlyDay {
		// 今月の1日を取得
		firstOfNowMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
		// 今月に実行済みか
		return checkTarget.LatestExecDateTime.Sub(firstOfNowMonth).Hours() >= 0
	}

	// 先月の該当日との比較
	lastMonth := now.AddDate(0, -1, 0)
	targetDayOfLastMonth := time.Date(lastMonth.Year(), lastMonth.Month(), targetMonthlyDay, 0, 0, 0, 0, time.Local)
	// 先月の該当日以降に実行済みか
	return checkTarget.LatestExecDateTime.Sub(targetDayOfLastMonth).Hours() >= 0
}

// 年の指定した日に実行されているか
func checkYearDate(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	// 指定指定された月日（MM-dd形式）
	targeDate, ok := checkTarget.DeadLineCheckSubSetting["yearDate"].(string)
	if !ok {
		return false
	}
	// 月と日に分解
	targetMonthInt, err := strconv.Atoi(strings.Split(targeDate, "-")[0])
	if err != nil {
		return false
	}
	targetMonth := time.Month(targetMonthInt)
	targetDay, err := strconv.Atoi(strings.Split(targeDate, "-")[1])
	if err != nil {
		return false
	}

	// 現在の方が月が小さい場合
	if now.Month() < targetMonth {
		// 昨年の該当日との比較
		beforeYearDate := time.Date(now.Year()-1, targetMonth, targetDay, 0, 0, 0, 0, time.Local)
		// 実行済みか
		return checkTarget.LatestExecDateTime.Sub(beforeYearDate).Hours() >= 0
	} else if now.Month() > targetMonth { // 現在の方が月が大きい場合
		// 今年の該当日との比較
		thisYearDate := time.Date(now.Year(), targetMonth, targetDay, 0, 0, 0, 0, time.Local)
		// 実行済みか
		return checkTarget.LatestExecDateTime.Sub(thisYearDate).Hours() >= 0
	} else { // 月が同じ場合
		// 現在の日以下の場合
		if now.Day() <= targetDay {
			// 昨年の該当日との比較
			beforeYearDate := time.Date(now.Year()-1, targetMonth, targetDay, 0, 0, 0, 0, time.Local)
			// 実行済みか
			return checkTarget.LatestExecDateTime.Sub(beforeYearDate).Hours() >= 0
		} else { // 現在の日付が大きい場合
			// 今年の該当日との比較
			thisYearDate := time.Date(now.Year(), targetMonth, targetDay, 0, 0, 0, 0, time.Local)
			// 実行済みか
			return checkTarget.LatestExecDateTime.Sub(thisYearDate).Hours() >= 0
		}
	}
}
