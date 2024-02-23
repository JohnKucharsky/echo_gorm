package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"os"
)

type (
	CustomValidator struct {
		validator *validator.Validate
	}
)

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("Can't load env")
	}

	//port
	port := "8080"
	portEnv := os.Getenv("PORT")
	var portString = portEnv
	if portEnv == "" {
		portString = port
	}
	//end port

	//db
	dbAddress := "postgres://postgres:pass@db:5432/data?sslmode=disable"
	dbAddressEnv := os.Getenv("DB_URL")
	var dbAddressString = dbAddressEnv
	if dbAddressEnv == "" {
		dbAddressString = dbAddress
	}
	//end db

	//server setup
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}

	logger := logrus.New()
	e.Use(
		middleware.RequestLoggerWithConfig(
			middleware.RequestLoggerConfig{
				LogURI:    true,
				LogStatus: true,
				LogValuesFunc: func(
					c echo.Context,
					values middleware.RequestLoggerValues,
				) error {
					logger.WithFields(
						logrus.Fields{
							"URI":    values.URI,
							"status": values.Status,
						},
					).Info("request")

					return nil
				},
			},
		),
	)
	//end server setup

	apiGroup := e.Group("/api")
	routes(apiGroup, dbAddressString)
	e.Logger.Fatal(e.Start(":" + portString))
}
