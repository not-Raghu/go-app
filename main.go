package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/not-raghu/go-app/api"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/helpers"
	"github.com/not-raghu/go-app/middleware"
)

func init() {
	helpers.LoggerInit()
	godotenv.Load(".env")
	db.ConnectDb()
	db.InitRedis()
}

func main() {

	c := middleware.CorsConfig()

	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	router := gin.New()
	handler := c.Handler(router)
	api.Api(router)

	port := ":" + os.Getenv("PORT")

	http.ListenAndServe(port, handler)
}
