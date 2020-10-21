package main

import (
	"go-project/Config"
	"go-project/Routes"
)

func main() {
	Config.DbConnection()
	r := Routes.SetupRouter()
	//running
	r.Run(":80")
}
