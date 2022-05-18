package controllers

import (
	"github.com/gin-gonic/gin"
	"log"
	"starbucks-app/models"
	"starbucks-app/services"
	"starbucks-app/utils"
)

type productController struct{}

type ProductController interface {
	Save(c *gin.Context)
	Delete(c *gin.Context)
}

var (
	productService services.ProductService
)

func NewProductController(service services.ProductService) ProductController {
	productService = service
	return &productController{}
}

func (cont *productController) Save(c *gin.Context) {
	var product models.Product
	
	err := c.ShouldBind(&product)
	if err != nil {
		log.Printf("error when binding form - %s", err)
		return
	}
	
	_, err = productService.Save(product)
	if err != nil {
		log.Printf("error when saving product - %s", err)
		return
	}
	
	utils.Redirect("/products", c)
}

func (cont *productController) Delete(c *gin.Context) {
	id := c.Query("id")
	err := productService.Delete(id)
	if err != nil {
		log.Printf("error when deleting product - %s", err)
		return
	}
	
	utils.Redirect("/products", c)
}
