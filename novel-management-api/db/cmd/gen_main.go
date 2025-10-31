package main

import (
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// envの読み込み
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// データベース接続
	db, err := gorm.Open(postgres.Open(os.Getenv("DB_CONNECT")), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// Genの設定
	g := gen.NewGenerator(gen.Config{
		OutPath:       "db/query",
		ModelPkgPath:  "db/db_model",
		Mode:          gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.UseDB(db)

	// 全テーブルを生成
	g.ApplyBasic(g.GenerateAllTable()...)

	g.Execute()
}
