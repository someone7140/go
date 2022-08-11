package util

import (
	"time"
)

// GetMongoDateIntToTime Mongoから取得したInt型の日付をTime型に変換
func GetMongoDateIntToTime(dateMongoInt int64) time.Time {
	return time.Unix(dateMongoInt/1000, 0).UTC()
}
