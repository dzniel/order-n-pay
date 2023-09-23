package model

type Order struct {
	ID            int
	CustomerID    int
	Items         string
	Price         int
	PaymentStatus string
}
