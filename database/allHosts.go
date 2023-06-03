package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func initAllHosts(ctx *context.Context, redisClient *redis.Client) {
	result, err := redisClient.Type(*ctx, allHosts).Result()
	if err != nil {
		log.Fatalln("[FATAL]InitRedis: redisClient.Type(ctx, allHosts).Result() error", err)
	}
	if result != "set" {
		log.Printf("[WARN]InitRedis: allHosts(%s) is not set, rewriting it\n", allHosts)
		err := redisClient.Del(*ctx, allHosts).Err()
		if err != nil {
			log.Fatalf("[FATAL]InitRedis: cannot remove the allHosts(%s) key.\n%s\n", allHosts, err)
		}
		log.Printf("[INFO]InitRedis: the key allHosts(%s) removed\n", allHosts)
		addRes, err := redisClient.SAdd(*ctx, allHosts, "").Result()
		if err != nil || addRes != 1 {
			log.Fatalf("[FATAL]InitRedis: cannot add the allHosts(%s) key.\n%s\n", allHosts, err)
		}
		log.Printf("[INFO]InitRedis: the SET key allHosts(%s) created\n", allHosts)
	} else {
		log.Printf("[INFO]InitRedis: the SET key allHosts(%s) exists\n", allHosts)
	}
}
