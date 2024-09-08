package main

import (
	"go-api/controller"

	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	ProductController := controller.NewProductControler()

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/products", ProductController.GetProducts)
	server.Run(":8000")
}
