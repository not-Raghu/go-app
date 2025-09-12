package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/not-raghu/goober/api"
)

func main() {

	godotenv.Load(".env")

	//gin
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)
	router := gin.New()

	api.Api(router)

	port := ":" + os.Getenv("PORT")
	if err := router.Run(port); err != nil {
		fmt.Print("process exit")
	}

}
