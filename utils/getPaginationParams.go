package utils

import (
	"github.com/labstack/echo/v4"
	"strconv"
)

type PaginationParams struct {
	Limit int `json:"limit" validate:"numeric,gte=0,lte=200"`
	Page  int `json:"page" validate:"numeric,gte=0"`
}

func GetPaginationParams(c echo.Context) (*PaginationParams, error) {
	limitParam := c.QueryParam("limit")
	pageParam := c.QueryParam("page")
	limit, err := strconv.Atoi(limitParam)
	if limitParam != "" && err != nil {
		return nil, err
	}
	page, err := strconv.Atoi(pageParam)
	if pageParam != "" && err != nil {
		return nil, err
	}

	var paginationParams = PaginationParams{
		Limit: limit,
		Page:  page,
	}

	if err := c.Validate(paginationParams); err != nil {
		return nil, err
	}

	return &paginationParams, nil
}
