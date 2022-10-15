package database

import (
	"time"

	"gorm.io/gorm"
)

type Role struct {
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	Id        uint           `gorm:"autoIncrement;primary key" json:"-"`
	Role      string         `gorm:"type:varchar(50)"`
}
