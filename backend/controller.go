package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func chat(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)

	log.Print("chat")
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

		var msg Message
		err2 := json.Unmarshal(data, &msg)

		if err2 != nil {
			log.Println("Unmarshal error:", err2)
			break

		}

		log.Print("username " + msg.Username)
		log.Print("message " + msg.Message)

		toRedis <- msg
		broadcast <- msg
		ownmessage := []byte(" hi! :)  from server " + string(message))
		err = c.WriteMessage(mt, ownmessage)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

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

		msg := Message{Message: "from server " + string(message), Username: "militska"}
		broadcast <- msg
		ownmessage := []byte(" hi! 71 :)  from server " + string(message))
		err = c.WriteMessage(mt, ownmessage)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}
