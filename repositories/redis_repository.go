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
	Redis redis.Conn
}

func NewRedisRepository(redis redis.Conn) RedisRepository {
	return redisRepository{
		Redis: redis,
	}
}

func (r redisRepository) SetDataToRedis(data interface{}, key string, expire int) error {
	marshal, _ := json.Marshal(data)
	_, err := r.Redis.Do("SET", key, marshal)
	if err != nil {
		log.Error("Error connection to redis store data : %v ", data)
		return err
	}
	_, err = r.Redis.Do("EXPIRE", key, expire)
	if err != nil {
		log.Error("Error set expired key to redis : %v ", data)
		return err
	}
	log.Info("Sucess put data to redis")
	return nil
}

func (r redisRepository) GetDataFromRedis(key string) ([]byte, error) {
	data, err := r.Redis.Do("GET", key)
	if err != nil || data == nil {
		return nil, err
	}
	bytes := []byte(convertToString(data))
	return bytes, nil
}

func convertToString(bs interface{}) string {
	ba := []byte{}
	for _, b := range bs.([]uint8) {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
