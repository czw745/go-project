package main

import (
	"fmt"
	"go-project/routes"
)

func main() {
	r := routes.InitRouter()
	fmt.Println(r)
}
