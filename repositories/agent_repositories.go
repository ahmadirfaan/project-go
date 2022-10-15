package repositories

import (
	"errors"
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type AgentRepository interface {
	Save(agent database.Agent) (database.Agent, error)
	GiveRatingAgent(agentId uint, agentRating float32, totalReviewCustomer uint) error
	FindByAgentId(agentId uint) (*database.Agent, error)
	WithTrx(trxHandle *gorm.DB) agentRepo
}

type agentRepo struct {
	DB *gorm.DB
}

func NewAgentRepository(db *gorm.DB) AgentRepository {
	return &agentRepo{
		DB: db,
	}
}

func (a agentRepo) Save(agent database.Agent) (database.Agent, error) {
	err := a.DB.Debug().Create(&agent).Error
	log.Printf("Agent:%+v\n", agent)
	return agent, err
}

func (a agentRepo) GiveRatingAgent(agentId uint, agentRating float32, totalReviewCustomer uint) error {
	var agent database.Agent
	err := a.DB.Debug().Model(&agent).Where(" id = ?", agentId).
		Update("agent_rating", agentRating).Update("total_review_customer", totalReviewCustomer).Error
	return err

}

func (a agentRepo) FindByAgentId(agentId uint) (*database.Agent, error) {
	log.Println("Agent Id Di Anjinglah : ")
	var agent *database.Agent
	err := a.DB.Debug().Where("id = ?", agentId).First(&agent).Error
	log.Println("Agent Id Di Repositories : ", agent.Id)
	if agent.Id != nil {
		return agent, err
	} else {
		return agent, errors.New("No found record agentId")
	}

}

func (a agentRepo) WithTrx(trxHandle *gorm.DB) agentRepo {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return a
	}
	a.DB = trxHandle
	return a
}
