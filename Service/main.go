package main

import (
	"context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

func main() {
	opt, err := redis.ParseURL("redis://<user>:<pass>@localhost:6379/<db>")
	if err != nil {
		panic(err)
	}
	
	client := redis.NewClient(opt)
	
}