package serializer

type ProductBody struct {
	Name   string  `json:"name" validate:"required"`
	Serial *string `json:"serial"`
}
