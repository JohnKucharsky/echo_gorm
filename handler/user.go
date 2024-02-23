package handler

import (
	"github.com/JohnKucharsky/echo_gorm/serializer"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (apiConfig *DatabaseController) UserPost(c echo.Context) error {

	return c.JSON(
		http.StatusCreated, struct {
		}{},
	)
}

func (apiConfig *DatabaseController) GetUsers(c echo.Context) error {

	return c.JSON(
		http.StatusCreated, struct {
		}{},
	)
}

func (apiConfig *DatabaseController) GetOneUser(c echo.Context) error {
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
		http.StatusCreated, struct {
			dbId int32
		}{dbId: dbId},
	)
}

func (apiConfig *DatabaseController) UpdateUser(c echo.Context) error {
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

	var userBody serializer.UserBody
	if err := c.Bind(&userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(
		http.StatusCreated, struct {
			dbId int32
		}{dbId},
	)
}

func (apiConfig *DatabaseController) DeleteUser(c echo.Context) error {
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
