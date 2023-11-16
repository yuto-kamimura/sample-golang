package handler

import (
	"net/http"
	"sampleApi/db"
	"sampleApi/entity"
	"time"

	"github.com/labstack/echo/v4"
)

type CreateItemRequest struct {
	Name string `json:"Name"`
}

func CreateItem(c echo.Context) error {
	req := &CreateItemRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}
	dbInstance := db.GetDB()

	item := entity.Item{
		ID:        0, // Itemはautoincrementを指定しているため、自動的にIDが振られます
		Name:      req.Name,
		CreatedAt: time.Now(),
	}
	if err := dbInstance.Create(&item).Error; err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// レスポンス返却
	if err := c.JSON(http.StatusCreated, "ok!"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
