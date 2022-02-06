package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/leetrout/terraform-sim/internal/resources"
	"github.com/leetrout/terraform-sim/internal/webserver/util"
)

// EntityList lists all entities.
func EntityList(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	fmt.Fprint(w, "List of entities goes here")
}

// EntityDetail presents the details
// of one entity.
func EntityDetail(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	fmt.Fprint(w, "entity detail goes here")
}

// EntityCreate creates a new Entity
func EntityCreate(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	e := resources.NewEntity()
	jsonOut, _ := json.Marshal(e)
	w.WriteHeader(http.StatusCreated)
	w.Write(jsonOut)
}
