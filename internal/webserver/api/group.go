package api

import (
	"encoding/json"
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

// GroupDetail presents the details of one entity.
func GroupDetail(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)

	uuidStrs := util.UUID_RX.FindStringSubmatch(r.URL.Path)
	if uuidStrs == nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	uuid := uuidStrs[0]
	g, ok := store.Global.Groups[uuid]
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(g)
}

func GroupList(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	json.NewEncoder(w).Encode(store.Global.GroupList())
}
