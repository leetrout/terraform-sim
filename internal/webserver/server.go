package webserver

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leetrout/terraform-sim/internal/ws"
)

func handler(w http.ResponseWriter, r *http.Request) {
	var html, err = f.ReadFile("static/index.html")
	if err != nil {
		// Fixme
		return
	}
	fmt.Fprintf(w, string(html))
}

// func staticHandler(w http.ResponseWriter, r *http.Request) {
// 	fs := http.FileServer(http.FS(f))
// }

// RunServer runs a server at (TODO) address
func RunServer(addr string) {
	// Setup websocket exchange
	ws.Init()
	http.Handle("/static/", http.FileServer(http.FS(f)))
	http.HandleFunc("/", handler)
	http.HandleFunc("/api/", apiHandler)
	http.HandleFunc("/ws", ws.SocketHandler)
	log.Fatal(http.ListenAndServe(addr, nil))
}
