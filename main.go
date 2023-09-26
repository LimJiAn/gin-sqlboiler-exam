package main

import (
	"github.com/LimJiAn/gin-sqlboiler-example/api/route"
	"github.com/LimJiAn/gin-sqlboiler-example/database"
	"github.com/gin-gonic/gin"
)

func init() {
	database.ConnectDB()
}

func main() {
	r := gin.New()

	// Middlewares
	r.Use(gin.Logger())

	// Routes
	route.SetupRoutes(r)

	// Run the server
	r.Run(":8080")
}
