package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os"
	"starbucks-app/router"
)

func main() {
	r := gin.Default()
	
	router.InitRoutes(r)
	
	port := os.Getenv("PORT")
	
	if port == "" {
		port = "8080"
	}
	log.Printf("listen and serving on port: %s", port)
	address := fmt.Sprintf(":%s", port)
	err := r.Run(address)
	if err != nil {
		log.Fatalf("Cannot start the server: %v ", err.Error())
		return
	}
}
