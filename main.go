package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load(".env")

	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	port := ":" + os.Getenv("PORT")
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "this route doesn't exist brother",
		})
	})

	if err := router.Run(port); err != nil {
		fmt.Print("process exit")
	}
}
