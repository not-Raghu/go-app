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

	seed()

	fmt.Println("database migraion succesfulll")
}

var users = []models.User{
	{
		Name:     "Alice",
		Email:    "alice@example.com",
		Password: "password123",
	},
	{
		Name:     "Bob",
		Email:    "bob@example.com",
		Password: "password456",
	},
}

func seed() {

	for i := range users {
		result := Db.FirstOrCreate(&users[i], models.User{Email: users[i].Email})
		if result.Error != nil {
			log.Fatalf("Failed to seed user %s: %v", users[i].Name, result.Error)
		}
	}

	// for i := range blogs {
	// 	result := Db.FirstOrCreate(&blogs[i], models.Blog{Title: blogs[i].Title})
	// 	if result.Error != nil {
	// 		log.Fatalf("Failed to seed blog '%s': %v", blogs[i].Title, result.Error)
	// 	}
	// }
}
