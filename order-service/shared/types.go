package shared

type CreateOrderRequest struct {
	UserID           string  `validate:"required"`
	ProductCode      string  `validate:"required"`
	CustomerFullName string  `validate:"required"`
	ProductName      string  `validate:"required"`
	TotalAmount      float64 `validate:"required"`
	CreatedAt        string  `validate:"required"`
}
