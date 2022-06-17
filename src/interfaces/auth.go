package interfaces

import (
	"github.com/depri11/e-commerce/src/database/models"
)

type AuthService interface {
	Login(user models.User) (string, error)
}
