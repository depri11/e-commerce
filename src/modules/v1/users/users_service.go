package users

import (
	"errors"
	"fmt"
	"mime/multipart"
	"time"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
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
		res := helper.ResponseJSON("data user not found", 404, "error", errors.New("no found data"))
		return res, nil
	}

	res := helper.ResponseJSON("Success", 200, "OK", user)
	return res, nil
}

func (s *service) GetUserID(id string) (*helper.Res, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		res := helper.ResponseJSON("User Not Found", 404, "error", errors.New("no found data"))
		return res, nil
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Insert(input *input.UserInput) (*helper.Res, error) {
	var user models.User

	user.Email = input.Email
	user.Name = input.Name
	user.Gender = input.Gender
	user.Avatar = input.Avatar
	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()
	user.Role = "user"
	hashPass, err := helper.HashPassword(input.Password)
	if err != nil {
		res := helper.ResponseJSON("Error password", 404, "error", err.Error())
		return res, nil
	}

	user.Password = hashPass

	data, err := s.repository.Insert(&user)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Update(id string, input *input.UserInput) (*helper.Res, error) {
	_, err := s.repository.FindByID(id)
	if err != nil {
		res := helper.ResponseJSON("User Not Found", 404, "error", errors.New("no found data"))
		return res, nil
	}

	var user models.User

	user.Email = input.Email
	user.Name = input.Name
	user.Gender = input.Gender
	user.Avatar = input.Avatar
	hashPass, err := helper.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}
	user.Password = hashPass
	user.UpdatedAt = time.Now()

	data, err := s.repository.Update(id, &user)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Delete(id string) (*helper.Res, error) {
	_, err := s.repository.FindByID(id)
	if err != nil {
		res := helper.ResponseJSON("User Not Found", 404, "error", errors.New("no found data"))
		return res, nil
	}

	data, err := s.repository.Delete(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) ForgotPassword(input *input.ForgotPasswordInput) (*helper.Res, error) {
	data, err := s.repository.FindByEmail(input.Email)
	if err != nil {
		return nil, err
	}

	token := helper.GenToken(32)

	data.ResetPassToken = token
	data.ResetPassExpire = time.Now().Add(time.Hour * 2)

	id := data.ID.Hex()

	url := fmt.Sprintf("http://localhost:4000/api/v1/password/reset/%s", token)

	email := []string{data.Email}
	cc := []string{data.Email}

	helper.SendMail(email, cc, "Reset Password", url)

	_, err = s.repository.Update(id, data)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", url)
	return res, nil
}

func (s *service) ResetPassword(token string, user *models.User) (*helper.Res, error) {
	data, err := s.repository.FindByResetPassToken(token)
	if err != nil {
		return nil, err
	}

	if data.ResetPassExpire.Before(time.Now()) {
		return nil, errors.New("token expired")
	}

	hash, err := helper.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	data.Password = hash
	data.ResetPassToken = ""
	data.ResetPassExpire = time.Now()

	id := data.ID.Hex()

	r, err := s.repository.Update(id, data)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", r)
	return res, nil
}

func (s *service) UploadAvatar(id string, file multipart.File, handle *multipart.FileHeader) (*helper.Res, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	var input models.User
	input.Name = data.Name
	input.Email = data.Email
	input.Gender = data.Gender

	avatar := "avatar"

	images, err := helper.UploadImages(avatar, file, handle)
	if err != nil {
		return nil, err
	}

	input.Avatar = images.URL

	r, err := s.repository.Update(id, &input)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", r)
	return res, nil
}
