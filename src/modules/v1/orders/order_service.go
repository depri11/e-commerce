package orders

import (
	"time"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	repository interfaces.OrderRepository
	userRepo   interfaces.UserRepository
}

func NewService(repository interfaces.OrderRepository, userRepo interfaces.UserRepository) *service {
	return &service{repository, userRepo}
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

func (s *service) Create(id string, order models.Order) (*helper.Res, error) {
	order.PaidAt = time.Now()
	order.Status = "pending"
	order.CreatedAt = time.Now()
	order.UpdatedAt = time.Now()
	order.UserID = id
	data, err := s.repository.Insert(order)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Update(id string, order models.Order) (*helper.Res, error) {
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
