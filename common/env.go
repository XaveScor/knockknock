package common

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	RedisAddr string
}

func GetEnvs() *Env {
	_ = godotenv.Load()
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		panic("env REDIS_ADDR is not set")
	}

	return &Env{
		RedisAddr: redisAddr,
	}
}
