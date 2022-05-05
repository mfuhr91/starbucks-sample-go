package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"starbucks-app/models"
	"starbucks-app/services"
	"starbucks-app/utils"
)

type customerController struct{}

type CustomerController interface {
	Save(c *gin.Context)
	Delete(c *gin.Context)
}

var (
	customerService services.CustomerService
)

func NewCustomerController(service services.CustomerService) CustomerController {
	customerService = service
	return &customerController{}
}

func (cont *customerController) Save(c *gin.Context) {
	var customer models.Customer
	
	err := c.ShouldBind(&customer)
	if err != nil {
		log.Printf("error when binding form - %s", err)
		return
	}
	
	_, err = customerService.Save(customer)
	if err != nil {
		log.Printf("error when saving customer - %s", err)
		return
	}
	
	utils.Redirect("/customers", c)
}

func (cont *customerController) Delete(c *gin.Context) {
	id := c.Query("id")
	err := customerService.Delete(id)
	if err != nil {
		log.Printf("error when deleting customer - %s", err)
		return
	}
	
	utils.Redirect("/customers", c)
}
