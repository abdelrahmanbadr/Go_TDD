package common

import (
	"github.com/go-chi/chi"
	"net/http"
	"user_tdd/routes"
	"encoding/json"
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
func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
