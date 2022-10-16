package redisconn

import (
	"github.com/ahmadirfaan/project-go/app"
	"github.com/gomodule/redigo/redis"
)

func InitRedisClient() *redis.Pool {
	app := app.Init()
	return &redis.Pool{
		MaxIdle:   80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", app.Config.RedisAddress)
			if err != nil {
				panic(err.Error())
			}
			return c, err
		},
	}

}
