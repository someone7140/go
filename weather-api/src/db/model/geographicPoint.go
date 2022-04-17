package db

type GeographicPoint struct {
	Id               string  `xorm:"ID"`
	Name             string  `xorm:"NAME"`
	UserId           string  `xorm:"USER_ID"`
	Lat              float64 `xorm:"LAT"`
	Lon              float64 `xorm:"LON"`
	RegisterDateTime int64   `xorm:"REGISTER_DATE_TIME"`
}
