package auth

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type tokenResponse struct {
	Token string `json:"token"`
}

type Service interface {
	Login(user models.User) *helper.Res
}

type service struct {
	auth interfaces.UserRepository
}

func NewService(auth interfaces.UserRepository) *service {
	return &service{auth}
}

func (s *service) Login(user models.User) *helper.Res {
	data, err := s.auth.FindByEmail(user.Email)
	if err != nil {
		response := helper.ResponseJSON("User Not Found", 404, "error", nil)
		return response
	}

	if !helper.CheckPassword(data.Password, user.Password) {
		response := helper.ResponseJSON("Internal Server Error", 500, "error", nil)
		return response
	}

	id := data.ID.Hex()

	token := helper.NewToken(id, data.Email, data.Name)
	tokens, err := token.Create()
	if err != nil {
		response := helper.ResponseJSON("Internal Server Error", 500, "error", nil)
		return response
	}

	response := helper.ResponseJSON("Success", 200, "OK", tokenResponse{Token: tokens})
	return response
}
