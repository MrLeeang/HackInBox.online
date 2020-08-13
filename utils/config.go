package utils

import "fmt"

const (
	DbType = "mysql"
	DbHost = "118.31.237.114"
	DbPort = "3307"
	DbUser = "root"
	DbPass = "123456"
	DbName = "HackInBox"

	RedisHost     = "118.31.237.114"
	RedisPort     = "6379"
	RedisPassword = "ccadmin1QAZ"
)

var SessionRedisAddress = fmt.Sprintf("%s:%s", RedisHost, RedisPort)
