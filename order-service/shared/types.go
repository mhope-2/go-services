package shared

import _ "github.com/go-playground/validator/v10"

type CreateOrderRequest struct {
	UserID           string  `validate:"required"`
	ProductCode      string  `validate:"required"`
	CustomerFullName string  `validate:"required"`
	ProductName      string  `validate:"required"`
	TotalAmount      float64 `validate:"required"`
	CreatedAt        string  `validate:"required"`

	// TODO: validate request in handler
	//	err := validate.Struct(req)
	//	if err != nil {
	//	// Return an error response if validation fails
	//	c.JSON(http.StatusBadRequest, gin.H{
	//	"error": "Validation failed",
	//})
	//	return
	//}
}
