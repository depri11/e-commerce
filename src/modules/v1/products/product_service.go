package products

import (
	"fmt"
	"time"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
)

type service struct {
	repository     interfaces.ProductRepository
	userRepository interfaces.UserRepository
}

func NewService(repo interfaces.ProductRepository, userRepository interfaces.UserRepository) *service {
	return &service{repo, userRepository}
}

func (s *service) FindAll() (*helper.Res, error) {
	product, err := s.repository.FindAll()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", product)
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

func (s *service) Insert(product *models.Product) (*helper.Res, error) {
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()
	product.NumOfReviews = 0
	data, err := s.repository.Insert(product)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) Update(id string, product *models.Product) (*helper.Res, error) {
	product.UpdatedAt = time.Now()

	data, err := s.repository.Update(id, product)
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

func (s *service) Search(page, search, sort string) (*helper.Res, error) {
	data, err := s.repository.Search(page, search, sort)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) InsertReview(review *models.Review) (*helper.Res, error) {

	reviewData, err := s.repository.FindByID(review.ProductID)
	if err != nil {
		return nil, err
	}

	reviewData.NumOfReviews = reviewData.NumOfReviews + 1
	reviewData.Reviews = append(reviewData.Reviews, review)

	data, err := s.repository.Update(review.ProductID, reviewData)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) GetReviews(id string) (*helper.Res, error) {
	data, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) DeleteReview(id string) (*helper.Res, error) {
	data, err := s.repository.DeleteReview(id)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}
