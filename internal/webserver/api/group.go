package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/leetrout/terraform-sim/internal/resources"
	"github.com/leetrout/terraform-sim/internal/store"
	"github.com/leetrout/terraform-sim/internal/webserver/util"
)

func GroupCreate(w http.ResponseWriter, r *http.Request) {
	var ag resources.APICreateGroup

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&ag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	validate = validator.New()
	err = validate.Struct(ag)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	g := ag.AsGroup()
	store.AddGroup(g)
	w.WriteHeader(http.StatusCreated)
	util.MarkRespJSON(w)
	json.NewEncoder(w).Encode(g)

}

func GroupDetail(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	fmt.Fprint(w, "TODO")
}

func GroupList(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	json.NewEncoder(w).Encode(store.Global.GroupList())
}
