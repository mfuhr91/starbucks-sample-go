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
	
	productRepository = repositories.NewProductRepository()
	productService    = services.NewProductService(productRepository)
	productController = controllers.NewProductController(productService)
	
	orderRepository = repositories.NewOrderRepository(productRepository)
	orderService    = services.NewOrderService(orderRepository)
	orderController = controllers.NewOrderController(orderService)
)

func InitRoutes(r *gin.Engine) {
	r.NoRoute(templates.NoRouteHandler)
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
	
	r.GET("/products", templates.ProductRender)
	r.POST("/products", templates.ProductRender)
	r.GET("/products/new", templates.ProductRenderForm)
	r.GET("/products/edit", templates.ProductRenderForm)
	r.GET("/products/delete-confirm", templates.DeleteConfirmRender)
	
	r.POST("/products/save", productController.Save)
	r.POST("/products/delete", productController.Delete)
	
}
