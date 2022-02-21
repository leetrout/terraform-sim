package ws

import (
	"io"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Init() {
	// Initialize exchange if needed
	if SOCKEX == nil {
		SOCKEX = newSockEx()
		SOCKEX.run()
	}
}

// SocketHandler upgrades connections to a websocket
func SocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	// Initialize exchange if needed
	if SOCKEX == nil {
		SOCKEX = newSockEx()
		SOCKEX.run()
	}

	SOCKEX.addConnection(conn)
}

// SOCKEX is the global socket exchange instance
var SOCKEX *socketExchange

// socketExchange encapsulates the channels to coordinate websocket dispatching
type socketExchange struct {
	x           sync.Mutex
	lastID      int
	connections map[int]*websocket.Conn
	inbound     chan io.Reader
	started     bool
}

func newSockEx() *socketExchange {
	return &socketExchange{
		inbound:     make(chan io.Reader),
		connections: map[int]*websocket.Conn{},
	}
}

func (se *socketExchange) run() {
	// TODO
	if !se.started {
		se.started = true
		se.start()
	}
}

// addConnection stores the connection pointer and starts a NOOP readloop
func (se *socketExchange) addConnection(c *websocket.Conn) {
	se.x.Lock()
	defer se.x.Unlock()
	se.lastID += 1
	se.connections[se.lastID] = c
	go se.readloop(se.lastID, c)

}

func (se *socketExchange) Send(s string) {
	se.x.Lock()
	defer se.x.Unlock()
	for _, conn := range se.connections {
		go func(conn *websocket.Conn) {
			// conn.NextReader()
			w, err := conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write([]byte(s))
			w.Close()
		}(conn)
	}
}

func (se *socketExchange) start() {
	// ticker := time.NewTicker(2 * time.Second)
	// go func() {
	// 	for {
	// 		<-ticker.C
	// 		fmt.Println("=== SOCKEX STATE ===")
	// 		fmt.Printf("%+v\n", se)
	// 		if !se.started {
	// 			ticker.Stop()
	// 			return
	// 		}
	// 	}
	// }()
}

func (se *socketExchange) remove(id int) {
	se.x.Lock()
	delete(se.connections, id)
	se.x.Unlock()
}

func (se *socketExchange) readloop(id int, c *websocket.Conn) {
	for {
		_, _, err := c.NextReader()
		if err != nil {
			c.Close()
			se.remove(id)
			return
		}
	}
}
