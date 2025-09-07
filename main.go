package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/not-raghu/goober/controllers/authentication"
	"github.com/not-raghu/goober/controllers/test"
	"github.com/not-raghu/goober/middleware"
)

func main() {

	godotenv.Load(".env")

	//gin
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)
	router := gin.New()

	Api(router)

	port := ":" + os.Getenv("PORT")
	if err := router.Run(port); err != nil {
		fmt.Print("process exit")
	}
}

func Api(router *gin.Engine) {

	//midddleware
	router.Use(middleware.Logger(), middleware.ServerError(), gin.Recovery())
	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			auth := v1.Group("/auth")
			{
				auth.POST("/register", authentication.Register())
				// outh
				// auth.POST("/verify-otp")
				// auth.POST("/login")
				// auth.DELETE("/logout")
				// auth.POST("/password/reset")
			}
		}

	}

	router.NoRoute(func(c *gin.Context) {
		c.File("WrongRoute/404.html")
	})
	//test
	router.GET("/test", test.Test())
}
