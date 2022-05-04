package main

import (
	"fmt"
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
		port = "4747"
	}
	log.Printf("listen and serving on port: %s", port)
	address := fmt.Sprintf(":%s", "8080")
	err := router.Run(address)
	if err != nil {
		log.Fatalf("Cannot start the server: %v ", err.Error())
		return
	}
}
