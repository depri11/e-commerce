package products

import (
	"fmt"
	"math"
	"mime/multipart"
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

	var avg float64

	reviewData.NumOfReviews = reviewData.NumOfReviews + 1
	reviewData.Reviews = append(reviewData.Reviews, review)
	for _, v := range reviewData.Reviews {
		avg += v.Rating
	}

	reviewData.Ratings = avg / float64(reviewData.NumOfReviews)

	reviewData.Ratings = math.Round(reviewData.Ratings*10) / 10

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

func (s *service) DeleteReview(id string, review *models.ReviewInput) (*helper.Res, error) {
	dataProduct, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	var avg float64

	dataProduct.NumOfReviews = len(review.Review)
	for _, v := range review.Review {
		avg += v.Rating
	}

	dataProduct.Ratings = avg / float64(dataProduct.NumOfReviews)
	dataProduct.UpdatedAt = time.Now()
	dataProduct.Reviews = review.Review

	data, err := s.repository.Update(id, dataProduct)
	if err != nil {
		return nil, err
	}

	res := helper.ResponseJSON("Success", 200, "OK", data)
	return res, nil
}

func (s *service) UploadImages(id string, file multipart.File, handle *multipart.FileHeader) (*helper.Res, error) {
	_, err := s.repository.FindByID(id)
	if err != nil {
		return nil, err
	}

	ext := "products"

	images, err := helper.UploadImages(ext, file, handle)
	if err != nil {
		return nil, err
	}
	fmt.Println(images.SecureURL)

	fmt.Println(images.URL)
	return nil, nil

	// r, err := s.repository.Update(id, data)
	// if err != nil {
	// 	return nil, err
	// }

	// res := helper.ResponseJSON("Success", 200, "OK", r)
	// return res, nil
}
