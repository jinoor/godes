package main

import (
	"module"
	"util"
)

func main() {
	module.DBStart()
	module.WorldStart()
	module.MapStart()
	module.InitProtoPool()
	util.Profile()
	module.GameStart()
}
