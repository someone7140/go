package repository

import (
	"context"
	"main/db/db_model"
	"main/db/query"

	"gorm.io/gorm/clause"
)

func RegisterNovelContents(ctx context.Context, novelContents []*db_model.NovelContent) error {
	c := query.NovelContent
	return c.WithContext(ctx).Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(novelContents...)
}

func GetMyNovelContents(ctx context.Context, userAccountID string) ([]*db_model.NovelContent, error) {
	c := query.NovelContent
	var novelContents []*db_model.NovelContent
	err := c.WithContext(ctx).
		Where(c.OwnerUserAccountID.Eq(userAccountID)).
		UnderlyingDB().
		Order("display_order ASC NULLS LAST").
		Find(&novelContents).
		Error
	return novelContents, err
}

func GetNovelContentsByNovelID(ctx context.Context, userAccountID string, novelID string) ([]*db_model.NovelContent, error) {
	c := query.NovelContent
	var novelContents []*db_model.NovelContent
	err := c.WithContext(ctx).
		Where(c.OwnerUserAccountID.Eq(userAccountID), c.NovelID.Eq(novelID)).
		UnderlyingDB().
		Order("display_order ASC NULLS LAST").
		Find(&novelContents).
		Error
	return novelContents, err
}

func GetNovelContentsByParentID(ctx context.Context, userAccountID string, parentNovelContentsID string) ([]*db_model.NovelContent, error) {
	c := query.NovelContent
	var novelContents []*db_model.NovelContent
	err := c.WithContext(ctx).
		Where(c.ParentContentsID.Eq(parentNovelContentsID), c.OwnerUserAccountID.Eq(userAccountID)).
		UnderlyingDB().
		Order("display_order ASC NULLS LAST").
		Find(&novelContents).
		Error
	return novelContents, err
}

func DeleteNovelContents(ctx context.Context, userAccountID string, novelContentIDs []string) error {
	s := query.NovelSetting
	_, err := s.WithContext(ctx).
		Where(s.ID.In(novelContentIDs...), s.OwnerUserAccountID.Eq(userAccountID)).Delete()
	return err
}
