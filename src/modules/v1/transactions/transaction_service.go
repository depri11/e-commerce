package transaction

import (
	"fmt"
	"time"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	repository        interfaces.TransactionRepository
	productRepository interfaces.ProductRepository
	paymentService    interfaces.PaymentService
}

func NewService(repository interfaces.TransactionRepository, productRepository interfaces.ProductRepository, paymentService interfaces.PaymentService) *service {
	return &service{repository, productRepository, paymentService}
}

func (s *service) GetAll() (*helper.Res, error) {
	transaction, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", transaction)
	return res, nil
}

func (s *service) GetByID(id string) (*helper.Res, error) {
	transaction, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", transaction)
	return res, nil
}

func (s *service) GetByProductID(id string) (*helper.Res, error) {
	transaction, err := s.repository.FindByProductId(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", transaction)
	return res, nil
}

func (s *service) GetByUserID(id string) (*helper.Res, error) {
	transaction, err := s.repository.FindByUserId(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", transaction)
	return res, nil
}

func (s *service) Create(id string, transaction *models.Transaction) (*helper.Res, error) {
	transaction.UserID = id
	transaction.Status = "pending"
	transaction.Code = "ORDER-001"
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	transactionID, err := s.repository.Insert(transaction)
	if err != nil {
		return nil, err
	}

	newTransaction, err := s.repository.FindByID(transactionID)
	if err != nil {
		return nil, err
	}

	transaction.User.Name = "Devri"
	transaction.User.Email = "dev@gmail.com"

	// paymentTransaction := models.Payment{
	// 	ID:     transactionID,
	// 	Amount: transaction.Amount,
	// }

	newTransaction.Amount = transaction.Amount

	paymentUrl, err := s.paymentService.GetPaymentURL(newTransaction, &transaction.User)
	if err != nil {
		return nil, err
	}

	transaction.PaymentURL = paymentUrl

	fmt.Println(paymentUrl)

	data, err := s.repository.Update(newTransaction.ID.Hex(), transaction)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}