package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Create a Gin router
	router := gin.Default()

	// Define a route for the root endpoint
	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})

	// Start the server on port 8080
	router.Run(":8080")
}
