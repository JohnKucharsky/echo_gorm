package handler

import (
	"github.com/JohnKucharsky/echo_gorm/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) ProductPost(c echo.Context) error {
	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.String(
		http.StatusCreated,
		"",
	)
}

func (apiConfig *DatabaseController) GetProducts(c echo.Context) error {

	return c.String(
		http.StatusOK,
		"",
	)
}

func (apiConfig *DatabaseController) GetOneProduct(c echo.Context) error {
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

	return c.JSON(
		http.StatusOK, struct {
			dbId int32
		}{dbId},
	)
}

func (apiConfig *DatabaseController) UpdateProduct(c echo.Context) error {
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

	return c.String(
		http.StatusCreated,
		string(dbId),
	)
}

func (apiConfig *DatabaseController) DeleteProduct(c echo.Context) error {
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

	return c.JSON(
		http.StatusOK, struct {
			dbId int32
		}{dbId},
	)
}
