package main

import (
	"go-api/controller"
	"go-api/db"
	"go-api/repository"
	"go-api/service"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	productRepository := repository.NewProductRepository(dbConnection)
	productService := service.NewProductService(productRepository)
	ProductController := controller.NewProductController(productService)

	server.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
		})
	})

	server.GET("/products", ProductController.GetAll)
	server.POST("/products", ProductController.CreateProduct)

	server.Run(":8080")
}
