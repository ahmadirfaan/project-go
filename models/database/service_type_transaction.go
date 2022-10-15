package database

import (
	"gorm.io/gorm"
	"time"
)

type ServiceTypeTransaction struct {
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Id          *uint          `gorm:"autoIncrement;primary key"`
	NameService *string        `gorm:"varchar(255);not null" json:"nameServiceTransaction"`
}
