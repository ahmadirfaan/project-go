package repositories

import (
	"fmt"
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type RegencyRepository interface {
	FindByProvinceId(provinceId string) ([]database.Regencies, error)
	GetAll() ([]database.Regencies, error)
}

type regencyRepository struct {
	DB *gorm.DB
}

func NewRegencyRepository(db *gorm.DB) RegencyRepository {
	return &regencyRepository{
		DB: db,
	}
}

func (u regencyRepository) FindByProvinceId(provinceId string) ([]database.Regencies, error) {
	log.Println("Ini Error sebelum di save")
	var regency []database.Regencies
	err := u.DB.Debug().Where("province_id = ?", provinceId).Find(&regency).Error
	fmt.Println(err)
	log.Printf("Regency Repositories:%+v\n", regency)
	return regency, err
}

func (u regencyRepository) GetAll() ([]database.Regencies, error) {
	log.Println("Ini Error sebelum di save")
	var Regency []database.Regencies
	err := u.DB.Debug().Find(&Regency).Error
	fmt.Println(err)
	log.Printf("Province Repositories:%+v\n", Regency)
	return Regency, err
}
