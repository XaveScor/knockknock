package database

import (
	"context"
	"github.com/redis/go-redis/v9"
)

type Database struct {
	ctx         *context.Context
	redisClient *redis.Client
}

func Init(redisAddr string) *Database {
	ctx := context.Background()
	redisClient := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	initAllHosts(&ctx, redisClient)
	initBannedHosts(&ctx, redisClient)

	return &Database{
		ctx:         &ctx,
		redisClient: redisClient,
	}
}
