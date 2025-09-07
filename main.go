package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/not-raghu/goober/controllers/test"
	"github.com/not-raghu/goober/middleware"
)

func main() {

	godotenv.Load(".env")
	ginMode := os.Getenv("GIN_MODE")

	gin.SetMode(ginMode)

	router := gin.New()

	//midddleware
	router.Use(middleware.Logger(), gin.Recovery())

	api := router.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

	router.GET("/test", test.Test())

	router.NoRoute(func(c *gin.Context) {
		c.File("WrongRoute/404.html")
	})

	port := ":" + os.Getenv("PORT")

	if err := router.Run(port); err != nil {
		fmt.Print("process exit")
	}
}
