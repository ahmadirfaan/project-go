package database

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Id         *uint          `gorm:"autoIncrement;primary key" json:"id"`
	Username   string         `gorm:"uniqueIndex;type:varchar(255);not null" json:"-"`
	Password   string         `gorm:"type:varchar(255);not null" json:"-"`
	RoleId     uint           `json:"-"`
	Role       Role           `gorm:"foreignkey:RoleId;references:Id" json:"-"`
	CustomerId *uint          `json:"-"`
	Customer   *Customer      `gorm:"foreignkey:CustomerId;references:Id" json:"customer,omitempty"`
	AgentId    *uint          `json:"-"`
	Agent      *Agent         `gorm:"foreignkey:AgentId;references:Id" json:"agent,omitempty"`
}
