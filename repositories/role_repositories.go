package repositories

import (
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type RoleRepository interface {
	Save(role database.Role) (database.Role, error)
}

type roleRepository struct {
	DB *gorm.DB
}

func NewRoleRepository(db *gorm.DB) RoleRepository {
	return roleRepository{
		DB: db,
	}
}

func (r roleRepository) Save(role database.Role) (database.Role, error) {
	err := r.DB.Debug().Create(&role).Error
	log.Printf("Role:%+v\n", role)
	return role, err
}
