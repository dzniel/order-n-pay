package main

import (
	"log"
	"net/http"

	"github.com/dzniel/order-n-pay/services/order/internal/controller"
	"github.com/dzniel/order-n-pay/services/order/pkg/database"
	"github.com/dzniel/order-n-pay/services/order/pkg/helper"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	config, err := helper.LoadConfig("./services/order")
	if err != nil {
		log.Fatal("load config failed")
		return
	}

	db, err := database.DBConnect(config)
	if err != nil {
		log.Fatal("database connection failed")
		return
	}

	app := gin.Default()
	app.Use(gin.Logger())
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
		v1.GET("/order", controller.ListOrder)
		v1.POST("/order", controller.CreateOrder)
	}

	app.Run(":" + config.Port)
}
