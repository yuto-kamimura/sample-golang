package handler

import (
	"net/http"
	"sampleApi/entity"

	"github.com/labstack/echo/v4"
)

type (
	DeleteItemsRequest struct {
		IDs []int `json:"ids"`
	}
)

func DeleteItems(c echo.Context) error {
	req := &DeleteItemsRequest{}
	if err := c.Bind(req); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	var items entity.Items
	if err := items.DeleteItemsByItemID(req.IDs); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	if err := c.JSON(http.StatusOK, "delete items"); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return nil
}
