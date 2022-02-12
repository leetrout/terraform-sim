package util

import "net/http"

func MarkRespJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// CORS config...
	// TODO: Localhost not allowed so using * but should make another util
	// and check if the incoming host is localhost?
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
