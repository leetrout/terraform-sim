package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/leetrout/terraform-sim/internal/store"
	"github.com/leetrout/terraform-sim/internal/webserver"
	"github.com/leetrout/terraform-sim/internal/ws"
)

var addrFlag = flag.String("addr", ":9321", "address for webserver to listen")
var interactiveFlag = flag.Bool("i", false, "interactive mode (console)")

func main() {
	flag.Parse()

	addr := ":9321"
	if addrFlag != nil {
		addr = *addrFlag
	}

	store.Initialize()
	if *interactiveFlag == true {
		go console()
	}

	fmt.Printf("+++ API server starting at %s\n", addr)
	webserver.RunServer(addr)
}

func console() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("TerraSim")
	fmt.Println("exit: Close the program")
	fmt.Println("sock: Show websocket info")
	fmt.Println("sock send <string>: Broadcast the string to all WS clients")
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.ReplaceAll(text, "\n", "")

		if strings.Compare("sock", text) == 0 {
			fmt.Printf("%+v\n", ws.SOCKEX)
		}

		if strings.HasPrefix(text, "sock send") {
			ws.SOCKEX.Send(strings.ReplaceAll(text, "sock send ", ""))
		}

		if strings.Compare("exit", text) == 0 {
			os.Exit(0)
		}

	}
}
