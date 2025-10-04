package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/not-raghu/go-app/api"
	"github.com/not-raghu/go-app/db"
	"github.com/not-raghu/go-app/middleware"
)

// special functions
// init() - runs first before any other function from that package

// func init(){
// 	gin.SetMode()
// }

func main() {

	godotenv.Load(".env")

	db.ConnectDb()
	db.InitRedis()

	c := middleware.CorsConfig()
	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	router := gin.New()
	handler := c.Handler(router)
	api.Api(router)

	port := ":" + os.Getenv("PORT")

	http.ListenAndServe(port, handler)
}
