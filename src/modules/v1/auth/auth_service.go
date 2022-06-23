package auth

import (
	"errors"

	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	UserRepository interfaces.UserRepository
}

func NewService(auth interfaces.UserRepository) *service {
	return &service{auth}
}

func (s *service) Login(input input.AuthLogin) (string, error) {
	data, err := s.UserRepository.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("email/password incorrect, please correct this")
	}

	if !helper.CheckPassword(data.Password, input.Password) {
		return "", errors.New("email/password incorrect, please correct this")
	}

	id := data.ID.Hex()

	token := helper.NewToken(id, data.Email, data.Name, data.Role)
	tokens, err := token.Create()
	if err != nil {
		return "", errors.New("failed to create token")
	}

	return tokens, nil
}
