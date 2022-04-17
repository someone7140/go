package repository

import (
	"time"
	db "weather-api/src/db/model"
	"weather-api/src/util"

	"xorm.io/xorm"
)

func AddGeographicPoint(engine *xorm.Engine, userId string, name string, lat float64, lon float64) error {
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
		RegisterDateTime: nowTime,
	}
	_, err = engine.Table("GEOGRAPHIC_POINT").Insert(geographicPoint)
	if err != nil {
		return err
	}
	return nil
}
