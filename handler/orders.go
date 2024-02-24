package handler

import (
	"github.com/JohnKucharsky/echo_gorm/db"
	"github.com/JohnKucharsky/echo_gorm/models"
	"github.com/JohnKucharsky/echo_gorm/serializer"
	"github.com/JohnKucharsky/echo_gorm/utils"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm/clause"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) CreateOrder(c echo.Context) error {
	var orderBody serializer.OrderBody
	if err := c.Bind(&orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(orderBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var order = serializer.OrderBodyToOrder(orderBody)
	result := apiConfig.Database.DB.Model(&order).Preload("User").Preload("Product").Clauses(clause.Returning{}).Create(&order).First(&order)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusConflict, result.Error.Error())
	}

	return c.JSON(
		http.StatusCreated, order,
	)
}

func (apiConfig *DatabaseController) GetOrders(c echo.Context) error {
	paginationParams, err := utils.GetPaginationParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var orders []models.Order
	var pagination db.Pagination

	apiConfig.Database.DB.Preload("User").Preload("Product").Scopes(
		db.PaginateAndOrder(
			orders,
			&pagination,
			apiConfig.Database.DB,
			*paginationParams,
		),
	).Find(&orders)

	return c.JSON(
		http.StatusCreated, db.PaginationRes(orders, pagination),
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

	var order = serializer.OrderBodyToOrder(orderBody)

	apiConfig.Database.DB.Model(&order).Preload(clause.Associations).Where(
		"id = ?",
		dbId,
	).Updates(&order).First(&order)

	if order.ID == 0 {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(
		http.StatusCreated, order,
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

	var order models.Order
	apiConfig.Database.DB.Clauses(clause.Returning{}).Delete(&order, dbId)

	if order.ID == 0 {
		return c.NoContent(http.StatusOK)
	}

	return c.JSON(http.StatusOK, order)
}
