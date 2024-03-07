package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Customer struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

var customers []Customer

func init() {
	// Data dummy pelanggan awal
	customers = append(customers, Customer{ID: "1", Name: "John Doe", Email: "john@example.com"})
	customers = append(customers, Customer{ID: "2", Name: "Jane Smith", Email: "jane@example.com"})
}

func GetCustomers(c *gin.Context) {
	c.JSON(http.StatusOK, customers)
}

func GetCustomersById(c *gin.Context) {
	id := c.Param("id")

	for _, customer := range customers {
		if customer.ID == id {
			// Jika ditemukan, kirimkan data pelanggan dalam format JSON
			c.JSON(http.StatusOK, customer)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}

func CreateCustomer(c *gin.Context) {
	var newCustomer Customer
	if err := c.BindJSON(&newCustomer); err != nil {
		return
	}
	customers = append(customers, newCustomer)
	c.JSON(http.StatusCreated, newCustomer)
}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var updatedCustomer Customer

	if err := c.BindJSON(&updatedCustomer); err != nil {
		return
	}

	for i, a := range customers {
		if a.ID == id {
			customers[i] = updatedCustomer
			c.JSON(http.StatusOK, updatedCustomer)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}

func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	for i, a := range customers {
		if a.ID == id {
			customers = append(customers[:i], customers[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "customer deleted"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"message": "customer not found"})
}
