package users

import (
	"time"

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
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", user)
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

func (s *service) Insert(user *models.User) (*helper.Res, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	hashPass, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hashPass
	data, err := s.repository.Insert(user)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Update(id string, user *models.User) (*helper.Res, error) {
	user.UpdatedAt = time.Now()
	hash, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	data, err := s.repository.Update(id, user)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) UpdateProfile(id string, user *models.UpdateProfile) (*helper.Res, error) {
	user.UpdatedAt = time.Now()
	hash, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	user.Password = hash

	data, err := s.repository.UpdateProfile(id, user)
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
