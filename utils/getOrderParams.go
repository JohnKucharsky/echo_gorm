package utils

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/samber/lo"
	"strings"
)

type OrderParams struct {
	OrderBy string `json:"orderBy"`
	Desc    bool   `json:"desc"`
}

func GetOrderParams(c echo.Context, columns []string) (OrderParams, error) {
	sortOrder := c.QueryParam("sortOrder")
	orderBy := c.QueryParam("orderBy")

	if orderBy == "" {
		return OrderParams{}, nil
	}
	if !lo.Contains(columns, orderBy) {
		return OrderParams{}, errors.New(
			"param should be one of " + strings.Join(
				columns,
				",",
			),
		)
	}

	if sortOrder != "" && !lo.Contains([]string{"asc", "desc"}, sortOrder) {
		return OrderParams{}, errors.New(
			"param should be one either asc or desc",
		)
	}

	return OrderParams{
		OrderBy: orderBy,
		Desc:    sortOrder == "desc",
	}, nil
}
