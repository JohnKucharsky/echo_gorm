package serializer

type UserBody struct {
	FirstName string  `json:"first_name" validate:"required"`
	LastName  *string `json:"last_name"`
}
