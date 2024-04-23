package db

import (
	"os"

	"github.com/go-redis/redis"
)

var redisDB  *redis.Client

func useRedis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     os.Getenv("REDIS_HOST")+":"+os.Getenv("REDIS_PORT"),
        Password: os.Getenv("REDIS_PWD"), 
        DB:       0,  
    })
	return client
}

func GetRedis() *redis.Client {
	if redisDB == nil {
		redisDB = useRedis()
	}
	return redisDB
}