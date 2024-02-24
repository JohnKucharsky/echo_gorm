package serializer

import (
	"github.com/JohnKucharsky/echo_gorm/models"
)

type ProductBody struct {
	Name         string `json:"name" validate:"required"`
	SerialNumber string `json:"serial_number" validate:"required"`
}

func ProductBodyToProduct(productBody ProductBody) models.Product {
	return models.Product{
		Name:         productBody.Name,
		SerialNumber: productBody.SerialNumber,
	}
}
