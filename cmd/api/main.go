package main

import (
	"item-comparison-ai-api/config"
	"item-comparison-ai-api/internal/database"
	"item-comparison-ai-api/internal/logger"
	"item-comparison-ai-api/internal/routes"
	"item-comparison-ai-api/internal/server"

	"github.com/gin-gonic/gin"
)

func main() {
	var engine = gin.New()
	var config = config.New()
	var loggerAdapter = logger.NewLogger(config.Environment)
	var logger = loggerAdapter.GetLogger()
	db := database.NewClient(&database.Database{})
	if db == nil {
		logger.Fatal("Failed to create database client")
	}
	var server = server.New(config, db, engine, loggerAdapter).
		WithMiddlewares().
		WithHealthcheck().
		WithHandlers("",
			&routes.ProductRouter{},
		)

	logger.Println("Start Item Comparison AI API...")
	server.Start()
}
