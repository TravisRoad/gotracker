package models

import (
	"strconv"

	"github.com/gin-gonic/gin"
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
	db.AutoMigrate(
		&User{}, &Review{}, &Collection{}, &Seen{},
		&Metadata{}, &Keyword{},
		&Movie{}, &Book{}, &Game{}, &Manga{}, &Policy{})
}

func Paginate(c *gin.Context) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(c.Query("page"))
		if page <= 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(c.Query("page_size"))
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
