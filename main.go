package main

import (
	"log"
	"os"

	"github.com/Maged-Zaki/gin-rest-api/db"
	"github.com/Maged-Zaki/gin-rest-api/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables file")
	}

	// Connect to Sqlite3
	db.InitializeDatabase()

	server := gin.Default()

	// Register routes
	routes.RegisterRoutes(server)

	// Run server
	server.Run(os.Getenv("PORT"))
}
