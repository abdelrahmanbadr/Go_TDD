package main

import (
	"net/http"
	"user_tdd/common"
	"user_tdd/user/repository"
	"user_tdd/user/service"
)

func main() {
	http.HandleFunc("/users", GetAllUsers)
	http.ListenAndServe(":8080", nil)
	//routes := routes.GetRoutes()
	//common.StartServer("8081", routes)
}

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	repo := repository.NewInMemoryRepository()
	service := service.NewUserService(repo)
	users, _ := service.List()
	common.RespondWithJSON(w, http.StatusOK, users)
}

//https://github.com/eminetto/clean-architecture-go
//https://medium.com/@eminetto/clean-architecture-using-golang-b63587aa5e3f

//https://github.com/bxcodec/go-clean-arch
// for infrastructre only https://manuel.kiessling.net/2012/09/28/applying-the-clean-architecture-to-go-applications/

//nice to read (do not apply it )https://medium.com/@hatajoe/clean-architecture-in-go-4030f11ec1b1

//https://github.com/hatajoe/8am
