package auth

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type service struct {
	auth interfaces.UserRepository
}

func NewService(auth interfaces.UserRepository) *service {
	return &service{auth}
}

func (s *service) Login(user models.User) (*helper.Res, error) {
	data, err := s.auth.FindByEmail(user.Email)
	if err != nil {
		return nil, err
	}

	if !helper.CheckPassword(data.Password, user.Password) {
		return nil, err
	}

	id := data.ID.Hex()

	token := helper.NewToken(id, data.Email, data.Name, data.Role)
	tokens, err := token.Create()
	if err != nil {
		return nil, err
	}

	response := helper.ResponseJSON("Success", 200, "OK", tokenResponse{Token: tokens})
	return response, nil
}
