package database

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	Id          uint           `gorm:"autoIncrement;primary key" json:"-"`
	Name        *string        `gorm:"varchar(255);not null"`
	NoHandphone *string        `gorm:"varchar(12);not null"`
}
