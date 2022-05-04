package templates

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starbucks-app/controllers"
	"starbucks-app/models"
	"starbucks-app/repositories"
	"starbucks-app/services"
)

var (
	customerRepository = repositories.NewCustomerRepository()
	customerService    = services.NewCustomerService(customerRepository)
	customerController = controllers.NewCustomerController(customerService)
	orderController    = controllers.NewOrderController()
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func CustomersRender(c *gin.Context) {
	customers, err := customerService.GetAll()
	if err != nil {
		c.HTML(http.StatusOK, "customers.html", gin.H{
			"Title":     "Clientes",
			"Customers": []models.Customer{},
			"Status":    "Empty",
		})
		return
	}
	c.HTML(http.StatusOK, "customers.html", gin.H{
		"Title":     "Clientes",
		"Customers": customers,
		"Status":    "OK",
	})
}

func CustomerRenderForm(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.HTML(http.StatusOK, "customer-form.html", gin.H{
			"Title": "Nuevo Cliente",
		})
		return
	}
	customer, err := customerService.GetById(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err,
		})
		return
	}
	c.HTML(http.StatusOK, "customer-form.html", gin.H{
		"Title":    "Editar Cliente",
		"Customer": customer,
	})
}

func OrdersRender(c *gin.Context) {
	
	c.HTML(http.StatusOK, "orders.html", gin.H{
		"Title": "Pedidos",
	})
}
