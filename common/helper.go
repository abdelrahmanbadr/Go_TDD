package common

import (
	"encoding/json"
	"github.com/globalsign/mgo"
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"user_tdd/routes"
)

type DbConfig struct {
	DbHost string
	DbName string
}

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

func GetMongoDbConnection(dbConfig DbConfig) (*mgo.Database, error) {
	session, err := mgo.Dial(dbConfig.DSN)
	if err != nil {
		log.Fatalf(CANT_NOT_CONNECT_DB, dbConfig.DSN, err)
		return nil, err
	}
	return session.DB(dbConfig.DbName), nil
}
