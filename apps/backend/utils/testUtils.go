package utils

import (
	"os/exec"
	"travisroad/gotracker/di"
)

func PreTest() {
	dbFile := "/tmp/sqlmock_db.sqlite"
	configFile := "/tmp/config.yaml"
	exec.Command("rm", "-f", dbFile).Run()

	di.InitDI(configFile)
}
