package main

import (
	"github.com/gin-gonic/gin"
	"rest/db"
	"rest/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.SetupRoutes(server)

	server.Run(":8080")
}
