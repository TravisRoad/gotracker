package utils

import (
	"os/exec"
	"travisroad/gotracker/config"
	"travisroad/gotracker/di"
)

func PreTest() {
	// dbFile := "/home/travis/Documents/shabby-toys/proj/gotracker/apps/backend/services/db.sqlite"
	// exec.Command("rm", "-f", dbFile).Run()
	configFile := "/tmp/config.yaml"

	di.InitDI(configFile)
}

func PostTest() {
	di.C.Invoke(func(config *config.Config) {
		dbFile := config.DbPath
		exec.Command("rm", "-f", dbFile).Run()
	})
}
