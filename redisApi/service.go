package redisApi

import (
	"context"
	"strconv"
	"time"

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

func Set(key string, value string) error {
	return client.Set(context.Background(), key, value, 600*time.Second).Err()
}

func Get(key string) (string, error) {
	return client.Get(context.Background(), key).Result()
}

func Exists(key string) (bool, error) {
	_, err := client.Get(context.Background(), key).Result()
	if err == redis.Nil {
		return false, nil
	} else if err == nil {
		return true, nil
	} else {
		return false, err
	}
}
