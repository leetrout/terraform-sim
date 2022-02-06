package main

import (
	"fmt"

	"github.com/leetrout/terraform-sim/internal/store"
	"github.com/leetrout/terraform-sim/internal/webserver"
)

func main() {
	fmt.Println("Starting up...")
	store.Initialize()
	webserver.RunServer()
}
