package repository

import (
	"fmt"
	"log"

	"github.com/dzniel/order-n-pay/services/payment/internal/dto"
	"github.com/dzniel/order-n-pay/services/payment/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaymentRepository struct {
	db *gorm.DB
}

func NewPaymentRepository(db *gorm.DB) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (repo *PaymentRepository) CreatePayment(ctx *gin.Context, payment model.Payment) (id int, err error) {
	db := ctx.MustGet("db").(*gorm.DB)
	qx := db.Save(&payment)
	if qx.Error != nil {
		log.Println("error:", qx.Error)
		err = fmt.Errorf(qx.Error.Error())
		return
	}

	return payment.ID, nil
}

func (repo *PaymentRepository) ListPayment(ctx *gin.Context) (payments []dto.Payment) {
	db := ctx.MustGet("db").(*gorm.DB)
	db.Model(model.Payment{}).Find(&payments)
	return
}
