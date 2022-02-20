package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/leetrout/terraform-sim/internal/resources"
	"github.com/leetrout/terraform-sim/internal/store"
	"github.com/leetrout/terraform-sim/internal/webserver/util"
)

// EntityList lists all entities.
func EntityList(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)

	json.NewEncoder(w).Encode(
		store.Global.EntityList(),
	)
}

// EntityDetail presents the details of one entity.
func EntityDetail(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)

	uuidStrs := util.UUID_RX.FindStringSubmatch(r.URL.Path)
	if uuidStrs == nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	uuid := uuidStrs[0]
	e, ok := store.Global.Entities[uuid]
	if !ok {
		http.NotFound(w, r)
		return
	}

	json.NewEncoder(w).Encode(e)
}

var validate *validator.Validate

// EntityCreate creates a new Entity
func EntityCreate(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)

	var ae resources.APICreateEntity

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&ae)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// e := resources.NewEntity()

	validate = validator.New()
	err = validate.Struct(ae)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	e := ae.AsEntity()
	store.AddEntity(e)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(e)
}

// EntityDelete deletes the given entity
func EntityDelete(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)

	uuidStrs := util.UUID_RX.FindStringSubmatch(r.URL.Path)
	if uuidStrs == nil {
		http.Error(w, "invalid UUID", http.StatusBadRequest)
		return
	}

	uuid := uuidStrs[0]
	_, ok := store.Global.Entities[uuid]
	if !ok {
		http.NotFound(w, r)
		return
	}

	delete(store.Global.Entities, uuid)
	json.NewEncoder(w).Encode(map[string]string{
		"removed": uuid,
	})
}
