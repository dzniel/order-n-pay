package dto

type Payment struct {
	ID            int    `json:"id"`
	OrderID       int    `json:"order_id"`
	PaymentStatus string `json:"payment_status"`
}

type CreatePaymentRequest struct {
	OrderID int `json:"order_id"`
}
