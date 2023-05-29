package main

import (
	"context"
	"github.com/redis/go-redis/v9"
	"knockknocker/requester"
	"os"
	"strconv"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	var cursor uint64
	count := 0
	for {
		var keys []string
		keys, cursor, _ = rdb.SScan(ctx, "all-hosts", cursor, "", 10).Result()
		if cursor == 0 {
			break
		}
		for _, dirtyUrl := range keys {
			count++
			err := requester.TouchWebsite(dirtyUrl)
			if err != nil {
				rdb.SAdd(ctx, "banned", dirtyUrl)
				println(strconv.Itoa(count) + "|" + dirtyUrl + " is banned")
			} else {
				println(strconv.Itoa(count) + "|" + dirtyUrl + " is not banned")
			}
		}
	}

	println("done")
}
