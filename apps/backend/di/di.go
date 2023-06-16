package di

import (
	"fmt"
	"log"
	"travisroad/gotracker/auth"
	"travisroad/gotracker/config"
	"travisroad/gotracker/controllers"
	"travisroad/gotracker/models"
	"travisroad/gotracker/services"

	"go.uber.org/dig"
	"gorm.io/gorm"
)

var C *dig.Container

func InitDI(configPath string) {
	c := dig.New()
	var err error

	err = c.Provide(func() (*config.Config, error) {
		return config.ConfigFromFile(configPath)
	})
	if err != nil {
		msg := fmt.Sprintf("config error: %s", err.Error())
		panic(msg)
	}

	err = c.Provide(func(cf *config.Config) (*gorm.DB, error) {
		log.Println("db path is", cf.DbPath)
		return models.ConnectSqlite(cf.DbPath)
	})
	if err != nil {
		msg := fmt.Sprintf("db error: %s", err.Error())
		panic(msg)
	}

	err = c.Provide(services.NewTmdbService)
	if err != nil {
		msg := fmt.Sprintf("tmdbService init %s", err.Error())
		panic(msg)
	}

	err = c.Provide(services.NewMovieService)
	if err != nil {
		msg := fmt.Sprintf("tmdbService init %s", err.Error())
		panic(msg)
	}

	err = c.Provide(auth.NewJWTAuthHelper)
	if err != nil {
		msg := fmt.Sprintf("jwtHelper %s", err.Error())
		panic(msg)
	}

	log.Print("DI Initialized")

	c.Invoke(func(db *gorm.DB) {
		models.DB = db
	})
	err = c.Invoke(controllers.Inject)
	if err != nil {
		panic(err.Error())
	}

	C = c
}
