package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/not-raghu/go-app/helpers"
	"github.com/not-raghu/go-app/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var users = []models.User{
	{
		Email:    "alice@example.com",
		Password: "password123",
	},
	{
		Email:    "bob@example.com",
		Password: "password456",
	},
}

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

	fmt.Println("database migraion succesfulll")

	if len(os.Args) > 1 && os.Args[len(os.Args)-1] == "seed" {
		fmt.Println("seeding database")

		for i := range users {
			hashPass, err := bcrypt.GenerateFromPassword([]byte(users[i].Password), bcrypt.DefaultCost)

			if err != nil {
				log.Printf("error hashing password")
				continue
			}

			db.Create(&models.User{
				Name:        helpers.GeneateNames(),
				Email:       users[i].Email,
				Password:    string(hashPass),
				Is_Verified: false,
			})
		}

		fmt.Println("seeded database")
		return
	}

}
