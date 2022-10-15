package repositories

import (
	"log"

	"github.com/ahmadirfaan/project-go/models/database"
	"gorm.io/gorm"
)

type TransactionRepository interface {
	Save(transaction database.Transactions) (database.Transactions, error)
	FindTransactionWithUserId(userId string) ([]database.Transactions, error)
	UpdateStatusTransaction(transactionId string, statusTransaction uint8) error
	FindTransactionById(transactionId string) (*database.Transactions, error)
	DeleteTransactionById(transactionId string) error
	GiveRating(transactionId string, rating uint8) error
	WithTrx(trxHandle *gorm.DB) transactionRepo
}

type transactionRepo struct {
	DB *gorm.DB
}

func NewTransactionRepository(db *gorm.DB) TransactionRepository {
	return &transactionRepo{
		DB: db,
	}
}

func (t transactionRepo) Save(transactions database.Transactions) (database.Transactions, error) {
	err := t.DB.Debug().Create(&transactions).Error
	log.Printf("Transaction:%+v\n", transactions)
	return transactions, err
}

func (t transactionRepo) FindTransactionWithUserId(userId string) ([]database.Transactions, error) {
	var transactionList []database.Transactions
	err := t.DB.Debug().Where("agent_id = ? or customer_id = ?", userId, userId).
		Preload("UserCustomer").Preload("UserCustomer.Customer").
		Preload("TransactionType").Preload("TransactionType.ServiceTypeTransaction").
		Preload("UserAgent").Preload("UserAgent.Agent").Find(&transactionList).Error
	return transactionList, err
}

func (t transactionRepo) GiveRating(transactionId string, rating uint8) error {
	var transaction *database.Transactions
	err := t.DB.Debug().Model(&transaction).Where("id = ?", transactionId).Update("rating", rating).Error
	return err
}

func (t transactionRepo) FindTransactionById(transactionId string) (*database.Transactions, error) {
	var transaction *database.Transactions
	err := t.DB.Debug().Where("id = ?", transactionId).First(&transaction).Error
	return transaction, err
}

func (t transactionRepo) UpdateStatusTransaction(transactionId string, statusTransaction uint8) error {
	var transaction database.Transactions
	err := t.DB.Debug().Model(&transaction).Where(" id = ?", transactionId).
		Update("status_transaction", statusTransaction).Error
	return err
}

func (t transactionRepo) DeleteTransactionById(transactionId string) error {
	var transaction database.Transactions
	err := t.DB.Debug().Model(&transaction).Delete(&transaction, transactionId).Error
	return err
}

func (t transactionRepo) WithTrx(trxHandle *gorm.DB) transactionRepo {
	if trxHandle == nil {
		log.Print("Transaction Database not  found")
		return t
	}
	t.DB = trxHandle
	return t
}
