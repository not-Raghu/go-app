package api

import (
	"github.com/gin-gonic/gin"
	"github.com/not-raghu/go-app/controllers/authentication"
	"github.com/not-raghu/go-app/controllers/test"
	"github.com/not-raghu/go-app/middleware"
)

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
				auth.POST("/verify-otp", authentication.VerifyOtp())
				auth.POST("/forgot-password", authentication.ForgotPass())
				auth.POST("/change-password", authentication.ChangePass())
				auth.POST("/login")
				auth.DELETE("/logout")
				auth.POST("/password/reset")
			}

			blogs := v1.Group("/blog")
			{

				blogs.GET("/blogs")
				blogs.POST("/blogs")
			}
		}
		api.GET("/test", test.Test())
	}

	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{
			"message": "invalid route",
		})
	})

	//test

}
