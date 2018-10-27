package routes

import (
	"net/http"

)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc func(w http.ResponseWriter, r *http.Request)
}

func GetRoutes() []Route {

	var repoController = controllers.RespoController{repos.ReposClient{}}
	var appRoutes =[]Route{
		{http.MethodGet, "/repos", repoController.GetAllReposHandler},
	}
	return appRoutes
}
