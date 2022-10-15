package repositories

import (
	"errors"
	"fmt"
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type UserRepository interface {
	Save(user database.User) (database.User, error)
	CheckUsernameAndPassword(username string, roleId uint) (database.User, error)
	IsUserAgent(userId uint) (*bool, error)
	IsExist(userId uint) (bool, error)
	FindByUserId(userId string) (database.User, error)
	FindAgentByDistrictId(districtId string) ([]database.User, error)
	WithTrx(trxHandle *gorm.DB) userRepository
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		DB: db,
	}
}

func (u userRepository) Save(user database.User) (database.User, error) {
	err := u.DB.Debug().Save(&user).Error
	fmt.Println(err)
	log.Printf("Users Repositories:%+v\n", user)
	return user, err
}

func (u userRepository) CheckUsernameAndPassword(username string, roleId uint) (database.User, error) {
	var user database.User
	err := u.DB.Debug().Where("username = ? AND role_id = ?", username, roleId).Preload("Role").First(&user).Error
	if user.Id == nil {
		return user, errors.New("No matched user in the database")
	}
	log.Printf("Login Repositories:%+v\n", user)
	return user, err
}

func (u userRepository) IsUserAgent(userId uint) (*bool, error) {
	var user database.User
	err := u.DB.Debug().Where("id = ? AND role_id = ?", userId, 1).Preload("Role").First(&user).Error
	isAgent := true
	isCustomer := false
	if user.Role.Id == 1 {
		return &isAgent, err
	} else {
		return &isCustomer, err
	}
}

func (u userRepository) IsExist(userId uint) (bool, error) {
	var user database.User
	err := u.DB.Debug().Where("id = ?", userId).First(&user).Error
	if user.Id != nil {
		return true, err
	} else {
		return false, err
	}
}
func (u userRepository) FindByUserId(userId string) (database.User, error) {
	var user database.User
	err := u.DB.Debug().Where("id = ?", userId).First(&user).Error
	return user, err
}

func (u userRepository) FindAgentByDistrictId(districtId string) ([]database.User, error) {
	var users []database.User
	err := u.DB.Debug().Joins("Agent").Where("agent_id is NOT NULL").
		Where("Agent.district_id = ?", districtId).Find(&users).Error
	return users, err
}

func (u userRepository) WithTrx(trxHandle *gorm.DB) userRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return u
	}
	u.DB = trxHandle
	return u
}
