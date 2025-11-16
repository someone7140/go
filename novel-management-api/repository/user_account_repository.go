package repository

import (
	"context"
	"main/db/db_model"
	"main/db/query"
)

func GetUserAccountByGmail(ctx context.Context, gmail string) (*db_model.UserAccount, error) {
	u := query.UserAccount
	return u.WithContext(ctx).Where(u.Gmail.Eq(gmail)).First()
}

func GetUserAccountByUserSettingID(ctx context.Context, userSettingID string) (*db_model.UserAccount, error) {
	u := query.UserAccount
	return u.WithContext(ctx).Where(u.UserSettingID.Eq(userSettingID)).First()
}

func GetUserAccountByID(ctx context.Context, userAccountID string) (*db_model.UserAccount, error) {
	u := query.UserAccount
	return u.WithContext(ctx).Where(u.ID.Eq(userAccountID)).First()
}

func AddUserAccount(ctx context.Context, userAccount db_model.UserAccount) error {
	u := query.UserAccount
	return u.WithContext(ctx).Create(&userAccount)
}

func UpdateUserAccountByInput(ctx context.Context, userAccountID string, userSettingID string, name string) error {
	u := query.UserAccount
	_, err := u.WithContext(ctx).
		Where(u.ID.Eq(userAccountID)).
		Updates(map[string]any{"user_setting_id": userSettingID, "name": name})
	return err
}

func UpdateUserAccountImageUrl(ctx context.Context, userAccountID string, imageURL string) error {
	u := query.UserAccount
	_, err := u.WithContext(ctx).
		Where(u.ID.Eq(userAccountID)).
		UpdateSimple(u.ImageURL.Value(imageURL))
	return err
}
