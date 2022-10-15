package service

import (
	"fmt"

	"github.com/ahmadirfaan/project-go/models/database"
	"github.com/ahmadirfaan/project-go/models/web"
	"github.com/ahmadirfaan/project-go/repositories"
	"github.com/ahmadirfaan/project-go/utils"
	"gorm.io/gorm"
)

type AgentService interface {
	RegisterAgent(request web.RegisterAgentRequest) error
	FindByDistrictId(districtId string) ([]database.User, error)
}

type agentService struct {
	agentRepository repositories.AgentRepository
	userRepository  repositories.UserRepository
	locationService LocationService
	DB              *gorm.DB
}

func NewAgentService(ar repositories.AgentRepository,
	ur repositories.UserRepository, lr LocationService, db *gorm.DB) AgentService {
	return &agentService{
		agentRepository: ar,
		userRepository:  ur,
		locationService: lr,
		DB:              db,
	}
}

func (a *agentService) RegisterAgent(request web.RegisterAgentRequest) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return err
	}
	_, err = a.locationService.FindDistrictById(request.DistrictId)
	if err != nil {
		return err
	}
	tx := a.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	agent := database.Agent{
		AgentName:   request.AgentName,
		NoHandphone: request.NoHandphone,
		Address:     request.Address,
		DistrictId:  request.DistrictId,
	}
	agent, err = a.agentRepository.WithTrx(tx).Save(agent)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	user := database.User{
		RoleId:   1,
		AgentId:  agent.Id,
		Username: request.Username,
		Password: utils.HashPassword(request.Password),
	}
	user, err = a.userRepository.WithTrx(tx).Save(user)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
}

func (a *agentService) FindByDistrictId(districtId string) ([]database.User, error) {
	_, err := a.locationService.FindDistrictById(districtId)
	if err != nil {
		return nil, err
	}
	userAgent, err := a.userRepository.FindAgentByDistrictId(districtId)
	if err != nil {
		return nil, err
	}
	return userAgent, err

}
