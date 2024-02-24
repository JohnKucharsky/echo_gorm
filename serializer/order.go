package serializer

import "github.com/JohnKucharsky/echo_gorm/models"

type OrderBody struct {
	ProductId int `json:"product_id" validate:"required"`
	UserId    int `json:"user_id" validate:"required"`
}

func OrderBodyToOrder(orderBody OrderBody) models.Order {
	return models.Order{
		ProductID: orderBody.ProductId,
		UserID:    orderBody.UserId,
	}
}
