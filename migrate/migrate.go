package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type Blog struct {
	gorm.Model
	Title   string `gorm:"not null"`
	Content string `gorm:"type:text;not null"`

	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID"`
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

	err = db.AutoMigrate(&User{}, &Blog{})

	if err != nil {
		log.Fatal("failed to do database migrations.")
	}

	fmt.Println("database migraion succesfulll")
}
