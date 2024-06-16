package redisApi

import (
	"context"
	"strconv"

	"github.com/go-redis/redis/v8"
)

var client *redis.Client

func Init(host string, port int, password string, db int) error {
	client = redis.NewClient(&redis.Options{
		Addr:     host + ":" + strconv.Itoa(port),
		Password: password,
		DB:       db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		return err
	}

	return nil
}
