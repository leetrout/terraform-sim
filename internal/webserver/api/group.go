package api

import (
	"fmt"
	"net/http"

	"github.com/leetrout/terraform-sim/internal/webserver/util"
)

func GroupCreate(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	fmt.Fprint(w, "TODO")
}

func GroupDetail(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	fmt.Fprint(w, "TODO")
}

func GroupList(w http.ResponseWriter, r *http.Request) {
	util.MarkRespJSON(w)
	fmt.Fprint(w, "TODO")
}
