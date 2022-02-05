package main

import (
	"fmt"

	"github.com/leetrout/terraform-sim/internal/webserver"
)

func main() {
	fmt.Println("Hello world.")
	webserver.RunServer()
}
