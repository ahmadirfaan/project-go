package repositories

import (
	"fmt"
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type ProvinceRepository interface {
	GetAll() ([]database.Provinces, error)
}

type provinceRepository struct {
	DB *gorm.DB
}

func NewProvinceRepository(db *gorm.DB) ProvinceRepository {
	return &provinceRepository{
		DB: db,
	}
}

func (u provinceRepository) GetAll() ([]database.Provinces, error) {
	log.Println("Ini Error sebelum di save")
	var province []database.Provinces
	err := u.DB.Debug().Find(&province).Error
	fmt.Println(err)
	log.Printf("Province Repositories:%+v\n", province)
	return province, err
}
