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
	orderRepository    = repositories.NewOrderRepository()
	orderService       = services.NewOrderService(orderRepository)
	orderController    = controllers.NewOrderController(orderService)
)

func InitRoutes(router *gin.Engine) {
	router.GET("/", templates.OrdersRender)
	router.GET("/ping", templates.Ping)
	
	router.Static("assets", "resources/assets")
	
	router.LoadHTMLGlob("resources/templates/*")
	
	router.GET("/customers", templates.CustomersRender)
	router.GET("/customers/new", templates.CustomerRenderForm)
	router.GET("/customers/edit", templates.CustomerRenderForm)
	router.GET("/customers/delete-confirm", templates.DeleteConfirmRender)
	
	router.POST("/customers/save", customerController.Save)
	router.POST("/customers/delete", customerController.Delete)
	
	router.GET("/orders", templates.OrdersRender)
	router.GET("/orders/new", templates.OrderRenderForm)
	router.GET("/orders/edit", templates.OrderRenderForm)
	router.GET("/orders/delete-confirm", templates.DeleteConfirmRender)
	
	router.POST("/orders/save", orderController.Save)
	router.POST("/orders/delete", orderController.Delete)
	
}
