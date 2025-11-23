package db_model

import "github.com/lib/pq"

const TableNameNovelSetting = "novel_settings"

// NovelSetting mapped from table <novel_settings>
type NovelSetting struct {
	ID                 string         `gorm:"column:id;primaryKey" json:"id"`
	Name               string         `gorm:"column:name;not null" json:"name"`
	NovelID            string         `gorm:"column:novel_id;not null" json:"novel_id"`
	OwnerUserAccountID string         `gorm:"column:owner_user_account_id;not null" json:"owner_user_account_id"`
	ParentSettingID    *string        `gorm:"column:parent_setting_id" json:"parent_setting_id"`
	DisplayOrder       *int32         `gorm:"column:display_order" json:"display_order"`
	Attributes         pq.StringArray `gorm:"column:attributes;not null;type:text[]" json:"attributes"`
	Description        *string        `gorm:"column:description" json:"description"`
}

// TableName NovelSetting's table name
func (*NovelSetting) TableName() string {
	return TableNameNovelSetting
}
