package db

import (
    "github.com/go-redis/redis"
)

var redisDB  *redis.Client

func useRedis() *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     "redis-16295.c302.asia-northeast1-1.gce.cloud.redislabs.com:16295",
        Password: "uIHhmDDsYQJnS5y5Hdkr866A1bDA5Ois", // Redis 无密码设置的情况下，此处应该为空字符串
        DB:       0,  // 使用默认的 DB
    })
	return client
}

func GetRedis() *redis.Client {
	if redisDB == nil {
		redisDB = useRedis()
	}
	return redisDB
}