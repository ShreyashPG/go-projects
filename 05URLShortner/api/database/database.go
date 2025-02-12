package database


import ( 
		"context"
		"github.com/go-redis/redis/v8"
		"os"
)


//The context package in Go helps you control how long a function or Goroutine runs. 
//It is mainly used to cancel tasks, set timeouts, 
//and pass values between functions.
var Ctx =context.Background()

func CreateClient(dbNo  int ) *redis.Client{
	rdb := redis.NewClient(&redis.Options{
		Addr: os.Getenv("DB_ADDR"),
		Password : os.Getenv("DB_PASS"),
		DB :dbNo,

	})
	return rdb;
}