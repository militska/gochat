package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var (
	addr      = flag.String("addr", getIp().String()+":8074", "http service address")
	clients   = make(map[*websocket.Conn]bool)
	broadcast = make(chan Message)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

type Chat struct {
	Message string
	Name    string
	//Ip      string
}

/**
* send message to all connected clients
 */
func broadcastSender() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				err = client.Close()
				if err != nil {
					log.Println("broadcastSender. error close client:", err)

				}
				delete(clients, client)
			}
		}
	}
}

func main() {
	log.Print("hi! my ipv4 " + getIp().String())
	go broadcastSender()
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/chat", chat)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func (m *Chat) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
