package main

import (
	"log"
	"net/http"
	"prd_mngt/api"
	"prd_mngt/cache"
	"prd_mngt/db"
	"prd_mngt/logger"
	"prd_mngt/mq"
)

func main() {
	// Initialize logger
	logger.InitLogger()

	// Initialize Redis cache
	cache.InitRedis()

	// Initialize DB connection
	db.InitDatabase()

	// Initialize RabbitMQ for asynchronous processing
	mq.Connect()

	// Initialize the router
	r := api.InitRoutes()

	// Start the HTTP server
	log.Println("Server started on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
