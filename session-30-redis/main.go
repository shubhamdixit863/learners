package main

import (
	"github.com/redis/go-redis/v9"
	"log"
	"session-30-redis/internal"
)

func connectToRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return client

}

func main() {
	// redis connection
	client := connectToRedis()
	cache := internal.NewRedisCache(client)
	cache.SetData("logs", "This is a log from the api")
	data, err := cache.GetData("thename")
	if err != nil {
		log.Println("Error", err)
		return
	}
	log.Println("The data", data)
}
