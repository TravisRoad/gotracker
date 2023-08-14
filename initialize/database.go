package initialize

import (
	"github.com/TravisRoad/gotracker/global"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitDatabase() *gorm.DB {
	var db *gorm.DB
	dbConfig := global.Config.Db

	switch dbConfig.Type {
	case "mysql":
		db = initMysql()
	case "sqlite":
		db = initSqlite()
	default:
		panic("db type error: " + dbConfig.Type)
	}

	return db
}

func initMysql() *gorm.DB {
	return nil
}

func initSqlite() *gorm.DB {
	dbConfig := global.Config.Db
	db, err := gorm.Open(sqlite.Open(dbConfig.Sqlite.Path), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}
