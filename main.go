package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"starbucks-app/routes"
)

func main() {
	router := gin.Default()
	
	routes.InitRoutes(router)
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = ":8080"
	}
	err := router.Run(port)
	if err != nil {
		log.Fatalf("Cannot start the server: %v ", err.Error())
		return
	}
}
