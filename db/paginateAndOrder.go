package db

import (
	"github.com/JohnKucharsky/echo_gorm/utils"
	"gorm.io/gorm"
)

type Pagination struct {
	Limit     int   `json:"limit"`
	Page      int   `json:"page"`
	TotalRows int64 `json:"total_rows"`
}

func PaginateAndOrder(
	value interface{},
	pagination *Pagination,
	db *gorm.DB,
	pp utils.PaginationParams,
) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	var limit = pp.Limit
	if limit < 20 {
		limit = 20
	}

	pagination.TotalRows = totalRows
	pagination.Page = pp.Page
	pagination.Limit = limit

	return func(db *gorm.DB) *gorm.DB {
		offset := (pp.Page - 1) * limit

		return db.Offset(offset).Limit(limit)
	}
}

type PaginatedResponse struct {
	Data       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}

func PaginationRes(
	user interface{},
	pagination Pagination,
) PaginatedResponse {
	return PaginatedResponse{
		Data:       user,
		Pagination: pagination,
	}
}
