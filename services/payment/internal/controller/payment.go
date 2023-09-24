package controller

import (
	"net/http"

	"github.com/dzniel/order-n-pay/services/payment/internal/dto"
	"github.com/dzniel/order-n-pay/services/payment/internal/service"
	"github.com/gin-gonic/gin"
)

func CreatePayment(ctx *gin.Context) {
	var request dto.CreatePaymentRequest
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := service.NewPaymentService().CreatePayment(ctx, request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	result := gin.H{
		"result": &dto.CreatePaymentRequest{OrderID: id},
	}
	ctx.JSON(http.StatusOK, result)
}

func ListPayment(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, service.NewPaymentService().ListPayment(ctx))
}
