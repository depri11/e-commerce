package payment

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/veritrans/go-midtrans"
)

type service struct {
	repository        interfaces.TransactionRepository
	productRepository interfaces.ProductRepository
}

func NewService(repository interfaces.TransactionRepository, productRepository interfaces.ProductRepository) *service {
	return &service{repository, productRepository}
}

func (s *service) GetPaymentURL(transaction *models.Transaction, user *models.User) (string, error) {
	id := transaction.ID.Hex()

	midclient := midtrans.NewClient()
	midclient.ClientKey = "SB-Mid-client-Fg55R6OSZynaFTNA"
	midclient.ServerKey = "SB-Mid-server-Dc-ShUJ8AJYb9EvWzoVZKCq0"
	midclient.APIEnvType = midtrans.Sandbox

	coreGateway := midtrans.SnapGateway{
		Client: midclient,
	}

	chargeReq := &midtrans.SnapReq{
		CustomerDetail: &midtrans.CustDetail{
			FName: user.Name,
			Email: user.Email,
		},
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  id,
			GrossAmt: int64(transaction.Amount),
		},
	}

	snapTokenResp, err := coreGateway.GetToken(chargeReq)
	if err != nil {
		return "", err
	}

	return snapTokenResp.RedirectURL, nil
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

	product, err := s.productRepository.FindByID(updateTransaction.ProductID)
	if err != nil {
		return nil, err
	}

	product.Stock = product.Stock - 1

	_, err = s.productRepository.Update(product.ID.Hex(), product)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", updateTransaction)
	return res, nil
}
