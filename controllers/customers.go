package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"starbucks-app/models"
	"starbucks-app/services"
	"starbucks-app/utils"
)

type customerController struct{}

type CustomerController interface {
	GetAll(c *gin.Context)
	Save(c *gin.Context)
}

var (
	customerService services.CustomerService
)

func NewCustomerController(service services.CustomerService) CustomerController {
	customerService = service
	return &customerController{}
}

func (cont *customerController) GetAll(c *gin.Context) {
	customers, err := customerService.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	
	c.JSON(http.StatusOK, customers)
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
