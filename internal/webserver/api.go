package webserver

import (
	"errors"
	"net/http"
	"strings"

	"github.com/leetrout/terraform-sim/internal/webserver/api"
)

// apiHandler handles routing API specific requests.
func apiHandler(w http.ResponseWriter, r *http.Request) {
	resource, err := splitAPIPath(strings.ToLower(r.URL.Path))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if resource == "entity" {
		switch r.Method {
		case http.MethodGet:
			api.EntityDetail(w, r)
			return
		case http.MethodPost:
			api.EntityCreate(w, r)
			return
		case http.MethodDelete:
			api.EntityDelete(w, r)
			return
		}
	}

	if resource == "entities" {
		switch r.Method {
		case http.MethodGet:
			api.EntityList(w, r)
			return
		}
	}

	if resource == "group" {
		switch r.Method {
		case http.MethodGet:
			api.GroupDetail(w, r)
			return
		case http.MethodPost:
			api.GroupCreate(w, r)
			return
		case http.MethodDelete:
			api.GroupDelete(w, r)
			return
		case http.MethodPatch:
			api.GroupUpdate(w, r)
			return
		}

	}

	if resource == "groups" {
		switch r.Method {
		case http.MethodGet:
			api.GroupList(w, r)
			return
		}
	}

	w.WriteHeader(http.StatusBadRequest)
}

// splitAPIPath returns the resource on which we are operating.
func splitAPIPath(urlPath string) (string, error) {
	// Strip leading slash
	urlPath = urlPath[1:]

	// Strip trailing slash if present
	lastIdx := len(urlPath) - 1
	if string(urlPath[lastIdx]) == "/" {
		urlPath = urlPath[:lastIdx]
	}

	// Split on any remaining slashes
	pathParts := strings.Split(urlPath, "/")

	if len(pathParts) < 2 {
		return "", errors.New("invalid path")
	}

	// Expecting [api, <resource>, ...] at this point
	// so return the resource name
	return pathParts[1], nil
}
