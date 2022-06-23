package auth

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	UserRepository interfaces.UserRepository
}

func NewService(auth interfaces.UserRepository) *service {
	return &service{auth}
}

func (s *service) Login(user models.User) (string, error) {
	data, err := s.UserRepository.FindByEmail(user.Email)
	if err != nil {
		return "", err
	}

	if !helper.CheckPassword(data.Password, user.Password) {
		return "", err
	}

	id := data.ID.Hex()

	token := helper.NewToken(id, data.Email, data.Name, data.Role)
	tokens, err := token.Create()
	if err != nil {
		return "", err
	}

	return tokens, nil
}
