package service

import (
	"github.com/ahmadirfaan/project-go/models/database"
	"github.com/ahmadirfaan/project-go/repositories"
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
}

func NewLocationService(pr repositories.ProvinceRepository, rr repositories.RegencyRepository, dr repositories.DistrictRepository) LocationService {
	return &locationService{
		ProvinceRepository: pr,
		RegencyRepository:  rr,
		DistrictRepository: dr,
	}
}

func (l *locationService) GetAllLocationProvince() ([]database.Provinces, error) {
	provinces, err := l.ProvinceRepository.GetAll()
	return provinces, err
}

func (l *locationService) GetAllRegencyByProvince(provinceId string) ([]database.Regencies, error) {
	regencies, err := l.RegencyRepository.FindByProvinceId(provinceId)
	return regencies, err
}

func (l *locationService) GetAllDistrictByRegency(regencyId string) ([]database.Districts, error) {
	districts, err := l.DistrictRepository.FindByRegencyId(regencyId)
	return districts, err
}

func (l *locationService) FindDistrictById(districtId string) (database.Districts, error) {
	district, err := l.DistrictRepository.FindById(districtId)
	return district, err
}
