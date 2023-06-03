package common

import "os"

type Env struct {
	RedisAddr string
}

func GetEnvs() *Env {
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		panic("REDIS_ADDR is not set")
	}

	return &Env{
		RedisAddr: redisAddr,
	}
}
