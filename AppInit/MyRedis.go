package AppInit

import (
	"log"
	"time"

	"github.com/go-redis/redis"
)

var (
	Redis *redis.Client
)

func init() {
	Redis = newRedis()
}

func newRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:         "localhost:6379",
		//Password:     Config.Redis.Password,
		DB:           0,
		PoolSize:     128,
		PoolTimeout:  30 * time.Second,
		IdleTimeout:  50 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		MaxRetries:   3,
	})
	_, err := client.Ping().Result()
	if err != nil {
		log.Fatal(err)
	}
	return client
}
