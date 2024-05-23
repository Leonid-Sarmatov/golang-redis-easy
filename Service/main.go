package main

import (
	"context"
    "fmt"
    "github.com/redis/go-redis/v9"
)

func main() {
	/*opt, err := redis.ParseURL("redis://leonid:password@redis:6379/10")
	if err != nil {
		panic(err)
	}*/

	client := redis.NewClient(&redis.Options{
        Addr:     "redis:6379",
        Password: "",
        DB:       0,
    })
	
	//client := redis.NewClient(opt)

	// Записать простую строку
	ctx := context.Background()
	key := "MyKey"
	value := "MyValue"
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		panic(err)
	}
	
	// Извлечь строку
	res, err := client.Get(ctx, key).Result()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Key: %v, Value: %v\n", key, res)

	for {
		
	}
}