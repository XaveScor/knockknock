package database

import (
	"context"
	"github.com/redis/go-redis/v9"
	"knockknocker/common"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func writeDefaultDataIntoAllHosts(ctx *context.Context, redisClient *redis.Client) (int64, error) {
	ownDir := common.GetEnvs().OwnDirectory
	hostsFilename := filepath.Join(ownDir, "hosts.txt")
	hostsFile, err := os.OpenFile(hostsFilename, os.O_RDONLY, 0666)
	if err != nil {
		log.Println("[WARN]InitRedis: cannot open the hosts.txt file, skipping it")
		return redisClient.SAdd(*ctx, allHosts, "").Result()
	}
	defer func(hostsFile *os.File) {
		err := hostsFile.Close()
		if err != nil {
			log.Println("[WARN]InitRedis: cannot close the hosts.txt file")
		}
	}(hostsFile)

	hostsFileContent, err := os.ReadFile(hostsFilename)
	domains := strings.Split(string(hostsFileContent), "\n")
	addRes, err := redisClient.SAdd(*ctx, allHosts, domains).Result()
	if err != nil {
		return 0, err
	}
	log.Println("[INFO]InitRedis: the data from hosts.txt copied to the allHosts key")
	return addRes, nil
}

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
		addRes, err := writeDefaultDataIntoAllHosts(ctx, redisClient)
		if err != nil || addRes == 0 {
			log.Fatalf("[FATAL]InitRedis: cannot add the allHosts(%s) key.\n%s\n", allHosts, err)
		}
		log.Printf("[INFO]InitRedis: the SET key allHosts(%s) created\n", allHosts)
	} else {
		log.Printf("[INFO]InitRedis: the SET key allHosts(%s) exists\n", allHosts)
	}
}

func validateAllHostsQueue(ctx *context.Context, redisClient *redis.Client) {
	result, err := redisClient.Type(*ctx, allHostsQueue).Result()
	if err != nil {
		log.Fatalf("[FATAL]InitRedis: redisClient.Type(ctx, allHostsQueue).Result() error\n%s\n", err)
	}
	if result != "set" {
		log.Printf("[WARN]InitRedis: allHostsQueue(%s) is not set, removing it\n", allHostsQueue)
		err := redisClient.Del(*ctx, allHostsQueue).Err()
		if err != nil {
			log.Fatalf("[FATAL]InitRedis: cannot remove the allHostsQueue(%s) key.\n%s\n", allHostsQueue, err)
		}
	}
}
