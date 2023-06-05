package models

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectSqlite(sqlitePath string) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(sqlitePath), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	migrate(db)
	return db, err
}

func migrate(db *gorm.DB) {
	DB.AutoMigrate(
		&User{}, &Review{}, &Collection{}, &Seen{},
		&Metadata{}, &Keyword{},
		&Movie{}, &Book{}, &Game{}, &Manga{}, &Policy{})
}
