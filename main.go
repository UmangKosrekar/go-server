package main

import (
	"first_server/api/routes"

	"first_server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	// connect mongoDB
	database.ConnectMongoDB()

	server := gin.Default()
	routes.RegisterRoutes(server)

	server.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
