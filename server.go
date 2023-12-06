package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var (
	clients  []websocket.Conn
	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(
			w, r, nil)

		clients = append(clients, *conn)

		for {
			msgType, byte, err := conn.ReadMessage()
			if err != nil {
				return
			}

			fmt.Printf("send by: %s\n%s", conn.RemoteAddr(), string(byte))
			for _, client := range clients {
				if err = client.WriteMessage(msgType, byte); err != nil {
					return
				}
			}
		}
	})
	http.ListenAndServe(":8090", nil)
}
