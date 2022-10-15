package service

import (
	"errors"
	"github.com/ahmadirfaan/project-go/models/database"
	"github.com/ahmadirfaan/project-go/models/web"
	"github.com/ahmadirfaan/project-go/repositories"
	"github.com/ahmadirfaan/project-go/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"strconv"
)

type TransactionService interface {
	CreateTransaction(request web.CreateTransactionRequest, customerId string) error
	GetAllTransactionByUserId(userId string) ([]database.Transactions, error)
	IsUserAgent(userId string) (*bool, error)
	DeleteTransaction(transactionId string, userId string) (int, error)
	GiveRatingTransaction(request web.RequestRating, userId string, transactionId string) (int, error)
	ChangeStatusTransaction(transactionId string, userId uint, request web.ChangeTransactionRequest) (int, error)
}

type transactionService struct {
	transactionsRepository    repositories.TransactionRepository
	transactionTypeRepository repositories.TransactionTypeRepository
	locationServices          LocationService
	userRepository            repositories.UserRepository
	agentRepository           repositories.AgentRepository
	DB                        *gorm.DB
}

func NewTransactionService(tr repositories.TransactionRepository,
	ttr repositories.TransactionTypeRepository,
	lr LocationService,
	ur repositories.UserRepository,
	ar repositories.AgentRepository,
	db *gorm.DB) TransactionService {
	return &transactionService{
		transactionsRepository:    tr,
		DB:                        db,
		transactionTypeRepository: ttr,
		locationServices:          lr,
		userRepository:            ur,
		agentRepository:           ar,
	}
}

func (t *transactionService) GetAllTransactionByUserId(userId string) ([]database.Transactions, error) {
	transactions, err := t.transactionsRepository.FindTransactionWithUserId(userId)
	return transactions, err
}

func (t *transactionService) CreateTransaction(request web.CreateTransactionRequest, customerId string) error {
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return err
	}
	err = t.validateUserForTransaction(request, customerId)
	if err != nil {
		return err
	}
	tx := t.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return err
	}
	customerIdUint, _ := strconv.ParseUint(customerId, 10, 32)
	transaction := database.Transactions{
		TransactionTypeId: request.TransactionTypeId,
		CustomerId:        uint(customerIdUint),
		AgentId:           request.AgentId,
		Amount:            request.Amount,
		Address:           request.Address,
		DistrictId:        request.DistrictId,
		StatusTransaction: 0,
	}
	transaction, err = t.transactionsRepository.WithTrx(tx).Save(transaction)
	if err != nil {
		tx.Debug().Rollback()
		return err
	}
	return tx.Commit().Error
}

func (t *transactionService) GiveRatingTransaction(request web.RequestRating, userId string, transactionId string) (int, error) {
	returnOk := fiber.StatusOK
	defaultError := fiber.StatusInternalServerError
	err := utils.NewValidator().Struct(&request)
	if err != nil {
		return defaultError, err
	}
	code, err := t.validateUserGiveRating(userId, transactionId)
	if err != nil {
		return code, err
	}
	tx := t.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()
	if err := tx.Error; err != nil {
		return defaultError, err
	}
	transaction, err := t.transactionsRepository.FindTransactionById(transactionId)
	if err != nil {
		tx.Debug().Rollback()
		return defaultError, err
	}
	userIdString := strconv.FormatUint(uint64(transaction.AgentId), 10)
	user, err := t.userRepository.FindByUserId(userIdString)
	if err != nil {
		tx.Debug().Rollback()
		return defaultError, err
	}
	agent, err := t.agentRepository.WithTrx(tx).FindByAgentId(*user.AgentId)
	if err != nil {
		tx.Debug().Rollback()
		return defaultError, err
	}
	var insertRating float32
	insertTotalReview := agent.TotalReviewCustomer + 1
	if agent.AgentRating == nil {
		insertRating = float32(request.Rating)
	} else {
		insertRating = (float32(request.Rating) + *agent.AgentRating) / 2
	}
	err = t.transactionsRepository.WithTrx(tx).GiveRating(transactionId, request.Rating)
	if err != nil {
		tx.Debug().Rollback()
		return defaultError, err
	}
	err = t.agentRepository.WithTrx(tx).GiveRatingAgent(*agent.Id, insertRating, insertTotalReview)
	if err != nil {
		tx.Debug().Rollback()
		return defaultError, err
	}
	return returnOk, tx.Commit().Error
}

func (t *transactionService) IsUserAgent(userId string) (*bool, error) {
	agentID, err := strconv.ParseUint(userId, 10, 64)
	if err != nil {
		return nil, errors.New("error handling for converstion customerId")
	}
	isAgent, err := t.userRepository.IsUserAgent(uint(agentID))
	return isAgent, err
}

func (t *transactionService) ChangeStatusTransaction(transactionId string, userId uint, request web.ChangeTransactionRequest) (int, error) {
	err := utils.NewValidator().Struct(&request)
	defaultReturn := 0
	if err != nil {
		return defaultReturn, err
	}
	code, err := t.validateUpdateTransaction(transactionId, userId, request.StatusTransaction)
	if err != nil {
		return code, err
	}
	err = t.transactionsRepository.UpdateStatusTransaction(transactionId, request.StatusTransaction)
	if err != nil {
		return defaultReturn, err
	}
	return defaultReturn, err

}

