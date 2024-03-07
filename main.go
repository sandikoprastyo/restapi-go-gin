package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inisialisasi router Gin
	router := gin.Default()

	// Definisikan rute
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to my microservice!",
		})
	})

	// Customer routes
	router.GET("/customers", GetCustomers)
	router.GET("/customers/:id", GetCustomersById)
	router.POST("/customers", CreateCustomer)
	router.PUT("/customers/:id", UpdateCustomer)
	router.DELETE("/customers/:id", DeleteCustomer)

	// Product routes
	router.GET("/products", GetProducts)
	router.GET("/products/:id", GetProductsById)
	router.POST("/products", CreateProduct)
	router.PUT("/products/:id", UpdateProduct)
	router.DELETE("/products/:id", DeleteProduct)

	// Mulai server pada port tertentu
	router.Run(":8080")
}
