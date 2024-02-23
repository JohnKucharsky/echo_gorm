package serializer

type OrderBody struct {
	ProductId int32 `json:"product_id" validate:"required"`
	UserId    int32 `json:"user_id" validate:"required"`
}