func (t *transactionService) validateUpdateTransaction(transactionId string, userId uint, statusTransaction uint8) (int, error) {
	transaction, err := t.transactionsRepository.FindTransactionById(transactionId)
	if err != nil {
		return fiber.StatusBadRequest, err
	}
	waitingStatus := uint8(0)
	if statusTransaction == uint8(3) && transaction.StatusTransaction == waitingStatus {
		return fiber.StatusBadRequest, errors.New("can't change status to done")
	}
	if isUserDoTransaction := isUserDoTransaction(*transaction, userId); !isUserDoTransaction {
		return fiber.StatusForbidden, errors.New("you're not allowed change transaction")
	}
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	if transaction.Id == nil {
		return fiber.StatusOK, errors.New("not found transaction")
	}
	switch transaction.StatusTransaction {
	case uint8(2):
		return fiber.StatusBadRequest, errors.New("cant change because the transaction canceled")
	case uint8(3):
		return fiber.StatusBadRequest, errors.New("cant change because the transaction done")
	}
	isAgent, err := t.userRepository.IsUserAgent(userId)
	if !*isAgent && statusTransaction == uint8(1) {
		return fiber.StatusForbidden, errors.New("customer can't change the order waiting transaction")
	}
	return fiber.StatusOK, nil
}

func isUserDoTransaction(transaction database.Transactions, userId uint) bool {
	userTransactions := []uint{transaction.AgentId, transaction.CustomerId}
	for _, v := range userTransactions {
		if userId == v {
			return true
		}
	}
	return false
}

func (t *transactionService) validateUserGiveRating(userId string, transactionId string) (int, error) {
	transaction, err := t.transactionsRepository.FindTransactionById(transactionId)
	if err != nil {
		return fiber.StatusBadRequest, err
	}
	if transaction.Rating != 0 {
		return fiber.StatusBadRequest, errors.New("transaction have a rating")
	}
	if transaction.StatusTransaction != uint8(3) {
		return fiber.StatusBadRequest, errors.New("please complete transaction to give rating")
	}
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	userIdInt, _ := strconv.Atoi(userId)
	if isUserDoTransaction := isUserDoTransaction(*transaction, uint(userIdInt)); !isUserDoTransaction {
		return fiber.StatusForbidden, errors.New("you're not allowed change transaction")
	}
	if isUserAgent, _ := t.IsUserAgent(userId); *isUserAgent {
		return fiber.StatusForbidden, errors.New("must user customer to can giver rating")
	}
	return fiber.StatusOK, nil
}

func (t *transactionService) DeleteTransaction(transactionId string, userId string) (int, error) {
	code, err := t.validateDeleteTransaction(transactionId, userId)
	defaultReturn := 200
	if err != nil {
		return code, err
	}
	err = t.transactionsRepository.DeleteTransactionById(transactionId)
	if err != nil {
		return fiber.StatusInternalServerError, err
	}
	return defaultReturn, err
}

func (t *transactionService) validateDeleteTransaction(transactionId string, userId string) (int, error) {
	transaction, err := t.transactionsRepository.FindTransactionById(transactionId)
	if err != nil {
		return fiber.StatusBadRequest, err
	}
	userIdInt, _ := strconv.Atoi(userId)
	if isUserDoTransaction := isUserDoTransaction(*transaction, uint(userIdInt)); !isUserDoTransaction {
		return fiber.StatusForbidden, errors.New("you're not allowed change transaction")
	}
	if transaction.StatusTransaction != uint8(2) {
		return fiber.StatusBadRequest, errors.New("must canceled transaction to be deleted")
	}

	return fiber.StatusOK, nil
}

func (t *transactionService) validateUserForTransaction(request web.CreateTransactionRequest, customerId string) error {
	//validate userId must not same with agentId
	customerIdUint, _ := strconv.ParseUint(customerId, 10, 32)
	if strconv.FormatUint(uint64(request.AgentId), 32) == customerId {
		return errors.New("agentId and customerId must not the same")
	}
	//validate exist districtId
	_, err := t.locationServices.FindDistrictById(request.DistrictId)
	if err != nil {
		return err
	}
	//validate exist transactionType
	err = t.transactionTypeRepository.FindById(request.TransactionTypeId)
	if err != nil {
		return err
	}
	//validate if the request customerId is customer
	isAgent, err := t.IsUserAgent(customerId)
	//validate the exist is customerId
	isExistUserCustomer, err := t.userRepository.IsExist(uint(customerIdUint))
	//validate if the request agentId is agent
	isNotAgent, err := t.IsUserAgent(strconv.FormatUint(uint64(request.AgentId), 10))
	//validate the exist is AgentId
	isExistUserAgent, err := t.userRepository.IsExist(request.AgentId)
	if err != nil {
		return err
	}
	if !*isNotAgent {
		return errors.New("your input agentId is the user customer")
	}
	if *isAgent {
		return errors.New("your input customerId is The User agent")
	}
	if !isExistUserCustomer {
		return errors.New("your customerId is not exist")
	}
	if !isExistUserAgent {
		return errors.New("your agentId is not exist")
	}
	return err
}
