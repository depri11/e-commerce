package users

import (
	"net/http"

	"github.com/depri11/e-commerce/src/database/models"
	"github.com/depri11/e-commerce/src/helper"
	"github.com/depri11/e-commerce/src/input"
	"github.com/depri11/e-commerce/src/interfaces"
	"github.com/labstack/echo/v4"
)

type handler struct {
	service interfaces.UserService
}

func NewHandler(service interfaces.UserService) *handler {
	return &handler{service}
}

func (h *handler) FindAll(c echo.Context) error {
	data, err := h.service.FindAll()
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetUserID(c echo.Context) error {
	id := c.Param("id")
	data, err := h.service.GetUserID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) GetUserDetails(c echo.Context) error {
	id := c.Request().Header.Get("user_id")
	data, err := h.service.GetUserID(id)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, data)
}

func (h *handler) UpdateProfile(c echo.Context) error {
	id := c.Request().Header.Get("user_id")
	var input input.UserInput

	if err := c.Bind(&input); err != nil {
		return err
	}

	if err := helper.ValidationError(input); err != nil {
		return c.JSON(400, err.Error())
	}

	res, err := h.service.Update(id, &input)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) Register(c echo.Context) error {
	var input input.UserInput
	if err := c.Bind(&input); err != nil {
		return err
	}

	if err := helper.ValidationError(input); err != nil {
		return c.JSON(400, err.Error())
	}

	res, err := h.service.Insert(&input)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) UpdateUser(c echo.Context) error {
	id := c.Param("id")
	var input input.UserInput

	if err := c.Bind(&input); err != nil {
		return err
	}

	if err := helper.ValidationError(input); err != nil {
		return c.JSON(400, err.Error())
	}

	res, err := h.service.Update(id, &input)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)

}

func (h *handler) DeletUser(c echo.Context) error {
	id := c.Param("id")
	res, err := h.service.Delete(id)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) ForgotPassword(c echo.Context) error {
	var input input.ForgotPasswordInput
	if err := c.Bind(&input); err != nil {
		return err
	}

	if err := helper.ValidationError(input); err != nil {
		return c.JSON(400, err.Error())
	}

	res, err := h.service.ForgotPassword(&input)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) ResetPassword(c echo.Context) error {
	var user models.User
	token := c.Param("token")
	if err := c.Bind(&user); err != nil {
		return err
	}

	if err := helper.ValidationError(user); err != nil {
		return c.JSON(400, err.Error())
	}

	res, err := h.service.ResetPassword(token, &user)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}

func (h *handler) UploadAvatar(c echo.Context) error {
	id := c.Request().Header.Get("user_id")
	file, handle, err := c.Request().FormFile("avatar")
	if err != nil {
		return err
	}

	res, err := h.service.UploadAvatar(id, file, handle)
	if err != nil {
		return c.JSON(http.StatusNotAcceptable, err.Error())
	}

	return c.JSON(http.StatusOK, res)
}
