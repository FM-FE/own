package main

import (
	"context"
	"example/redis/db/redis_common"
	"github.com/redis"
	"log"
)

func main() {
	options := redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	}
	client, e := redis_common.GetRedisClient(options)
	if e != nil {
		log.Println(">")
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">>>>> ", e.Error())
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">")
		return
	}
	set := client.Set(context.Background(), "key", "value", 0)
	result, e := set.Result()
	if e != nil {
		log.Println(">")
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">>>>> ", e.Error())
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">")
		return
	}
	log.Println("set result is: " + result)
	get := client.Get(context.Background(), "key")
	s, e := get.Result()
	if e != nil {
		log.Println(">")
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">>>>> ", e.Error())
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">")
		return
	}
	log.Println("get result is: " + s)

	get2 := client.Get(context.Background(), "key2")
	s2, e := get.Result()
	if e != nil {
		log.Println(">")
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">>>>> ", e.Error())
		log.Println(">>> ERROR OCCUR !!!")
		log.Println(">")
		return
	}
	log.Println("get result is: " + s2)
}
