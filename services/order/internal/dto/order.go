package dto

type Order struct {
	ID            int    `json:"id"`
	CustomerID    int    `json:"customer_id"`
	Items         string `json:"items"`
	Price         int    `json:"price"`
	PaymentStatus string `json:"payment_status"`
}

type CreateOrderRequest struct {
	CustomerID int    `json:"customer_id"`
	Items      string `json:"items"`
	Price      int    `json:"price"`
}

type CreateOrderResponse struct {
	ID int `json:"id"`
}

type PaidRequest struct {
	OrderID int `json:"order_id"`
}
