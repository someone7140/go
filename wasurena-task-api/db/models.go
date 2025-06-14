// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package db

import (
	"database/sql/driver"
	"fmt"
	"time"

	db_type "wasurena-task-api/db/type"
)

type DeadLineCheckEnum string

const (
	DeadLineCheckEnumDailyOnce         DeadLineCheckEnum = "DailyOnce"
	DeadLineCheckEnumDailyHour         DeadLineCheckEnum = "DailyHour"
	DeadLineCheckEnumWeeklyDay         DeadLineCheckEnum = "WeeklyDay"
	DeadLineCheckEnumWeeklyDayInterval DeadLineCheckEnum = "WeeklyDayInterval"
	DeadLineCheckEnumMonthOnce         DeadLineCheckEnum = "MonthOnce"
	DeadLineCheckEnumMonthDate         DeadLineCheckEnum = "MonthDate"
	DeadLineCheckEnumYearOnceDate      DeadLineCheckEnum = "YearOnceDate"
)

func (e *DeadLineCheckEnum) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = DeadLineCheckEnum(s)
	case string:
		*e = DeadLineCheckEnum(s)
	default:
		return fmt.Errorf("unsupported scan type for DeadLineCheckEnum: %T", src)
	}
	return nil
}

type NullDeadLineCheckEnum struct {
	DeadLineCheckEnum DeadLineCheckEnum
	Valid             bool // Valid is true if DeadLineCheckEnum is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullDeadLineCheckEnum) Scan(value interface{}) error {
	if value == nil {
		ns.DeadLineCheckEnum, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.DeadLineCheckEnum.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullDeadLineCheckEnum) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.DeadLineCheckEnum), nil
}

type TaskCategory struct {
	ID           string
	Name         string
	OwnerUserID  string
	DisplayOrder *int32
}

type TaskDefinition struct {
	ID                      string
	Title                   string
	OwnerUserID             string
	DisplayFlag             bool
	NotificationFlag        bool
	CategoryID              *string
	DeadLineCheck           *DeadLineCheckEnum
	DeadLineCheckSubSetting db_type.Jsonb
	Detail                  *string
}

type TaskExecute struct {
	ID               string
	TaskDefinitionID string
	ExecuteUserID    string
	ExecuteDateTime  time.Time
	Memo             *string
}

type UserAccount struct {
	ID              string
	UserSettingID   string
	LineID          string
	UserName        string
	ImageUrl        *string
	IsLineBotFollow bool
}
