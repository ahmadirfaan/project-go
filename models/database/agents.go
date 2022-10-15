package database

import (
	"time"

	"gorm.io/gorm"
)

type Agent struct {
	CreatedAt           time.Time      `json:"-"`
	UpdatedAt           time.Time      `json:"-"`
	DeletedAt           gorm.DeletedAt `gorm:"index" json:"-"`
	Id                  *uint          `gorm:"autoIncrement;primary key" json:"-"`
	AgentName           string         `gorm:"type:varchar(250);not null" json:"agentName"`
	DistrictId          string         `gorm:"type:char(7);not null" json:"districtId"`
	Address             string         `gorm:"type:text;not null" json:"address"`
	NoHandphone         string         `gorm:"type:varchar(12);not null" json:"noHandphone"`
	AgentRating         *float32       `gorm:"type:float" json:"agentRating,omitempty"`
	TotalReviewCustomer uint           `json:"totalReviewTransactions,omitempty"`
}
