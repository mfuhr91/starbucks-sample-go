package templates

import (
	"github.com/gin-gonic/gin"
	"log"
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
	
	productRepository = repositories.NewProductRepository()
	productService    = services.NewProductService(productRepository)
	productController = controllers.NewProductController(productService)
	
	orderRepository = repositories.NewOrderRepository(productRepository)
	orderService    = services.NewOrderService(orderRepository)
	orderController = controllers.NewOrderController(orderService)
	order           models.Order
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
	// clean order
	order = models.Order{}
	
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
	productId := c.Query("productId")
	
	var product models.Product
	var err error
	if productId != "" {
		product, err = productService.GetById(productId)
	}
	
	if customerId != "" && order.ID == "" {
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
		
		order.FinalPrice = setItemsGetFinalPrice(productId, product)
		
		order.Customer = customer
		order.Time = date
		
		c.HTML(http.StatusOK, "order-form.html", gin.H{
			"Title": "Nuevo Pedido",
			"Order": order,
			/*"Item":  item,*/
		})
		return
	}
	
	id := c.Query("id")
	if c.Query("id") == "" {
		id = order.ID
	}
	
	order, err = orderService.GetById(id)
	
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	
	order.FinalPrice = setItemsGetFinalPrice(productId, product)
	
	order.Time = order.Time[0:16]
	c.HTML(http.StatusOK, "order-form.html", gin.H{
		"Title": "Editar Pedido",
		"Order": order,
	})
}

func ProductRender(c *gin.Context) {
	
	var items []models.Item
	
	err := c.ShouldBind(&items)
	if err != nil {
		log.Printf("error when binding form - %s", err)
		return
	}
	var finalPrice float64
	
	for _, item := range items {
		if contains, index := containsProduct(order.Items, item.ProductID); contains {
			order.Items[index].Quantity = item.Quantity
			
			quantity := float64(item.Quantity)
			finalPrice += quantity * item.Price
		} else {
			order.Items = append(order.Items, item)
		}
	}
	order.FinalPrice = finalPrice
	
	searching, _ := strconv.ParseBool(c.Query("search"))
	if searching {
		products, err := productService.GetAll(false)
		if err != nil {
			c.HTML(http.StatusOK, "error.html", gin.H{
				"Error": err,
				"Msg":   "Toca el botón para reintentar",
			})
			
			return
		}
		
		customerId := c.Query("customerId")
		c.HTML(http.StatusOK, "products.html", gin.H{
			"Title":      "Seleccionar Producto",
			"Products":   products,
			"CustomerID": customerId,
			"Searching":  searching,
		})
		
		return
	}
	
	products, err := productService.GetAll(true)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		
		return
	}
	
	c.HTML(http.StatusOK, "products.html", gin.H{
		"Title":    "Productos",
		"Products": products,
	})
	
}

func ProductRenderForm(c *gin.Context) {
	id := c.Query("id")
	if id == "" {
		c.HTML(http.StatusOK, "product-form.html", gin.H{
			"Title": "Nuevo Producto",
		})
		return
	}
	product, err := productService.GetById(id)
	if err != nil {
		c.HTML(http.StatusOK, "error.html", gin.H{
			"Error": err,
			"Msg":   "Toca el botón para reintentar",
		})
		return
	}
	c.HTML(http.StatusOK, "product-form.html", gin.H{
		"Title":   "Editar Producto",
		"Product": product,
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
	} else if from == "products" {
		c.HTML(http.StatusOK, "delete-confirm.html", gin.H{
			"Title": "Eliminar Producto",
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

func containsProduct(items []models.Item, productId string) (bool, int) {
	for i, p := range items {
		if p.ProductID == productId {
			return true, i
		}
	}
	
	return false, 0
}

func setItemsGetFinalPrice(productId string, product models.Product) float64 {
	var item models.Item
	if productId != "" {
		item = models.Item{
			ProductID: product.ID,
			Product:   product,
			Quantity:  1,
		}
		
		if contains, _ := containsProduct(order.Items, productId); !contains {
			order.Items = append(order.Items, item)
		}
	}
	var finalPrice float64
	for _, item = range order.Items {
		
		quantity := float64(item.Quantity)
		finalPrice += quantity * item.Price
	}
	
	return finalPrice
}
