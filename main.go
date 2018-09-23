package main

import (
	"tdd/routes"
	"tdd/common"
)

func main() {
	routes := routes.GetRoutes()
	common.StartServer("8081", routes)
}

