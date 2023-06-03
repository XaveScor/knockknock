package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
)

func initBannedHosts(ctx *context.Context, redisClient *redis.Client) {
	result, err := redisClient.Type(*ctx, bannedHosts).Result()
	if err != nil {
		log.Fatalln("[FATAL]InitRedis: redisClient.Type(ctx, bannedHosts).Result() error", err)
	}
	if result != "set" {
		log.Printf("[WARN]InitRedis: bannedHosts(%s) is not set, rewriting it\n", bannedHosts)
		err := redisClient.Del(*ctx, bannedHosts).Err()
		if err != nil {
			log.Fatalf("[FATAL]InitRedis: cannot remove the bannedHosts(%s) key.\n%s\n", bannedHosts, err)
		}
		log.Printf("[INFO]InitRedis: the key bannedHosts(%s) removed\n", bannedHosts)
		addRes, err := redisClient.SAdd(*ctx, bannedHosts, "").Result()
		if err != nil || addRes != 1 {
			log.Fatalf("[FATAL]InitRedis: cannot add the bannedHosts(%s) key.\n%s\n", bannedHosts, err)
		}
		log.Printf("[INFO]InitRedis: the SET key bannedHosts(%s) created\n", bannedHosts)
	} else {
		log.Printf("[INFO]InitRedis: the SET key bannedHosts(%s) exists\n", bannedHosts)
	}
}

func (db *Database) SaveBannedHost(host string) (bool, error) {
	res, err := db.redisClient.SAdd(*db.ctx, bannedHosts, host).Result()
	return res == 1, err
}
