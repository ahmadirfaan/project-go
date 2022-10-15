package database

import (
	"gorm.io/gorm"
	"time"
)

type Transactions struct {
	CreatedAt         time.Time       `json:"createdAt"`
	UpdatedAt         time.Time       `json:"-"`
	DeletedAt         gorm.DeletedAt  `gorm:"index" json:"-"`
	Id                *uint           `gorm:"autoIncrement;primary key" json:"id"`
	TransactionTypeId uint            `gorm:"not null" json:"-"`
	TransactionType   TransactionType `gorm:"foreignkey:TransactionTypeId;references:Id" json:"transactionType"`
	CustomerId        uint            `gorm:"not null" json:"-"`
	UserCustomer      User            `gorm:"foreignkey:CustomerId;references:Id" json:"userCustomer"`
	AgentId           uint            `gorm:"not null" json:"-"`
	UserAgent         User            `gorm:"foreignkey:AgentId;references:Id" json:"userAgent"`
	Address           string          `gorm:"type:text;not null" json:"address"`
	DistrictId        string          `gorm:"type:char(7);not null" json:"districtId"`
	Amount            uint64          `gorm:"not null" json:"amount"`
	StatusTransaction uint8           `gorm:"not null;type:tinyint" json:"statusTransaction"`
	Rating            uint8           `json:"rating,omitempty"`
}
