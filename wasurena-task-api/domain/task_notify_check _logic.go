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
	NextDeadLineDateTime    *time.Time
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
		checkResult := check.CheckTaskDeadLine(deadLineCheck.NowDateTime)
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

// タスクの期限超過チェック
func (checkTarget TaskDeadLineCheckTarget) CheckTaskDeadLine(now time.Time) bool {
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

// 次回の実行期限
func (checkTarget TaskDeadLineCheckTarget) GetNextDeadLineDateTime(now time.Time) *time.Time {
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumDailyOnce {
		return getNextDeadLineDateTimeDailyOnce(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumDailyHour {
		return getNextDeadLineDateTimeDailyHour(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumWeeklyDay {
		return getNextDeadLineDateTimeWeeklyDay(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumWeeklyDayInterval {
		return getNextDeadLineDateTimeWeeklyDayInterval(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumMonthOnce {
		return getNextDeadLineDateTimeMonthOnce(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumMonthDate {
		return getNextDeadLineDateTimeMonthDate(checkTarget, now)
	}
	if checkTarget.DeadLineCheck == db.DeadLineCheckEnumYearOnceDate {
		return getNextDeadLineDateTimeDateYearDate(checkTarget, now)
	}
	return nil
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

// 日単位のタスクの期限が超過していないか
func checkDailyOnce(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	diffHour := now.Sub(checkTarget.LatestExecDateTime).Hours()
	return diffHour <= 24
}

// 週の曜日単位のタスクの期限が超過していないか
func checkWeeklyDay(checkTarget TaskDeadLineCheckTarget, now time.Time) bool {
	// 現在の曜日
	nowWeekDay := int(now.Weekday())
	// 指定された曜日からチェック期限の日を取得
	weeklyDayFloat, ok := checkTarget.DeadLineCheckSubSetting["weeklyDay"].(float64)
	if !ok {
		return false
	}
	weeklyDay := int(weeklyDayFloat)

	// 設定曜日と現在曜日の比較
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
	// 今月の1日を取得してから1ヶ月戻る
	firstOfThisMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)
	firstOfLastMonth := firstOfThisMonth.AddDate(0, -1, 0)

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

// hour単位のタスクの次回実行期限
func getNextDeadLineDateTimeDailyHour(checkTarget TaskDeadLineCheckTarget, now time.Time) *time.Time {
	hourInterval, ok := checkTarget.DeadLineCheckSubSetting["hourInterval"].(float64)
	if !ok {
		return nil
	}

	diffHour := now.Sub(checkTarget.LatestExecDateTime).Hours()
	// すでに超過している場合は現在日付
	if diffHour > hourInterval {
		return &now
	} else {
		// 直近の実行日時から設定時間を足す
		nextDeadLineDateTime := checkTarget.LatestExecDateTime.Add(time.Duration(hourInterval) * time.Hour)
		return &nextDeadLineDateTime
	}
}

// 日単位のタスクの次回実行期限
func getNextDeadLineDateTimeDailyOnce(checkTarget TaskDeadLineCheckTarget, now time.Time) *time.Time {
	diffHour := now.Sub(checkTarget.LatestExecDateTime).Hours()
	// すでに超過している場合は現在日付
	if diffHour > 24 {
		return &now
	} else {
		// 直近の実行日時から24時間を足す
		nextDeadLineDateTime := checkTarget.LatestExecDateTime.Add(time.Duration(24) * time.Hour)
		return &nextDeadLineDateTime
	}

}

// 週の単位のタスクの次回実行期限
func getNextDeadLineDateTimeWeeklyDay(checkTarget TaskDeadLineCheckTarget, now time.Time) *time.Time {
	// 現在の曜日
	nowWeekDay := int(now.Weekday())
	// 指定された曜日からチェック期限の日を取得
	weeklyDayFloat, ok := checkTarget.DeadLineCheckSubSetting["weeklyDay"].(float64)
	if !ok {
		return nil
	}
	weeklyDay := int(weeklyDayFloat)

	// 設定曜日と現在曜日の比較
	var targetWeekDate time.Time
	if nowWeekDay == weeklyDay {
		targetWeekDate = now.AddDate(0, 0, 7)
	} else if nowWeekDay < weeklyDay {
		targetWeekDate = now.AddDate(0, 0, weeklyDay-nowWeekDay)
	} else {
		targetWeekDate = now.AddDate(0, 0, 7-(nowWeekDay-weeklyDay))
	}

	return &targetWeekDate
}

// 指定された週間隔の次回実行期限
func getNextDeadLineDateTimeWeeklyDayInterval(checkTarget TaskDeadLineCheckTarget, now time.Time) *time.Time {
	// 指定指定された週間隔
	weekInterval, ok := checkTarget.DeadLineCheckSubSetting["weekInterval"].(float64)
	if !ok {
		return nil
	}

	diffHour := now.Sub(checkTarget.LatestExecDateTime).Hours()
	// 指定した週以内に実行されていなければ現在日を返す
	if diffHour > 7*24*weekInterval {
		return &now
	} else {
		// 直近の実行日時から指定週を足す
		nextDeadLineDateTime := checkTarget.LatestExecDateTime.Add(time.Duration(7*24*weekInterval) * time.Hour)
		return &nextDeadLineDateTime
	}
}

// 月に1度実行設定の次回実行期限
func getNextDeadLineDateTimeMonthOnce(checkTarget TaskDeadLineCheckTarget, now time.Time) *time.Time {
	// 今月の1日
	firstOfThisMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, time.Local)

	// 今月の1日以降に実行されているか
	if checkTarget.LatestExecDateTime.Sub(firstOfThisMonth).Hours() >= 0 {
		// 来月の末日
		targetDateTime := firstOfThisMonth.AddDate(0, 2, -1)
		return &targetDateTime
	} else {
		// 今月の末日
		targetDateTime := firstOfThisMonth.AddDate(0, 1, -1)
		return &targetDateTime
	}
}

// 月の指定日の次回実行期限
func getNextDeadLineDateTimeMonthDate(checkTarget TaskDeadLineCheckTarget, now time.Time) *time.Time {
	// 指定指定された月の日
	targetMonthlyDay, ok := checkTarget.DeadLineCheckSubSetting["monthlyDay"].(int)
	if !ok {
		return nil
	}

	targetOfThisMonth := time.Date(now.Year(), now.Month(), targetMonthlyDay, 0, 0, 0, 0, time.Local)
	// 現在日時の日より対象日が小さい場合
	if now.Day() > targetMonthlyDay {
		// 来月の指定日
		targetOfNextMonth := targetOfThisMonth.AddDate(0, 1, 0)
		return &targetOfNextMonth
	}

	// 今月の指定日
	return &targetOfThisMonth
}

// 年の指定した日の次回実行期限
func getNextDeadLineDateTimeDateYearDate(checkTarget TaskDeadLineCheckTarget, now time.Time) *time.Time {
	// 指定指定された月日（MM-dd形式）
	targeDate, ok := checkTarget.DeadLineCheckSubSetting["yearDate"].(string)
	if !ok {
		return nil
	}
	// 月と日に分解
	targetMonthInt, err := strconv.Atoi(strings.Split(targeDate, "-")[0])
	if err != nil {
		return nil
	}
	targetMonth := time.Month(targetMonthInt)
	targetDay, err := strconv.Atoi(strings.Split(targeDate, "-")[1])
	if err != nil {
		return nil
	}

	targetOfThisMonth := time.Date(now.Year(), targetMonth, targetDay, 0, 0, 0, 0, time.Local)
	targetOfNextYearMonth := targetOfThisMonth.AddDate(1, 0, 0)

	if now.Month() < targetMonth {
		// まだその月が来てない場合は今年の日付
		return &targetOfThisMonth
	} else if now.Month() > targetMonth { // 現在の方が月が大きい場合
		return &targetOfNextYearMonth
	} else { // 月が同じ場合
		// 現在の日以下の場合
		if now.Day() <= targetDay {
			// 実行済みか
			return &targetOfThisMonth
		} else { // 現在の日付が大きい場合
			return &targetOfNextYearMonth
		}
	}
}
