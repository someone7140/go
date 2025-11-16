package repository

import (
	"context"
	"main/db/db_model"
	"main/db/query"
)

func AddNovel(ctx context.Context, novel db_model.Novel) error {
	n := query.Novel
	return n.WithContext(ctx).Create(&novel)
}

func EditNovel(ctx context.Context, novelID, userAccountID string, title string, description *string) error {
	n := query.Novel
	_, err := n.WithContext(ctx).
		Where(n.ID.Eq(novelID), n.OwnerUserAccountID.Eq(userAccountID)).
		Updates(map[string]any{"title": title, "description": description})
	return err
}

func DeleteNovel(ctx context.Context, userAccountID string, novelID string) error {
	n := query.Novel
	_, err := n.WithContext(ctx).
		Where(n.ID.Eq(novelID), n.OwnerUserAccountID.Eq(userAccountID)).Delete()
	return err
}

func GetMyNovels(ctx context.Context, userAccountID string) ([]*db_model.Novel, error) {
	n := query.Novel
	return n.WithContext(ctx).
		Where(n.OwnerUserAccountID.Eq(userAccountID)).Find()
}

func GetMyNovelByID(ctx context.Context, userAccountID string, novelID string) (*db_model.Novel, error) {
	n := query.Novel
	return n.WithContext(ctx).
		Where(n.ID.Eq(novelID), n.OwnerUserAccountID.Eq(userAccountID)).First()
}
