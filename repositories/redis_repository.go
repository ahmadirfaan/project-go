package repositories

import (
	"encoding/json"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

type RedisRepository interface {
	SetDataToRedis(data interface{}, key string, expire int) error
	GetDataFromRedis(key string) ([]byte, error)
}

type redisRepository struct {
	Redis *redis.Pool
}

func NewRedisRepository(redis *redis.Pool) RedisRepository {
	return redisRepository{
		Redis: redis,
	}
}

func (r redisRepository) SetDataToRedis(data interface{}, key string, expire int) error {
	marshal, _ := json.Marshal(data)
	connection := r.Redis.Get()
	defer func(connection redis.Conn) {
		err := connection.Close()
		if err != nil {
			log.Info("Error Connection to redis")
		}
	}(connection)
	_, err := connection.Do("SET", key, marshal)
	if err != nil {
		log.Error("Error connection to redis store data : %v ", data)
		return err
	}
	_, err = connection.Do("EXPIRE", key, expire)
	if err != nil {
		log.Error("Error set expired key to redis : %v ", data)
		return err
	}
	log.Info("Sucess put data to redis")
	return nil
}

func (r redisRepository) GetDataFromRedis(key string) ([]byte, error) {
	connection := r.Redis.Get()
	defer func(connection redis.Conn) {
		err := connection.Close()
		if err != nil {
			log.Info("Error Connection to redis")
		}
	}(connection)
	data, err := connection.Do("GET", key)
	if err != nil || data == nil {
		return nil, err
	}
	bytes := []byte(convertToString(data))
	return bytes, nil
}

func convertToString(bs interface{}) string {
	var ba []byte
	for _, b := range bs.([]uint8) {
		ba = append(ba, b)
	}
	return string(ba)
}
