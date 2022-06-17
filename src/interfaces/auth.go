package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
)

type AuthService interface {
	Login(user models.User) (*helper.Res, error)
}
