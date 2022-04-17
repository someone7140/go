package repository

import (
	"time"
	db "weather-api/src/db/model"

	"xorm.io/xorm"
)

func RegsiterUser(engine *xorm.Engine, userId string, userName string) error {
	nowTime := time.Now().UTC().Unix()

	registeredUser, err := SelectUser(engine, userId)
	if err != nil {
		return err
	}
	loginUser := db.LoginUser{
		Id:                userId,
		Name:              userName,
		LastLoginDateTime: nowTime,
	}

	if registeredUser == nil {
		_, err = engine.Table("LOGIN_USER").Insert(loginUser)
	} else {
		_, err = engine.Table("LOGIN_USER").Where("ID = ?", userId).Update(loginUser)
	}
	if err != nil {
		return err
	}
	return nil
}

func SelectUser(engine *xorm.Engine, userId string) (*db.LoginUser, error) {
	loginUser := db.LoginUser{}
	result, err := engine.Table("LOGIN_USER").Where("ID = ?", userId).Get(&loginUser)
	if err != nil {
		return nil, err
	}

	if !result {
		return nil, nil
	}
	return &loginUser, nil
}
