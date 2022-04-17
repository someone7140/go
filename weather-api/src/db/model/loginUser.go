package db

type LoginUser struct {
	Id                string `xorm:"ID"`
	Name              string `xorm:"NAME"`
	LastLoginDateTime int64  `xorm:"LAST_LOGIN_DATE_TIME"`
}
