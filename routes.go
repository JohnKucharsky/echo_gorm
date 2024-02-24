package main

import (
	"github.com/JohnKucharsky/echo_gorm/db"
	"github.com/JohnKucharsky/echo_gorm/handler"
	"github.com/labstack/echo/v4"
)

func routes(route *echo.Group, dbConnectionString string) {
	database := db.DatabaseConnection(dbConnectionString)
	h := handler.DatabaseController{Database: database}

	route.GET("/healthz", handler.CheckHealth)

	// users
	route.POST("/users", h.CreateUser)
	route.GET("/users", h.GetUsers)
	route.GET("/users/:id", h.GetOneUser)
	route.PUT("/users/:id", h.UpdateUser)
	route.DELETE("/users/:id", h.DeleteUser)
	// end users

	// products
	route.POST("/products", h.CreateProduct)
	route.GET("/products", h.GetProducts)
	route.GET("/products/:id", h.GetOneProduct)
	route.PUT("/products/:id", h.UpdateProduct)
	route.DELETE("/products/:id", h.DeleteProduct)
	// end products

	// orders
	route.POST("/orders", h.CreateOrder)
	route.GET("/orders", h.GetOrders)
	route.PUT("/orders/:id", h.UpdateOrder)
	route.DELETE("/orders/:id", h.DeleteOrder)
	// end orders
}
