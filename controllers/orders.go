package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"starbucks-app/models"
	"starbucks-app/services"
	"starbucks-app/utils"
)

type orderController struct{}

type OrderController interface {
	Save(c *gin.Context)
	Delete(c *gin.Context)
}

var (
	orderService services.OrderService
)

func NewOrderController(service services.OrderService) OrderController {
	orderService = service
	return &orderController{}
}

func (cont *orderController) Save(c *gin.Context) {
	var order models.Order
	
	err := c.ShouldBind(&order)
	if err != nil {
		log.Printf("error when binding form - %s", err)
		return
	}
	
	customerId := c.Request.FormValue("customerId")
	order.Customer.ID = customerId
	
	_, err = orderService.Save(order)
	if err != nil {
		log.Printf("error when saving order - %s", err)
		return
	}
	
	utils.Redirect("/orders", c)
}

func (cont *orderController) Delete(c *gin.Context) {
	id := c.Query("id")
	err := orderService.Delete(id)
	if err != nil {
		log.Printf("error when deleting order - %s", err)
		return
	}
	
	utils.Redirect("/orders", c)
}
