package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

var ctx = context.Background()

func getRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

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

//func ExampleClient() {
//
//	msg := Chat{Message: "ttt", User: "militska", Ip: "11"}
//
//	rdb := getRedisClient();
//	err := rdb.Set("key", &msg, 0).Err()
//	if err != nil {
//		panic(err)
//	}
//
//	val, err := rdb.Get("key").Result()
//	if err != nil {
//		panic(err)
//	}
//	fmt.Println("key", val)
//
//	val2, err := rdb.Get("key2").Result()
//	if err == redis.Nil {
//		fmt.Println("key2 does not exist")
//	} else if err != nil {
//		panic(err)
//	} else {
//		fmt.Println("key2", val2)
//	}
//	// Output: key value
//	// key2 does not exist
//}
