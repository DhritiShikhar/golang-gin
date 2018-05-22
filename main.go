package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default() // This creates a gin router with default middleware

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello world",
		})
	})
	router.Run()
}
