package db

import (
	"log"

	_ "github.com/mattn/go-oci8"
	"xorm.io/xorm"
)

func GetOracleDbEngine() (*xorm.Engine, error) {
	engine, err := xorm.NewEngine("oci8", "WEATHER_DB/password@localhost:1521/ORCLPDB1")
	if err != nil {
		log.Fatal(err)
	}
	return engine, err
}
