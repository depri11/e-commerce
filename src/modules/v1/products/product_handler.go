package products

import (
	"fmt"
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service interfaces.ProductService
}

func NewHandler(service interfaces.ProductService) *handler {
	return &handler{service}
}

func (h *handler) GetProducts(c echo.Context) error {
	data, err := h.service.FindAll()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetProduct(c echo.Context) error {
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
