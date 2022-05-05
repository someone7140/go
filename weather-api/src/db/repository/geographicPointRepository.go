package repository

import (
	"errors"
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

// 地点の更新
func UpdateGeographicPoint(
	engine *xorm.Engine,
	userId string,
	id string,
	name string,
	lat float64,
	lon float64,
	displayOrder int32,
) error {
	point := db.GeographicPoint{}
	result, err := engine.Table("GEOGRAPHIC_POINT").Where("ID = ?", id).And("USER_ID = ?", userId).Get(&point)
	if err != nil || !result {
		return errors.New("Not Found Error")
	}
	geographicPoint := db.GeographicPoint{
		Id:               id,
		Name:             name,
		UserId:           userId,
		Lat:              lat,
		Lon:              lon,
		DisplayOrder:     displayOrder,
		RegisterDateTime: point.RegisterDateTime,
	}
	_, err = engine.Table("GEOGRAPHIC_POINT").Where("ID = ?", id).And("USER_ID = ?", userId).Update(geographicPoint)
	if err != nil {
		return err
	}
	return nil
}

// 地点の削除
func DeleteGeographicPoint(
	engine *xorm.Engine,
	userId string,
	id string,
) error {
	point := db.GeographicPoint{}
	_, err := engine.Table("GEOGRAPHIC_POINT").Where("ID = ?", id).And("USER_ID = ?", userId).Delete(&point)
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
