package repository

import (
	"time"
	db "weather-api/src/db/model"
	"weather-api/src/util"

	"xorm.io/xorm"
)

// 地点の追加
func AddGeographicPoint(engine *xorm.Engine, userId string, name string, lat float64, lon float64, displayOrder int32) error {
	nowTime := time.Now().UTC().Unix()
	uuid, err := util.GenerateUUID()

	if err != nil {
		return err
	}
	geographicPoint := db.GeographicPoint{
		Id:               uuid,
		Name:             name,
		UserId:           userId,
		Lat:              lat,
		Lon:              lon,
		DisplayOrder:     displayOrder,
		RegisterDateTime: nowTime,
	}
	_, err = engine.Table("GEOGRAPHIC_POINT").Insert(geographicPoint)
	if err != nil {
		return err
	}
	return nil
}

// ユーザが登録した地点の一覧
func GetGeographicPointsByUserId(engine *xorm.Engine, userId string) ([]db.GeographicPoint, error) {
	points := []db.GeographicPoint{}
	err := engine.Table("GEOGRAPHIC_POINT").Where("USER_ID = ?", userId).OrderBy("DISPLAY_ORDER ASC").Find(&points)
	if err != nil {
		return nil, err
	}
	return points, nil
}
