package database

func (db *Database) ScanDomains(cursor uint64, count int64) ([]string, uint64, error) {
	return db.redisClient.SScan(*db.ctx, allHosts, cursor, "", count).Result()
}
