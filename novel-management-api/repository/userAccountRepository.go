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

func AddUserAccount(ctx context.Context, userAccount db_model.UserAccount) error {
	u := query.UserAccount
	return u.WithContext(ctx).Create(&userAccount)
}
