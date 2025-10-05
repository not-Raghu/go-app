package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/not-raghu/go-app/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func main() {

	godotenv.Load(".env")
	connStr := os.Getenv("DATABASE_URL")

	if connStr == "" {
		log.Fatal("NO DATABASE URL")
	}

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})

	if err != nil {
		log.Fatal("couldn't connect to the database")
		return
	}

	err = db.AutoMigrate(&models.User{}, &models.Blog{})

	if err != nil {
		log.Fatal("failed to do database migrations.")
	}

	Db = db

	fmt.Println("database migraion succesfulll")
}
