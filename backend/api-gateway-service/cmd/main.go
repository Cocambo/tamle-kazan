package main

import (
	"fmt"
	"log"

	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/config"
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/router"
)

func main() {
	config.LoadConfig()

	r := router.SetupRouter()

	port := config.AppConfig.ServerPort
	log.Printf("API Gateway running on port %s", port)

	if err := r.Run(fmt.Sprintf(":%s", port)); err != nil {
		log.Fatalf("Failed to start API Gateway: %v", err)
	}
}
