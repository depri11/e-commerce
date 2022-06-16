package products

import (
	"fmt"
	"time"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	repository interfaces.ProductRepository
}

func NewService(repo interfaces.ProductRepository) *service {
	return &service{repo}
}

func (s *service) FindAll() (*helper.Res, error) {
	product, err := s.repository.FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", product)
	return res, nil
}

func (s *service) GetUserID(id string) (*helper.Res, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Insert(product *models.Product) (*helper.Res, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	data, err := s.repository.Insert(product)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Update(id string, product *models.Product) (*helper.Res, error) {
	product.UpdatedAt = time.Now()

	data, err := s.repository.Update(id, product)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Delete(id string) (*helper.Res, error) {
	data, err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Search(query string) (*helper.Res, error) {
	data, err := s.repository.Search(query)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}
