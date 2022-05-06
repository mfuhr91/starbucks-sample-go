package templates

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"starbucks-app/controllers"
	"starbucks-app/models"
	"starbucks-app/repositories"
	"starbucks-app/services"
	"strconv"
	"strings"
	"time"
)

var (
	customerRepository = repositories.NewCustomerRepository()
	customerService    = services.NewCustomerService(customerRepository)
	customerController = controllers.NewCustomerController(customerService)
	orderRepository    = repositories.NewOrderRepository()
	orderService       = services.NewOrderService(orderRepository)
	orderController    = controllers.NewOrderController(orderService)
)

func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"ping": "pong"})
}

func CustomersRender(c *gin.Context) {
	customers, err := customerService.GetAll()
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	
	searching, _ := strconv.ParseBool(c.Query("search"))
	if searching {
		c.HTML(http.StatusOK, "customers.html", gin.H{
			"Title":     "Seleccionar cliente",
			"Customers": customers,
			"Searching": searching,
		})
		return
	}
	
	c.HTML(http.StatusOK, "customers.html", gin.H{
		"Title":     "Clientes",
		"Customers": customers,
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
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	c.HTML(http.StatusOK, "customer-form.html", gin.H{
		"Title":    "Editar Cliente",
		"Customer": customer,
	})
}

func OrdersRender(c *gin.Context) {
	orders, err := orderService.GetAll()
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	
	c.HTML(http.StatusOK, "orders.html", gin.H{
		"Title":  "Pedidos",
		"Orders": orders,
	})
}

func OrderRenderForm(c *gin.Context) {
	customerId := c.Query("customerId")
	
	if customerId != "" {
		customer, err := customerService.GetById(customerId)
		if err != nil {
			c.HTML(http.StatusOK, "error.html", gin.H{
				"Error": err,
				"Msg":   "Toca el botón para reintentar",
			})
			return
		}
		
		loc, _ := time.LoadLocation("America/Argentina/Buenos_Aires")
		newDate := time.Now().In(loc)
		date := newDate.String()[0:16]
		date = strings.Replace(date, " ", "T", 1)
		
		order := models.Order{
			Customer: customer,
			Time:     date,
		}
		c.HTML(http.StatusOK, "order-form.html", gin.H{
			"Title": "Nuevo Pedido",
			"Order": order,
		})
		return
	}
	
	id := c.Query("id")
	order, err := orderService.GetById(id)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	
	order.Time = order.Time[0:16]
	c.HTML(http.StatusOK, "order-form.html", gin.H{
		"Title": "Editar Pedido",
		"Order": order,
	})
}

func DeleteConfirmRender(c *gin.Context) {
	id := c.Query("id")
	from := c.Query("from")
	if from == "orders" {
		c.HTML(http.StatusOK, "delete-confirm.html", gin.H{
			"Title": "Eliminar Pedido",
			"From":  from,
			"ID":    id,
		})
		return
	}
	
	c.HTML(http.StatusOK, "delete-confirm.html", gin.H{
		"Title": "Eliminar Cliente",
		"From":  from,
		"ID":    id,
	})
	
}
