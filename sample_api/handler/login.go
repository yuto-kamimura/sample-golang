package handler

import (
	"net/http"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

type (
	LoginRequest struct {
		UserID   string `json:"user_id"`
		Password string `json:"password"`
	}
)

func Login(c echo.Context) error {
	req := &LoginRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	user := entity.User{
		Username: req.UserID,
		Password: req.Password,
	}

	target, err := entity.FindUser(user.Username)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := user.VerifyPassword(target); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := c.JSON(http.StatusOK, "successed login."); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil

}
