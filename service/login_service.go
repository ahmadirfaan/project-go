package service

import (
	"errors"
	"github.com/ahmadirfaan/project-go/models/database"
	"github.com/ahmadirfaan/project-go/models/web"
	"github.com/ahmadirfaan/project-go/repositories"
	"github.com/ahmadirfaan/project-go/utils"
	"log"
)

type LoginService interface {
	Login(request web.LoginRequest) (database.User, error)
}

type loginService struct {
	UserRepository repositories.UserRepository
}

func NewLoginService(ur repositories.UserRepository) LoginService {
	return &loginService{
		UserRepository: ur,
	}
}

func (l *loginService) Login(request web.LoginRequest) (database.User, error) {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return database.User{}, err
	}
	log.Println("Request Password: ", request.Password)
	user, err := l.UserRepository.CheckUsernameAndPassword(request.Username, request.Role)
	checkPassword := utils.CheckPasswordHash(request.Password, user.Password)
	log.Println("Hash Password Database: ", user.Password)
	log.Println("Check Password in Database: ", checkPassword)
	if !checkPassword {
		return database.User{}, errors.New("There is no match record in our database")
	}
	return user, err
}
