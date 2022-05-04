package routes

import (
	"github.com/gin-gonic/gin"
	"starbucks-app/controllers"
	"starbucks-app/repositories"
	"starbucks-app/services"
	"starbucks-app/templates"
)

var (
	customerRepository = repositories.NewCustomerRepository()
	customerService    = services.NewCustomerService(customerRepository)
	customerController = controllers.NewCustomerController(customerService)
	orderController    = controllers.NewOrderController()
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", templates.OrdersRender)
	router.GET("/ping", templates.Ping)
	
	router.Static("assets", "resources/assets")
	
	router.LoadHTMLGlob("resources/templates/*")
	router.GET("/customers", templates.CustomersRender)
	router.GET("/customers/new", templates.CustomerRenderForm)
	router.GET("/customers/edit", templates.CustomerRenderForm)
	
	router.GET("/orders", templates.OrdersRender)
	
	router.POST("/customer/save", customerController.Save)
}
