package main

import (
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

	go observer(ch)
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
		Addr:           ":8071",
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	//cores := runtime.NumCPU() - 4

	ChatHandler(ch)

	//for i := 1; i < cores; i++ {
	//	go internalSend(ch)
	//}

	log.Fatal(s.ListenAndServe())

}

type Chat struct {
	Message string
	User    string
	Ip      string
}
