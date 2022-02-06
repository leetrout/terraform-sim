package util

import "net/http"

func MarkRespJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
