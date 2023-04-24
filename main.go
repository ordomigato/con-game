package main

import (
	"fmt"
	"log"

	socketio "github.com/googollee/go-socket.io"
	"github.com/ordomigato/con-game/router"
)

func main() {
	server := socketio.NewServer(nil)
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go func() {
		if err := server.Serve(); err != nil {
			log.Fatalf("socketio listen error: %s\n", err)
		}
	}()
	defer server.Close()

	router := router.SetupRouter(server)
	router.Run() // listen and serve on 0.0.0.0:8080
}
