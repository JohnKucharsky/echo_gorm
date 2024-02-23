package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "I'm good, i'm running")
}
