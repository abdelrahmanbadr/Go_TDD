package common

import (
	"github.com/go-chi/chi"
	"net/http"
	"tdd/routes"
)

func StartServer(port string, routes []routes.Route) {

	http.ListenAndServe(":"+port, GetRouter(routes))
}
func GetRouter(routes []routes.Route) *chi.Mux {
	r := chi.NewRouter()
	for _, route := range routes {
		r.MethodFunc(route.Method, route.Pattern, route.HandlerFunc)
	}
	return r

}
