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

func (apiConfig *DatabaseController) CreateUser(c echo.Context) error {
	var userBody serializer.UserBody
	if err := c.Bind(&userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(userBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var user = serializer.UserBodyToUser(userBody)
	result := apiConfig.Database.DB.Model(&user).Clauses(clause.Returning{}).Create(&user)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusConflict, result.Error.Error())
	}

	return c.JSON(
		http.StatusCreated, user,
	)
}

func (apiConfig *DatabaseController) GetUsers(c echo.Context) error {
	paginationParams, err := utils.GetPaginationParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	orderClause, err := utils.GetOrderParams(
		c,
		[]string{"created_at", "first_name", "last_name"},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var orderColumn = "created_at"
	if orderClause.OrderBy != "" {
		orderColumn = orderClause.OrderBy
	}

	firstNameQuery := c.QueryParam("first_name")
	lastNameQuery := c.QueryParam("last_name")
	likeKey := "first_name like ? AND last_name is null or last_name like ?"

	var users []models.User
	var pagination db.Pagination

	apiConfig.Database.DB.Where(
		likeKey,
		"%"+firstNameQuery+"%",
		"%"+lastNameQuery+"%",
	).Clauses(
		//clause.Gt{
		//	Column: "created_at",
		//	Value:  "2024-02-23 19:29:46.316036 +00:00",
		//},
		clause.OrderBy{
			Columns: []clause.OrderByColumn{
				{
					Column: clause.Column{Name: orderColumn},
					Desc:   orderClause.Desc,
				},
			},
		},
	).Scopes(
		db.PaginateAndOrder(
			users,
			&pagination,
			apiConfig.Database.DB,
			*paginationParams,
		),
	).Find(&users)

	return c.JSON(
		http.StatusCreated, db.PaginationRes(users, pagination),
	)
}

func (apiConfig *DatabaseController) GetOneUser(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int(res)

	var user models.User
	apiConfig.Database.DB.First(&user, dbId)

	if user.ID == 0 {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(
		http.StatusOK, user,
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

	var user = serializer.UserBodyToUser(userBody)

	apiConfig.Database.DB.Model(&user).Clauses(clause.Returning{}).Where(
		"id = ?",
		dbId,
	).Updates(&user)

	if user.ID == 0 {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(
		http.StatusCreated, user,
	)
}

func (apiConfig *DatabaseController) DeleteUser(c echo.Context) error {
	var id = c.Param("id")
	if id == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "No id in the address")
	}
	var dbId int
	res, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	dbId = int(res)

	var user models.User
	apiConfig.Database.DB.Clauses(clause.Returning{}).Delete(&user, dbId)

	if user.ID == 0 {
		return c.NoContent(http.StatusOK)
	}

	return c.JSON(http.StatusOK, user)
}
