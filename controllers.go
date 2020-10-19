package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func ChatHandler(ch chan Chat) {
	http.HandleFunc("/chat", func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			_, _ = w.Write([]byte("Content-Type must be application/json"))
			w.WriteHeader(400)
			return
		}

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(400)
			return
		}
		var data Chat
		if err = json.Unmarshal(body, &data); err != nil {
			w.WriteHeader(400)
			return
		}

		ch <- data

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})

	http.HandleFunc("/chat/history", func(w http.ResponseWriter, r *http.Request) {

		rdb := getRedisClient()

		keys := rdb.Keys("")

		fmt.Print(keys)
		fmt.Print(rdb.Get("militska"))
		//err := rdb.Set("key", &data, 0).Err()
		//if err != nil {
		//	panic(err)
		//}
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("ok"))
	})
}
