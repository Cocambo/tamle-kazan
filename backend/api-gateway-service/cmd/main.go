package main

import (
	"fmt"
	"log"

	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/config"
	"github.com/Cocambo/tamle-kazan/backend/api-gateway-service/internal/router"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	r := router.SetupRouter(cfg)

	log.Printf("API Gateway running on port %s", cfg.ServerPort)

	if err := r.Run(fmt.Sprintf(":%s", cfg.ServerPort)); err != nil {
		log.Fatal(err)
	}
}
