package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	//"os"
	//"runtime"
	"time"
)

//
//var addr = flag.String("addr", "172.22.0.2:8082", "http service address fpr ws")
//
//var upgrader = websocket.Upgrader{} // use default options
//
//func echo(w http.ResponseWriter, r *http.Request) {
//	c, err := upgrader.Upgrade(w, r, nil)
//	if err != nil {
//		log.Print("upgrade:", err)
//		return
//	}
//	defer c.Close()
//	for {
//		mt, message, err := c.ReadMessage()
//		if err != nil {
//			log.Println("read:", err)
//			break
//		}
//		log.Printf("recv: %s", message)
//		err = c.WriteMessage(mt, message)
//		if err != nil {
//			log.Println("write:", err)
//			break
//		}
//	}
//}

//func home(w http.ResponseWriter, r *http.Request) {
//	homeTemplate.Execute(w, "ws://"+r.Host+"/echo")
//}

func main() {

	log.Print("hi!")

	ch := make(chan Chat, 100)

	msg := Chat{Message: "ttt", User: "militska", Ip: "11"}

	rdb := getRedisClient()
	err := rdb.Set("key2", &msg, 0).Err()
	if err != nil {
		panic(err)
	}

	//flag.Parse()
	//log.SetFlags(0)
	//http.HandleFunc("/echo", echo)
	//log.Fatal(http.ListenAndServe(*addr, nil))
	//
	////go observer(ch)
	initHttpServer(ch)
}

// @todo send to queue
func observer(ch chan Chat) {
	for {
		message := <-ch
		fmt.Println("[sender: " + message.User + "] text: " + message.Message)
	}

}
func initHttpServer(ch chan Chat) {

	s := &http.Server{
		Addr:           ":8074",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go setToRedis(ch)
	ChatHandler(ch)
	log.Fatal(s.ListenAndServe())

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
