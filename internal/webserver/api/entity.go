package api

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/leetrout/terraform-sim/internal/resources"
	"github.com/leetrout/terraform-sim/internal/webserver/util"
)

// EntityList lists all entities.
func EntityList(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	json.NewEncoder(w).Encode(
		[]*resources.Entity{
			resources.NewEntity(),
			resources.NewEntity(),
		},
	)
}

// EntityDetail presents the details
// of one entity.
func EntityDetail(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	json.NewEncoder(w).Encode(resources.NewEntity())
}

var validate *validator.Validate

// EntityCreate creates a new Entity
func EntityCreate(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)

	var e resources.Entity

	// Try to decode the request body into the struct. If there is an error,
	// respond to the client with the error message and a 400 status code.
	err := json.NewDecoder(r.Body).Decode(&e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// e := resources.NewEntity()

	validate = validator.New()
	err = validate.Struct(e)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	jsonOut, _ := json.Marshal(e)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonOut)
}
