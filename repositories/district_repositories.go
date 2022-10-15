package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type DistrictRepository interface {
	FindByRegencyId(regencyId string) ([]database.Districts, error)
	FindById(districtId string) (database.Districts, error)
}

type districtRepository struct {
	DB *gorm.DB
}

func NewDistrictRepository(db *gorm.DB) DistrictRepository {
	return &districtRepository{
		DB: db,
	}
}

func (d districtRepository) FindByRegencyId(regencyId string) ([]database.Districts, error) {
	var districs []database.Districts
	err := d.DB.Debug().Where("regency_id = ?", regencyId).Find(&districs).Error
	fmt.Println(err)
	log.Printf("District Repositories:%+v\n", districs)
	return districs, err
}

func (d districtRepository) FindById(districtId string) (database.Districts, error) {
	var district database.Districts
	err := d.DB.Debug().First(&district, districtId).Error
	log.Println("District Id : ", district.Id)
	if district.Id != nil {
		return district, err
	} else {
		return district, errors.New("No found record districtId")
	}
}
