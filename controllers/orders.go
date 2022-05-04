package controllers

import (
	"github.com/gin-gonic/gin"
	"starbucks-app/utils"
)

type pedidoController struct{}

type PedidoController interface {
	GetAll(c *gin.Context)
	Save(c *gin.Context)
}

func NewOrderController() PedidoController {
	return &pedidoController{}
}

func (cont *pedidoController) GetAll(c *gin.Context) {

}

func (cont *pedidoController) Save(c *gin.Context) {
	
	utils.Redirect("/orders", c)
}
