package main

import (
	"encoding/json"
	"flag"
	"fmt"
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

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func main() {
	log.Print("hi!")
	ch := make(chan Chat, 100)

	fmt.Println("IPv4: ", getIp())

	go observer(ch)
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/echo", echo)
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

// @todo send to queue
func observer(ch chan Chat) {
	for {
		message := <-ch

		//msg := Chat{Message: "ttt", User: "militska", Ip: "11"}
		//
		//rdb := getRedisClient()
		//err := rdb.Set("key2", &msg, 0).Err()
		//if err != nil {
		//	panic(err)
		//}

		fmt.Println("[sender: " + message.User + "] text: " + message.Message)
	}
}

func ContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Our middleware logic goes here...
		next.ServeHTTP(w, r)
	})
}

type Chat struct {
	Message string
	User    string
	Ip      string
}

func (m *Chat) MarshalBinary() ([]byte, error) {
	return json.Marshal(m)
}

func setToRedis(ch chan Chat) {
	for {
		message := <-ch

		rdb := getRedisClient()
		err := rdb.Set(message.User, &message, 0).Err()

		if err != nil {
			fmt.Print(err)
			panic(err)
		}

		fmt.Print(rdb.Get("militska"))
	}
}
