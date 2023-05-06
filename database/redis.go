package db

import (
    "github.com/go-redis/redis"
)

var redisDB  *redis.Client

func useRedis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "qlmsmart.com:36379",
        Password: "uIHhmDDsYQJnS5y5Hdkr866A1bDA5Ois", 
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