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

// GroupDelete presents the details of one entity.
func GroupDelete(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)

	uuidStrs := util.UUID_RX.FindStringSubmatch(r.URL.Path)
	if uuidStrs == nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	uuid := uuidStrs[0]
	_, ok := store.Global.Groups[uuid]
	if !ok {
		http.NotFound(w, r)
		return
	}

	delete(store.Global.Groups, uuid)
	json.NewEncoder(w).Encode(map[string]string{
		"removed": uuid,
	})
}

// GroupUpdate handles updating the given group
func GroupUpdate(w http.ResponseWriter, r *http.Request) {
	payload := map[string]any{}

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&payload)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	uuidStrs := util.UUID_RX.FindStringSubmatch(r.URL.Path)
	if uuidStrs == nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	uuid := uuidStrs[0]
	existingGroup, ok := store.Global.Groups[uuid]
	if !ok {
		http.NotFound(w, r)
		return
	}

	newName, ok := payload["name"]
	if ok {
		existingGroup.Name = newName.(string)
	}

	validate = validator.New()

	err = validate.Struct(existingGroup)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.AddGroup(existingGroup)
	w.WriteHeader(http.StatusCreated)
	util.MarkRespJSON(w)
	json.NewEncoder(w).Encode(existingGroup)

}
