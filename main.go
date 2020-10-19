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

func main() {

	log.Print("hi!")

	ch := make(chan Chat, 100)

	msg := Chat{Message: "ttt", User: "militska", Ip: "11"}

	rdb := getRedisClient()
	err := rdb.Set("key2", &msg, 0).Err()
	if err != nil {
		panic(err)
	}

	//go observer(ch)
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

	http.DefaultTransport.(*http.Transport).MaxIdleConnsPerHost = 100

	go setToRedis(ch)
	ChatHandler(ch)
	log.Fatal(s.ListenAndServe())

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

		fmt.Print(message)

		rdb := getRedisClient()
		err := rdb.Set(message.User, &message, 0).Err()

		if err != nil {
			fmt.Print(err)
			panic(err)
		}

		fmt.Print(rdb.Get("militska"))
	}
}
