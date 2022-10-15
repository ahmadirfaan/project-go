package repositories

import (
	"errors"
	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type TransactionTypeRepository interface {
	FindById(idTransactionType uint) error
}

type transactionTypeRepository struct {
	DB *gorm.DB
}

func NewTransactionTypeRepository(db *gorm.DB) TransactionTypeRepository {
	return &transactionTypeRepository{
		DB: db,
	}
}

func (t transactionTypeRepository) FindById(idTransactionType uint) error {
	var transactionType database.TransactionType
	err := t.DB.Debug().First(&transactionType, idTransactionType).Error
	if transactionType.Id != nil {
		return err
	} else {
		return errors.New("No found record transactionType")
	}
}
