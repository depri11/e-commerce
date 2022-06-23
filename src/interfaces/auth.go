package interfaces

import "github.com/depri11/e-commerce/src/input"

type AuthService interface {
	Login(input input.AuthLogin) (string, error)
}
