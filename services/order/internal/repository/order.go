package repository

import (
	"fmt"
	"log"

	"github.com/dzniel/order-n-pay/services/order/internal/dto"
	"github.com/dzniel/order-n-pay/services/order/internal/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db: db}
}

func (repo *OrderRepository) CreateOrder(ctx *gin.Context, order model.Order) (id int, err error) {
	db := ctx.MustGet("db").(*gorm.DB)
	qx := db.Save(&order)
	if qx.Error != nil {
		log.Println("error:", qx.Error)
		err = fmt.Errorf(qx.Error.Error())
		return
	}

	return order.ID, nil
}

func (repo *OrderRepository) ListOrder(ctx *gin.Context) (orders []dto.Order) {
	db := ctx.MustGet("db").(*gorm.DB)
	db.Model(model.Order{}).Find(&orders)
	return orders
}
