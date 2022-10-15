package repositories

import (
	"fmt"
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type CustomerRepository interface {
	Save(customer database.Customer) (database.Customer, error)
	WithTrx(trxHandle *gorm.DB) customerRepository
}

type customerRepository struct {
	DB *gorm.DB
}

func NewCustomerRepository(db *gorm.DB) CustomerRepository {
	return &customerRepository{
		DB: db,
	}
}

func (c customerRepository) Save(customer database.Customer) (database.Customer, error) {
	err := c.DB.Debug().Create(&customer).Error
	log.Printf("Customer Repositories:%+v\n ", customer)
	fmt.Println(err)
	return customer, err
}

func (c customerRepository) WithTrx(trxHandle *gorm.DB) customerRepository {
	if trxHandle == nil {
		log.Print("Transaction Database not found")
		return c
	}
	c.DB = trxHandle
	return c
}
