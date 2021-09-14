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
	broadcast = make(chan Message, 100)
	redis     = make(chan Message, 100)
	upgrader  = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

type Message struct {
	Username string
	Message  string
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

func redisHandler() {
	for recForRedis := range redis {
		setMsg(recForRedis)
	}
}

func main() {
	log.Print("hi! my ipv4 " + getIp().String())
	go broadcastSender()
	go redisHandler()
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/chat", chat)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func (m *Message) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
