package service

import (
	"github.com/ahmadirfaan/project-go/models/database"
	"github.com/ahmadirfaan/project-go/repositories"
	"github.com/gomodule/redigo/redis"
	log "github.com/sirupsen/logrus"
)

type LocationService interface {
	GetAllLocationProvince() ([]database.Provinces, error)
	GetAllRegencyByProvince(provinceId string) ([]database.Regencies, error)
	GetAllDistrictByRegency(regencyId string) ([]database.Districts, error)
	FindDistrictById(districtId string) (database.Districts, error)
}

type locationService struct {
	ProvinceRepository repositories.ProvinceRepository
	RegencyRepository  repositories.RegencyRepository
	DistrictRepository repositories.DistrictRepository
	RedisClient        redis.Conn
}

func NewLocationService(pr repositories.ProvinceRepository, rr repositories.RegencyRepository, dr repositories.DistrictRepository, rc redis.Conn) LocationService {
	return &locationService{
		ProvinceRepository: pr,
		RegencyRepository:  rr,
		DistrictRepository: dr,
		RedisClient:        rc,
	}
}

const (
	TIME_EXPIRED_KEY               = 10800
	KEY_ALL_PROVINCE               = "KEY_ALL_PROVINCE"
	KEY_ALL_REGENCY_BY_PROVINCE_ID = "KEY_ALL_REGENCY_BY_PROVINCE_ID"
	KEY_ALL_DISTRICT_BY_REGENCY_ID = "KEY_ALL_DISTRICT_BY_REGENCY_ID"
	KEY_DISCTRICT_ID               = "KEY_DISCTRICT_ID"
)

func (l *locationService) GetAllLocationProvince() ([]database.Provinces, error) {
	var provinces []database.Provinces
	data, err := redis.Values(l.RedisClient.Do("HGETALL", KEY_ALL_PROVINCE))
	if err != nil || data != nil {
		provinces, err = l.ProvinceRepository.GetAll()
		_, err2 := l.RedisClient.Do("HSET", KEY_ALL_PROVINCE, provinces)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", KEY_ALL_PROVINCE, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		err = redis.ScanStruct(data, &provinces)
		log.Info("Check data: %v ", data)
	}

	return provinces, err
}

func (l *locationService) GetAllRegencyByProvince(provinceId string) ([]database.Regencies, error) {
	var regencies []database.Regencies
	keyRedis := KEY_ALL_REGENCY_BY_PROVINCE_ID + "_" + provinceId
	data, err := redis.Values(l.RedisClient.Do("HGETALL", keyRedis))
	if err != nil || data != nil {
		regencies, err = l.RegencyRepository.FindByProvinceId(provinceId)
		_, err2 := l.RedisClient.Do("HSET", keyRedis, regencies)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", keyRedis, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		err = redis.ScanStruct(data, &regencies)
		log.Info("Check data: %v ", data)
	}
	err = nil
	return regencies, err
}

func (l *locationService) GetAllDistrictByRegency(regencyId string) ([]database.Districts, error) {
	var districts []database.Districts
	keyRedis := KEY_ALL_DISTRICT_BY_REGENCY_ID + "_" + regencyId
	data, err := redis.Values(l.RedisClient.Do("HGETALL", keyRedis))
	if err != nil || data != nil {
		districts, err = l.DistrictRepository.FindByRegencyId(regencyId)
		_, err2 := l.RedisClient.Do("HSET", keyRedis, districts)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", keyRedis, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		err = redis.ScanStruct(data, &districts)
		log.Info("Check data: %v ", data)
	}
	return districts, err
}

func (l *locationService) FindDistrictById(districtId string) (database.Districts, error) {
	var district database.Districts
	keyRedis := KEY_DISCTRICT_ID + "_" + districtId
	data, err := redis.Values(l.RedisClient.Do("HGETALL", keyRedis))
	if err != nil || data != nil {
		district, err = l.DistrictRepository.FindById(districtId)
		_, err2 := l.RedisClient.Do("HSET", keyRedis, district)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", keyRedis, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		err = redis.ScanStruct(data, &district)
		log.Info("Check data: %v ", data)
	}
	return district, err
}
