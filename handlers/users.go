package users

import (
	"net/http"

	user "pocCurd/models"

	"github.com/labstack/echo"
)

type handler struct {
	UserModel user.UserModelImpl
}

func NewHandler(u user.UserModelImpl) *handler {
	return &handler{u}
}

func (h *handler) SetDetail(c echo.Context) error {
	u := new(user.Employee)
	if err := c.Bind(u); err != nil {
		return err
	}
	_, err := h.UserModel.PostEmploye(*u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, u)
	}
	return c.JSON(http.StatusOK, u)
}
func (h *handler) UpdateDetail(c echo.Context) error {
	u := new(user.Employee)
	if err := c.Bind(u); err != nil {
		return err
	}
	_, err := h.UserModel.PutEmploye(*u)
	if err != nil {
		c.JSON(http.StatusInternalServerError, u)
	}
	return c.JSON(http.StatusOK, u)
}

func (h *handler) DeleteDetails(c echo.Context) error {
	id := c.Param("id")

	_, err := h.UserModel.DeleteEmploye(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, id)
	}
	return c.JSON(http.StatusOK, id)
}

func (h *handler) GetDetail(c echo.Context) error {

	u, err := h.UserModel.FetchEmploye()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, u)
}
