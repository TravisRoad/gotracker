package initialize

import (
	"github.com/TravisRoad/gotracker/global"
)

func InitGlobal() {
	global.Viper = initViper()
}
