package main

import (
	"log"
	"net/http"

	"github.com/dzniel/order-n-pay/services/payment/internal/controller"
	"github.com/dzniel/order-n-pay/services/payment/pkg/database"
	"github.com/dzniel/order-n-pay/services/payment/pkg/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := helper.LoadConfig("./services/payment")
	if err != nil {
		log.Fatal("load config failed")
		return
	}

	db, err := database.New().DBConnect(config)
	if err != nil {
		log.Fatal("database connection failed")
		return
	}

	app := gin.Default()
	app.Use(gin.Recovery())
	app.Use(database.InjectDB(db))
	app.UseRawPath = true
	app.UnescapePathValues = true

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "DELETE"}

	app.Use(cors.New(corsConfig))

	app.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	v1 := app.Group("/v1")
	{
		v1.POST("/payment", controller.CreatePayment)
		v1.GET("/payment", controller.ListPayment)
	}

	app.Run(":" + config.ServicePort)
}
