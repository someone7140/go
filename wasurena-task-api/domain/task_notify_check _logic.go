package service

import (
	"time"
	"wasurena-task-api/db"
)

// タスクの期限が超過しているか
func IsTaskDeadLineExceed(deadLineCheck db.DeadLineCheckEnum, latestDateTime time.Time) bool {
	return true
}
