package redisconn

import (
	"github.com/ahmadirfaan/project-go/app"
	"github.com/gomodule/redigo/redis"
)

func InitRedisClient() redis.Conn {
	app := app.Init()
	pool := redis.NewPool(
		func() (redis.Conn, error) {
			return redis.Dial("tcp", app.Config.RedisAddress)
		},
		10,
	)

	pool.MaxActive = 10
	conn := pool.Get()
	defer conn.Close()
	return conn
}
