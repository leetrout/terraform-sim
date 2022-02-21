package api

import (
	"encoding/json"
	"net/http"
	"reflect"

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
	e, ok := store.Global.Entities[uuid]
	if !ok {
		http.NotFound(w, r)
		return
	}

	store.RemoveEntity(e)
	json.NewEncoder(w).Encode(map[string]string{
		"removed": uuid,
	})
}

// EntityUpdate handles updating the given entity
func EntityUpdate(w http.ResponseWriter, r *http.Request) {
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
	existingEntity, ok := store.Global.Entities[uuid]
	if !ok {
		http.NotFound(w, r)
		return
	}

	newName, ok := payload["Name"]
	if ok {
		existingEntity.Name = newName.(string)
	}

	newTER, ok := payload["TurboEncabulationRate"]
	if ok && newTER != nil {
		newTERF := newTER.(float64)
		newTERI := int(newTERF)
		existingEntity.TurboEncabulationRate = newTERI
	}

	refRate, ok := payload["RefractionRate"]
	if ok {
		var maybeRefRate *int
		val := reflect.ValueOf(refRate)
		if val.CanFloat() {
			floatRefRate := refRate.(float64)
			intRefRate := int(floatRefRate)
			maybeRefRate = &intRefRate
		}
		existingEntity.RefractionRate = maybeRefRate
	}

	validate = validator.New()

	err = validate.Struct(existingEntity)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	store.AddEntity(existingEntity)
	w.WriteHeader(http.StatusCreated)
	util.MarkRespJSON(w)
	json.NewEncoder(w).Encode(existingEntity)
}
