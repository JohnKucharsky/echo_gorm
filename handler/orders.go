package handler

import (
	"github.com/JohnKucharsky/echo_gorm/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) OrderPost(c echo.Context) error {
	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		orderBody,
	)
}

func (apiConfig *DatabaseController) GetOrders(c echo.Context) error {

	return c.String(
		http.StatusOK,
		"",
	)
}

func (apiConfig *DatabaseController) UpdateOrder(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int32
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int32(res)

	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusCreated,
		struct {
			serializer.OrderBody
			dbId int32
		}{orderBody, dbId},
	)
}

func (apiConfig *DatabaseController) DeleteOrder(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int32
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int32(res)

	return c.String(http.StatusOK, string(dbId))
}
