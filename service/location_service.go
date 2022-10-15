package service

import (
	"encoding/json"
	"github.com/ahmadirfaan/project-go/models/database"
	"github.com/ahmadirfaan/project-go/repositories"
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
	RedisClient        repositories.RedisRepository
}

func NewLocationService(pr repositories.ProvinceRepository, rr repositories.RegencyRepository, dr repositories.DistrictRepository, rc repositories.RedisRepository) LocationService {
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
	data, err := l.RedisClient.GetDataFromRedis(KEY_ALL_PROVINCE)
	if err != nil || data == nil {
		provinces, err = l.ProvinceRepository.GetAll()
		err := l.RedisClient.SetDataToRedis(provinces, KEY_ALL_PROVINCE, TIME_EXPIRED_KEY)
		if err != nil {
			return nil, err
		}
	} else {
		err := json.Unmarshal(data, &provinces)
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
	data, err := l.RedisClient.GetDataFromRedis(keyRedis)
	if err != nil || data == nil {
		regencies, err = l.RegencyRepository.FindByProvinceId(provinceId)
		err := l.RedisClient.SetDataToRedis(regencies, keyRedis, TIME_EXPIRED_KEY)
		if err != nil {
			return nil, err
		}
	} else {
		err := json.Unmarshal(data, &regencies)
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
	data, err := l.RedisClient.GetDataFromRedis(keyRedis)
	if err != nil || data == nil {
		districts, err = l.DistrictRepository.FindByRegencyId(regencyId)
		err := l.RedisClient.SetDataToRedis(districts, keyRedis, TIME_EXPIRED_KEY)
		if err != nil {
			return nil, err
		}
	} else {
		err := json.Unmarshal(data, &districts)
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
	data, err := l.RedisClient.GetDataFromRedis(keyRedis)
	if err != nil || data == nil {
		district, err = l.DistrictRepository.FindById(districtId)
		err := l.RedisClient.SetDataToRedis(district, keyRedis, TIME_EXPIRED_KEY)
		if err != nil {
			return database.Districts{}, err
		}
	} else {
		err := json.Unmarshal(data, &district)
		if err != nil {
			return database.Districts{}, err
		}
		log.Info("Check data: %v ", district)
	}
	return district, err
}
