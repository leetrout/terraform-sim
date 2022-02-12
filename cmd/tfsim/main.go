package main

import (
	"fmt"

	"github.com/leetrout/terraform-sim/internal/store"
	"github.com/leetrout/terraform-sim/internal/webserver"
)

func main() {
	addr := ":9321"
	fmt.Printf("API server starting at %s\n", addr)
	store.Initialize()
	webserver.RunServer(addr)
}
