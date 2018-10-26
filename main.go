package main

import (
	"user_tdd/common"
	"user_tdd/routes"
)

func main() {
	routes := routes.GetRoutes()
	common.StartServer("8081", routes)
}

//https://github.com/eminetto/clean-architecture-go
//https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f

//https://github.com/bxcodec/go-clean-arch
// for infrastructre only https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/

//nice to read (do not apply it )https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1

//https://github.com/hatajoe/8am
