package controller

import (
	"net/http"

	"github.com/dzniel/order-n-pay/services/order/internal/dto"
	"github.com/dzniel/order-n-pay/services/order/internal/service"
	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context) {
	var request dto.CreateOrderRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.NewOrderService().CreateOrder(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := gin.H{
		"result": &dto.CreateOrderResponse{ID: id},
	}
	ctx.JSON(http.StatusOK, result)
}

func ListOrder(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.NewOrderService().ListOrder(ctx))
}
