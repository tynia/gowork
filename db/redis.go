package db

import (
	"gowork/extern/logging"
	e "gowork/error"
	_ "common/goredis"
)

var (
	Redis *goredis.Redis
)

func InitRedis(url string) *e.WError {
	redis, err := goredis.DialURL(url)
	if err != nil {
		logging.Error("[InitRedis] Connect to REDIS server: %s, error = %s", url, err.Error())
		return e.NewWError(e.ERR_CODE_DB, "Failed to connect to redis: %s", url)
	}

	Redis = redis
	logging.Info("[InitRedis] connect to REDIS server[%s] ok", url)
	return nil
}