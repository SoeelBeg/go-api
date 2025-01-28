package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Use(CORSMiddleware())

	// Handle root path
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to my API! Visit /hello-world for a message.")
	})

	// Handle favicon.ico requests
	router.GET("/favicon.ico", func(c *gin.Context) {
		c.Status(http.StatusNoContent) // Respond with 204 No Content
	})

	// Your API route
	router.GET("/hello-world", myGetFunction)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8060"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

type simpleMessage struct {
	Hello   string `json:"hello"`
	Message string `json:"message"`
}

func myGetFunction(c *gin.Context) {
	simpleMessage := simpleMessage{
		Hello:   "World!",
		Message: "Subscribe to my channel!",
	}

	c.IndentedJSON(http.StatusOK, simpleMessage)
}
