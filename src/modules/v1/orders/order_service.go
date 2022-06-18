package orders

import (
	"time"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	repository     interfaces.OrderRepository
	userRepository interfaces.UserRepository
	paymentService interfaces.PaymentService
}

func NewService(repository interfaces.OrderRepository, userRepository interfaces.UserRepository, paymentService interfaces.PaymentService) *service {
	return &service{repository, userRepository, paymentService}
}

func (s *service) GetAllOrders() (*helper.Res, error) {
	data, err := s.repository.FindAll()
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) FindByID(id string) (*helper.Res, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) FindByUserID(id string) (*helper.Res, error) {
	data, err := s.repository.FindByUserID(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Create(id string, order *models.Order) (*helper.Res, error) {
	orderID := helper.GenOrderID()

	order.PaidAt = time.Now()
	order.Status = "pending"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.UserID = id
	order.OrderID = orderID

	user, err := s.userRepository.FindByID(id)
	if err != nil {
		return nil, err
	}

	paymentUrl, err := s.paymentService.GetPaymentURL(orderID, order, user)
	if err != nil {
		return nil, err
	}

	order.PaymentURL = paymentUrl

	data, err := s.repository.Insert(order)
	if err != nil {
		return nil, err
	}

	// data, err = s.repository.UpdateByOrderID(orderID, order)
	// if err != nil {
	// 	return nil, err
	// }

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Update(id string, order *models.Order) (*helper.Res, error) {
	order.UpdatedAt = time.Now()
	data, err := s.repository.Update(id, order)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Delele(id string) (*helper.Res, error) {
	data, err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}
