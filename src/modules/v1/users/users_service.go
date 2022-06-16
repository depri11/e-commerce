package users

import (
	"fmt"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	repository interfaces.UserRepository
}

func NewService(repository interfaces.UserRepository) *service {
	return &service{repository}
}

func (s *service) FindAll() (*helper.Res, error) {
	user, err := s.repository.FindAll()
	if err != nil {
		fmt.Println(err.Error())

		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", user)
	return res, nil
}

func (s *service) Insert(user *models.User) (*helper.Res, error) {
	data, err := s.repository.Insert(user)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}
