package main

import (
	"encoding/json"
	"flag"
	"github.com/gorilla/websocket"
	"log"
	"net"
	"net/http"
	"os"
)

var (
	addr     = flag.String("addr", getIp().String()+":8074", "http service address")
	upgrader = websocket.Upgrader{
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
type Connections struct {
	Con net.Conn
}

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message) // broadcast channel

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	clients[c] = true
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		msg := Message{Message: "from server " + string(message), Username: "militska", Email: "w"}
		broadcast <- msg
		ownmessage := []byte("from server " + string(message))
		err = c.WriteMessage(mt, ownmessage)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func chat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()

	clients[c] = true
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)

		data := []byte(message)

		var ch Chat
		err2 := json.Unmarshal(data, &ch)

		if err2 != nil {
			log.Println("Unmarshal error:", err2)
			break

		}
		log.Print("name:  " + ch.Name)

		msg := Message{Message: "from server " + string(message), Username: "militska", Email: "w"}
		broadcast <- msg
		ownmessage := []byte("from server " + string(message))
		err = c.WriteMessage(mt, ownmessage)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func handleMessages() {
	for {
		msg := <-broadcast
		for client := range clients {
			err := client.WriteJSON(msg)
			if err != nil {
				log.Printf("error: %v", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}

func main() {
	log.Print("hi! my ipv4 " + getIp().String())

	go handleMessages()
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/chat", chat)

	log.Fatal(http.ListenAndServe(*addr, nil))
}

func getIp() net.IP {
	host, _ := os.Hostname()
	addrs, _ := net.LookupIP(host)
	for _, addr := range addrs {
		if ipv4 := addr.To4(); ipv4 != nil {
			return ipv4
		}
	}

	return nil
}

type Chat struct {
	Message string
	Name    string
	//Ip      string
}

func (m *Chat) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}
