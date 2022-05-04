package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"starbucks-app/routes"
)

func main() {
	router := gin.Default()
	
	routes.InitRoutes(router)
	
	err := router.Run(":8080")
	if err != nil {
		log.Fatalf("Cannot start the server: %v ", err.Error())
		return
	}
}
