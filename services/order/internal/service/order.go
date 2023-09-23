package service

import (
	"log"

	"github.com/dzniel/order-n-pay/services/order/internal/dto"
	"github.com/dzniel/order-n-pay/services/order/internal/model"
	"github.com/dzniel/order-n-pay/services/order/internal/repository"
	"github.com/gin-gonic/gin"
)

type OrderService struct {
	repository *repository.OrderRepository
}

type OrderServiceInterface interface {
	CreateOrder(ctx *gin.Context, request dto.CreateOrderRequest) (id int, err error)
	ListOrder(ctx *gin.Context) []dto.Order
}

func NewOrderService() OrderServiceInterface {
	return &OrderService{}
}

func (service *OrderService) CreateOrder(ctx *gin.Context, request dto.CreateOrderRequest) (id int, err error) {
	order := model.Order{
		CustomerID: request.CustomerID,
		Items:      request.Items,
		Price:      request.Price,
	}

	log.Println(order)

	id, err = service.repository.CreateOrder(ctx, order)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *OrderService) ListOrder(ctx *gin.Context) []dto.Order {
	return service.repository.ListOrder(ctx)
}
