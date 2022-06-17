package transaction

import (
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
	transaction.CreatedAt = time.Now()
	transaction.UpdatedAt = time.Now()

	transactionID, err := s.repository.Insert(transaction)
	if err != nil {
		return nil, err
	}

	paymentTransaction := models.Payment{
		ID:     transactionID,
		Amount: transaction.Amount,
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(paymentTransaction, transaction.User)
	if err != nil {
		return nil, err
	}

	transaction.PaymentURL = paymentUrl

	data, err := s.repository.Update(transactionID, transaction)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) ProcessPayment(input *models.TransactionNotification) (*helper.Res, error) {
	transaction, err := s.repository.FindByID(input.OrderID)
	if err != nil {
		return nil, err
	}

	if input.PaymentType == "credit_card" && input.TransactionStatus == "capture" && input.FraudStatus == "accept" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "settlement" {
		transaction.Status = "paid"
	} else if input.TransactionStatus == "deny" || input.TransactionStatus == "expired" || input.TransactionStatus == "cancel" {
		transaction.Status = "cancelled"
	}

	id := transaction.ID.Hex()

	updateTransaction, err := s.repository.Update(id, transaction)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", updateTransaction)
	return res, nil
}
