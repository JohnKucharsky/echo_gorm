package handler

import "github.com/JohnKucharsky/echo_gorm/db"

type DatabaseController struct {
	Database *db.ApiConfig
}
