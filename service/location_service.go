package service

import (
	"encoding/json"
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
	TIME_EXPIRED_KEY               = 604800
	KEY_ALL_PROVINCE               = "KEY_ALL_PROVINCE"
	KEY_ALL_REGENCY_BY_PROVINCE_ID = "KEY_ALL_REGENCY_BY_PROVINCE_ID"
	KEY_ALL_DISTRICT_BY_REGENCY_ID = "KEY_ALL_DISTRICT_BY_REGENCY_ID"
	KEY_DISCTRICT_ID               = "KEY_DISCTRICT_ID"
)

func (l *locationService) GetAllLocationProvince() ([]database.Provinces, error) {
	var provinces []database.Provinces
	data, err := l.RedisClient.Do("GET", KEY_ALL_PROVINCE)
	if err != nil || data == nil {
		provinces, err = l.ProvinceRepository.GetAll()
		marshal, _ := json.Marshal(provinces)
		_, err2 := l.RedisClient.Do("SET", KEY_ALL_PROVINCE, marshal)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", KEY_ALL_PROVINCE, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		bytes := []byte(convertToString(data))
		err := json.Unmarshal(bytes, &provinces)
		if err != nil {
			return nil, err
		}
		log.Info("Check data: %v ", provinces)
	}

	return provinces, err
}

func (l *locationService) GetAllRegencyByProvince(provinceId string) ([]database.Regencies, error) {
	var regencies []database.Regencies
	keyRedis := KEY_ALL_REGENCY_BY_PROVINCE_ID + "_" + provinceId
	data, err := l.RedisClient.Do("GET", keyRedis)
	if err != nil || data == nil {
		regencies, err = l.RegencyRepository.FindByProvinceId(provinceId)
		marshal, _ := json.Marshal(regencies)
		_, err2 := l.RedisClient.Do("SET", keyRedis, marshal)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", keyRedis, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		bytes := []byte(convertToString(data))
		err := json.Unmarshal(bytes, &regencies)
		if err != nil {
			return nil, err
		}
		log.Info("Check data: %v ", regencies)
	}
	err = nil
	return regencies, err
}

func (l *locationService) GetAllDistrictByRegency(regencyId string) ([]database.Districts, error) {
	var districts []database.Districts
	keyRedis := KEY_ALL_DISTRICT_BY_REGENCY_ID + "_" + regencyId
	data, err := l.RedisClient.Do("GET", keyRedis)
	if err != nil || data == nil {
		districts, err = l.DistrictRepository.FindByRegencyId(regencyId)
		marshal, _ := json.Marshal(districts)
		_, err2 := l.RedisClient.Do("SET", keyRedis, marshal)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", keyRedis, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		bytes := []byte(convertToString(data))
		err := json.Unmarshal(bytes, &districts)
		if err != nil {
			return nil, err
		}
		log.Info("Check data: %v ", districts)
	}
	return districts, err
}

func (l *locationService) FindDistrictById(districtId string) (database.Districts, error) {
	var district database.Districts
	keyRedis := KEY_DISCTRICT_ID + "_" + districtId
	data, err := l.RedisClient.Do("GET", keyRedis)
	if err != nil || data == nil {
		district, err = l.DistrictRepository.FindById(districtId)
		marshal, _ := json.Marshal(district)
		_, err2 := l.RedisClient.Do("SET", keyRedis, marshal)
		if err2 != nil {
			log.Error("Error connection to redis store data : %v ", data)
		}
		_, err2 = l.RedisClient.Do("EXPIRE", keyRedis, TIME_EXPIRED_KEY)
		if err2 != nil {
			log.Error("Error set expired key to redis : %v ", data)
		}
	} else {
		bytes := []byte(convertToString(data))
		err := json.Unmarshal(bytes, &district)
		if err != nil {
			return database.Districts{}, err
		}
		log.Info("Check data: %v ", district)
	}
	return district, err
}

func convertToString(bs interface{}) string {
	ba := []byte{}
	for _, b := range bs.([]uint8) {
		ba = append(ba, byte(b))
	}
	return string(ba)
}
