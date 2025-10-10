package main

import (
	"io"
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

	reqfile, _ := os.OpenFile("reqfile.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultWriter = io.MultiWriter(reqfile)

	errfile, _ := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	gin.DefaultErrorWriter = io.MultiWriter(errfile)

	ginMode := os.Getenv("GIN_MODE")
	gin.SetMode(ginMode)

	router := gin.New()
	handler := c.Handler(router)
	api.Api(router)

	port := ":" + os.Getenv("PORT")

	http.ListenAndServe(port, handler)
}
