package handler

import (
	"net/http"
	"sampleApi/entity"
	"time"

	"github.com/labstack/echo/v4"
)

type (
	CreateUserRequest struct {
		Name     string `json:"name"`
		Password string `json:"password"`
		Email    string `json:"email"`
	}
)

func CreateUser(c echo.Context) error {
	req := &CreateUserRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user := entity.User{
		Username:  req.Name,
		Password:  req.Password,
		Email:     req.Email,
		UpdatedAt: time.Now(),
	}
	if err := user.EncriptPassword(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := user.CreateUser(); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.JSON(http.StatusCreated, "OK"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
