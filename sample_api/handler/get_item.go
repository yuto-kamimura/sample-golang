package handler

import (
	"net/http"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

type (
	GetItemsRequest struct {
		IDs []int `json:"ids"`
	}

	GetItemsResponse struct {
		Items entity.Items
	}
)

func GetItems(c echo.Context) error {
	req := &GetItemsRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	items := entity.Items{}

	if err := items.SelectItemsByItemID(req.IDs); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.JSON(http.StatusOK, items); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	return nil
}
