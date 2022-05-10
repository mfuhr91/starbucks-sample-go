package router

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

func InitRoutes(r *gin.Engine) {
	r.NoRoute(templates.OrdersRender)
	r.GET("/ping", templates.Ping)
	
	r.Static("assets", "resources/assets")
	
	r.LoadHTMLGlob("resources/templates/*")
	
	r.GET("/customers", templates.CustomersRender)
	r.GET("/customers/new", templates.CustomerRenderForm)
	r.GET("/customers/edit", templates.CustomerRenderForm)
	r.GET("/customers/delete-confirm", templates.DeleteConfirmRender)
	
	r.POST("/customers/save", customerController.Save)
	r.POST("/customers/delete", customerController.Delete)
	
	r.GET("/orders", templates.OrdersRender)
	r.GET("/orders/new", templates.OrderRenderForm)
	r.GET("/orders/edit", templates.OrderRenderForm)
	r.GET("/orders/delete-confirm", templates.DeleteConfirmRender)
	
	r.POST("/orders/save", orderController.Save)
	r.POST("/orders/delete", orderController.Delete)
	
}
