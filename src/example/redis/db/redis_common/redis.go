package redis_common

import (
	"context"
	"errors"
	"github.com/redis"
	"log"
)

func GetRedisClient(options redis.Options) (client *redis.Client, e error) {
	client = redis.NewClient(&options)
	result, e := client.Ping(context.Background()).Result()
	if e != nil {
		return nil, e
	}
	if result != "PONG" {
		log.Println("redis ping result is: " + result)
		return nil, errors.New("redis server can not reach; redis ping result: " + result)
	}
	return client, nil
}
