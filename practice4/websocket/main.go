package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var clients = make(map[*websocket.Conn]bool)

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)
		clients[conn] = true

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s sent : %s and messageType : %v\n", conn.RemoteAddr(), string(msg), msgType)

			// for returning the message to the same user connected
			// if err = conn.WriteMessage(msgType, msg); err != nil {
			// 	return
			// }

			// Broadcast the message to all connected clients
			for client := range clients {
				fmt.Println("client", client.RemoteAddr())
				if err = client.WriteMessage(msgType, msg); err != nil {
					// If there's an error sending to a client, remove them from the clients map
					delete(clients, client)
					client.Close()
				}
			}
		}

	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})
	http.ListenAndServe(":8080", nil)
}
