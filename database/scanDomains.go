package database

import "log"

func (db *Database) GetNextDomains(count int64) []string {
	result, err := db.redisClient.SPopN(*db.ctx, allHostsQueue, count).Result()
	if err != nil {
		log.Printf("[INFO]GetNextDomains: cannot get next domain. Recreating the allHostsQueue(%s) key\n", allHostsQueue)

		copyRes, err := db.redisClient.Copy(*db.ctx, allHosts, allHostsQueue, 0, true).Result()
		if err != nil || copyRes != 1 {
			log.Fatalf("[FATAL]GetNextDomains: cannot copy the allHosts(%s) key to allHostsQueue(%s) key.\n%s\n", allHosts, allHostsQueue, err)
		}
		return db.GetNextDomains(count)
	}

	return result
}
