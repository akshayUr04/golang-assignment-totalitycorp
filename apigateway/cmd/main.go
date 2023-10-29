package main

import (
	"log"

	"github.com/akshayUr04/totalitycorp-apigateway/pkg/api"
	"github.com/akshayUr04/totalitycorp-apigateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func main() {
	cfg, cfgErr := config.LoadConfig()
	if cfgErr != nil {
		log.Fatalf("failed to load config: %v", cfgErr)
	}
	r := gin.Default()
	api.RegisterRoutes(r, &cfg)
	r.Run(cfg.Port)
	log.Printf("server listening at %v", cfg.Port)
}
