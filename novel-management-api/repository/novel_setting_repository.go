package repository

import (
	"context"
	"main/db/db_model"
	"main/db/query"

	"gorm.io/gorm/clause"
)

func RegisterNovelSettings(ctx context.Context, novelSettings []*db_model.NovelSetting) error {
	s := query.NovelSetting
	return s.WithContext(ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(novelSettings...)
}

func GetMyNovelSettings(ctx context.Context, userAccountID string) ([]*db_model.NovelSetting, error) {
	s := query.NovelSetting
	var novelSettings []*db_model.NovelSetting
	err := s.WithContext(ctx).
		Where(s.OwnerUserAccountID.Eq(userAccountID)).
		UnderlyingDB().
		Order("display_order ASC NULLS LAST").
		Find(&novelSettings).
		Error
	return novelSettings, err
}

func GetNovelSettingsByNovelID(ctx context.Context, userAccountID string, novelID string) ([]*db_model.NovelSetting, error) {
	s := query.NovelSetting
	var novelSettings []*db_model.NovelSetting
	err := s.WithContext(ctx).
		Where(s.OwnerUserAccountID.Eq(userAccountID), s.NovelID.Eq(novelID)).
		UnderlyingDB().
		Order("display_order ASC NULLS LAST").
		Find(&novelSettings).
		Error
	return novelSettings, err
}

func GetNovelSettingsByParentID(ctx context.Context, userAccountID string, parentNovelSettingID string) ([]*db_model.NovelSetting, error) {
	s := query.NovelSetting
	var novelSettings []*db_model.NovelSetting
	err := s.WithContext(ctx).
		Where(s.ParentSettingID.Eq(parentNovelSettingID), s.OwnerUserAccountID.Eq(userAccountID)).
		UnderlyingDB().
		Order("display_order ASC NULLS LAST").
		Find(&novelSettings).
		Error
	return novelSettings, err
}

func DeleteNovelSettings(ctx context.Context, userAccountID string, novelSettingIDs []string) error {
	s := query.NovelSetting
	_, err := s.WithContext(ctx).
		Where(s.ID.In(novelSettingIDs...), s.OwnerUserAccountID.Eq(userAccountID)).Delete()
	return err
}
