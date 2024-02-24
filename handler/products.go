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

func (apiConfig *DatabaseController) CreateProduct(c echo.Context) error {
	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var product = serializer.ProductBodyToProduct(productBody)
	result := apiConfig.Database.DB.Model(&product).Clauses(clause.Returning{}).Create(&product)
	if result.Error != nil {
		return echo.NewHTTPError(http.StatusConflict, result.Error.Error())
	}

	return c.JSON(
		http.StatusCreated,
		product,
	)
}

func (apiConfig *DatabaseController) GetProducts(c echo.Context) error {
	paginationParams, err := utils.GetPaginationParams(c)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	orderClause, err := utils.GetOrderParams(
		c,
		[]string{
			"created_at",
			"updated_at",
			"name",
			"serial_number",
		},
	)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var orderColumn = "updated_at"
	if orderClause.OrderBy != "" {
		orderColumn = orderClause.OrderBy
	}

	nameQuery := c.QueryParam("name")
	serialNumberQuery := c.QueryParam("serial_number")
	likeKey := "name like ? AND serial_number like ?"

	var products []models.Product
	var pagination db.Pagination

	apiConfig.Database.DB.Where(
		likeKey,
		"%"+nameQuery+"%",
		"%"+serialNumberQuery+"%",
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
			products,
			&pagination,
			apiConfig.Database.DB,
			*paginationParams,
		),
	).Find(&products)

	return c.JSON(
		http.StatusCreated, db.PaginationRes(products, pagination),
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

	var product models.Product
	apiConfig.Database.DB.First(&product, dbId)

	if product.ID == 0 {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(
		http.StatusOK, product,
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

	var productBody serializer.ProductBody
	if err := c.Bind(&productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	if err := c.Validate(productBody); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var product = serializer.ProductBodyToProduct(productBody)

	apiConfig.Database.DB.Model(&product).Clauses(clause.Returning{}).Where(
		"id = ?",
		dbId,
	).Updates(&product)

	if product.ID == 0 {
		return c.NoContent(http.StatusOK)
	}
	return c.JSON(
		http.StatusCreated, product,
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

	var product models.Product
	apiConfig.Database.DB.Clauses(clause.Returning{}).Delete(&product, dbId)

	if product.ID == 0 {
		return c.NoContent(http.StatusOK)
	}

	return c.JSON(http.StatusOK, product)
}
