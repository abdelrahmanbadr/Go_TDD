package controllers

import (
	"encoding/json"
	"net/http"
	"tdd/interfaces"
)

type RespoController struct {
	Repos interfaces.Client
}

func (a *RespoController) GetAllReposHandler(w http.ResponseWriter, r *http.Request) {
	//user := chi.URLParam(r, "user")
	//fmt.Println(user)
	//if user == "" {
	//	http.Error(w, "MISSING_ARG_USER", 400)
	//	return
	//}

	repos, err := a.Repos.Get("")

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	RespondWithJSON(w, http.StatusOK, repos)
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
