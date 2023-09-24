package service

import (
	"log"

	"github.com/dzniel/order-n-pay/services/payment/internal/dto"
	"github.com/dzniel/order-n-pay/services/payment/internal/model"
	"github.com/dzniel/order-n-pay/services/payment/internal/repository"
	"github.com/gin-gonic/gin"
)

type PaymentService struct {
	repository repository.PaymentRepository
}

type PaymentServiceInterface interface {
	CreatePayment(ctx *gin.Context, request dto.CreatePaymentRequest) (id int, err error)
	ListPayment(ctx *gin.Context) []dto.Payment
}

func NewPaymentService() PaymentServiceInterface {
	return &PaymentService{}
}

func (service *PaymentService) CreatePayment(ctx *gin.Context, request dto.CreatePaymentRequest) (id int, err error) {
	payment := model.Payment{
		OrderID:       request.OrderID,
		PaymentStatus: "PAID",
	}

	log.Println(payment)

	id, err = service.repository.CreatePayment(ctx, payment)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *PaymentService) ListPayment(ctx *gin.Context) []dto.Payment {
	return service.repository.ListPayment(ctx)
}
