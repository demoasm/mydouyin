package cache

import "github.com/go-redis/redis/v8"

var rdeisClient *redis.Client

func Init() {
	rdeisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	initVideoCache()
}
