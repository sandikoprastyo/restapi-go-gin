package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Product adalah struktur data yang merepresentasikan produk
type Product struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// Slice products untuk menyimpan daftar produk
var products []Product

func init() {
	products = append(products, Product{ID: "1", Name: "John Doe", Price: 12000})
	products = append(products, Product{ID: "2", Name: "Jane Smith", Price: 10000})
}

// Get all products
func GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, products)
}

func GetProductsById(c *gin.Context) {
	id := c.Param("id")

	for _, products := range products {
		if products.ID == id {
			// Jika ditemukan, kirimkan data pelanggan dalam format JSON
			c.JSON(http.StatusOK, products)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// Create a new product
func CreateProduct(c *gin.Context) {
	var newProduct Product
	if err := c.BindJSON(&newProduct); err != nil {
		return
	}
	products = append(products, newProduct)
	c.JSON(http.StatusCreated, newProduct)
}

// Update an existing product
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var updatedProduct Product

	if err := c.BindJSON(&updatedProduct); err != nil {
		return
	}

	for i, p := range products {
		if p.ID == id {
			products[i] = updatedProduct
			c.JSON(http.StatusOK, updatedProduct)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}

// Delete a product
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	for i, p := range products {
		if p.ID == id {
			products = append(products[:i], products[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "product not found"})
}
