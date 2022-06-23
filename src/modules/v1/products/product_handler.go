package products

import (
	"fmt"
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service interfaces.ProductService
}

func NewHandler(service interfaces.ProductService) *handler {
	return &handler{service}
}

func (h *handler) GetAllProducts(c echo.Context) error {
	data, err := h.service.FindAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetProductDetails(c echo.Context) error {
	id := c.Param("id")
	data, err := h.service.GetUserID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) CreateProduct(c echo.Context) error {
	var product models.Product
	if err := c.Bind(&product); err != nil {
		return err
	}

	data, err := h.service.Insert(&product)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var product models.Product

	if err := c.Bind(&product); err != nil {
		return err
	}

	res, err := h.service.Update(id, &product)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}

func (h *handler) DeletProduct(c echo.Context) error {
	id := c.Param("id")

	res, err := h.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) QueryProducts(c echo.Context) error {
	search := c.QueryParam("s")
	sort := c.QueryParam("sort")
	page := c.QueryParam("page")
	data, err := h.service.Search(page, search, sort)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) CreateReview(c echo.Context) error {
	user_id := c.Request().Header.Get("user_id")
	name := c.Request().Header.Get("user_name")

	id := c.Param("id")

	var review models.Review

	if err := c.Bind(&review); err != nil {
		return err
	}

	review.UserID = user_id
	review.ProductID = id
	review.Fullname = name

	res, err := h.service.InsertReview(&review)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) GetAllReviewByProductId(c echo.Context) error {
	query := c.QueryParam("id")
	data, err := h.service.GetReviews(query)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) DeleteReview(c echo.Context) error {
	id := c.QueryParam("id")
	var review *models.ReviewInput

	if err := c.Bind(&review); err != nil {
		return err
	}

	res, err := h.service.DeleteReview(id, review)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) UploadImages(c echo.Context) error {
	id := c.QueryParam("id")
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	var res *helper.Res

	files := form.File["images"]

	for i := 0; i < len(files); i++ {
		src, err := files[i].Open()
		if err != nil {
			return err
		}
		defer src.Close()

		res, err = h.service.UploadImages(id, src, files[i])
		if err != nil {
			return c.JSON(http.StatusNotAcceptable, err.Error())
		}

	}

	return c.JSON(http.StatusOK, res)
}
