package database

import (
	"context"
	"fmt"
	"spotify-clone/helpers"

	"github.com/go-redis/redis/v8"
)

const (
	Addr = "localhost:6379" 
	PasswordRedis= ""
	DB = 0 
)

func ConnectRedis() *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr:  Addr,
		Password: PasswordRedis,
		DB: DB,
	})

	ctx := context.Background()

	if _, err := rdb.Ping(ctx).Result(); err != nil{
		helpers.CheckPanic(err)
	}

	fmt.Println("redis connections successfull")
	return rdb;
}
