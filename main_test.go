package main_test

import (
	"os"
	"testing"

	"github.com/TravisRoad/gotracker/initialize"
)

func TestMain(t *testing.T) {
	os.Setenv("MODE", "TEST")
	initialize.InitGlobal()
}
