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

	var cursor uint64
	for {
		var keys []string
		keys, cursor, _ = db.ScanDomains(cursor, 100)
		if cursor == 0 {
			break
		}
		for _, dirtyUrl := range keys {
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

	log.Printf("finished")
}
