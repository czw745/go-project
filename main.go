package main

import (
	"go-project/config"
	"go-project/routes"
)

func main() {
	config.DbConnection()
	config.REDISConnection()
	r := routes.SetupRouter()
	//running
	r.Run(":80")
}
