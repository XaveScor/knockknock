package main

import (
	"knockknocker/common"
	"knockknocker/database"
	"knockknocker/requester"
	"log"
)

func main() {
	db := database.Init(common.GetEnvs().RedisAddr)
	defer db.Close()

	for {
		for _, dirtyUrl := range db.GetNextDomains(10) {
			log.Printf("[INFO]main: checking %s\n", dirtyUrl)
			err := requester.TouchWebsite(dirtyUrl)
			if err != nil {
				res, err := db.SaveBannedHost(dirtyUrl)
				if !res || err != nil {
					log.Println("[ERROR]main: cannot save banned host", dirtyUrl, err)
				} else {
					log.Printf("[INFO]main: %s is banned\n", dirtyUrl)
				}
			} else {
				log.Printf("[INFO]main: %s is alive\n", dirtyUrl)
			}
		}
	}
}
