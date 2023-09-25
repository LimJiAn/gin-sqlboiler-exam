package main

import (
	"github.com/LimJiAn/gin-sqlboiler-exam/api/route"
	"github.com/LimJiAn/gin-sqlboiler-exam/database"
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
