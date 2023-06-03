package common

import (
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	RedisAddr    string
	OwnDirectory string
}

func GetEnvs() *Env {
	_ = godotenv.Load()
	redisAddr := os.Getenv("REDIS_ADDR")
	if redisAddr == "" {
		panic("[FATAL]Env: REDIS_ADDR is not set")
	}

	ownDirectory := os.Getenv("OWN_DIRECTORY")
	if dir, err := os.Stat(ownDirectory); err != nil || !dir.IsDir() {
		panic("[FATAL]Env: OWN_DIRECTORY is not set or directory does not exist")
	}

	return &Env{
		RedisAddr:    redisAddr,
		OwnDirectory: ownDirectory,
	}
}
