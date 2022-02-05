package webserver

import (
	"fmt"
	"log"
	"mime"
	"net/http"
	"strings"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var html, err = f.ReadFile("static/index.html")
	if err != nil {
		// Fixme
		return
	}
	fmt.Fprintf(w, string(html))
}

func staticHandler(w http.ResponseWriter, r *http.Request) {
	// Fix me - find a better URL parser
	var parts = strings.Split(r.URL.Path[1:], "/")
	fmt.Println(len(parts))
	if len(parts) < 2 {
		// 404?
		return
	}
	fmt.Println(parts)
	var file = strings.Join(parts, "/")
	data, _ := f.ReadFile(file)
	var extParts = strings.Split(file, ".")
	if len(extParts) > 0 {
		var ext = extParts[len(extParts)-1]
		ext = fmt.Sprintf(".%s", ext)
		var ct = mime.TypeByExtension(ext)
		w.Header().Set("Content-Type", ct)
	}
	fmt.Fprint(w, string(data))
}

// RunServer runs a server at (TODO) address
func RunServer() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/static/", staticHandler)
	log.Fatal(http.ListenAndServe(":9321", nil))
}
