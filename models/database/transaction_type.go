package database

import (
	"gorm.io/gorm"
	"time"
)

type TransactionType struct {
	CreatedAt                time.Time              `json:"-"`
	UpdatedAt                time.Time              `json:"-"`
	DeletedAt                gorm.DeletedAt         `gorm:"index" json:"-"`
	Id                       *uint                  `gorm:"autoIncrement;primary key" json:"-"`
	ServiceTypeTransactionId uint                   `json:"-"`
	ServiceTypeTransaction   ServiceTypeTransaction `gorm:"foreignkey:ServiceTypeTransactionId;references:Id" json:"serviceTypeTransaction"`
	NameTypeTransaction      string                 `gorm:"varchar(255);not null" json:"nameTransactionType"`
}
