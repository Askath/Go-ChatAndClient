package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	connection, _, err := websocket.DefaultDialer.Dial("ws://localhost:8090/echo", nil)
	if err != nil {
		fmt.Printf("Could not connect to Websocket")
		fmt.Printf(err.Error())
		return
	}

	defer connection.Close()

	// go routine for recieving messages regularly
	go func() {
		for {
			_, message, err := connection.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			fmt.Printf("Received: %s\n", message)
		}
	}()

	// Read from stdin and send messages to the server
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		err = connection.WriteMessage(websocket.TextMessage, []byte(text))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}
