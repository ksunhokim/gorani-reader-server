package dbs

const DBName = "engbreaker"

func Init() {
	initMongo()
	initRedis()
}
