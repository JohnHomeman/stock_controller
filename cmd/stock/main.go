package main

import (
	"stock_controller/binder"
	"stock_controller/glob"
)

func main() {

	glob.InitYaml()
	binder.Run()
	select {}

}
