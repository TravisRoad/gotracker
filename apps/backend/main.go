package main

import (
	"flag"
	"travisroad/gotracker/di"
	"travisroad/gotracker/route"
)

var (
	configFile string
)

func init() {
	flag.StringVar(&configFile, "conf", "config.yaml", "config file path")
	flag.Parse()
}

func main() {
	di.InitDI(configFile)

	r := route.RouteInit()
	r.Run()
}
