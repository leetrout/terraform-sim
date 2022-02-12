package util

import (
	"net/http"
	"regexp"
)

var UUID_RX = regexp.MustCompile(`[a-f0-9]{8}-[a-f0-9]{4}-4[a-f0-9]{3}-[89aAbB][a-f0-9]{3}-[a-f0-9]{12}$`)

func MarkRespJSON(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	// CORS config...
	// TODO: Localhost not allowed so using * but should make another util
	// and check if the incoming host is localhost?
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
