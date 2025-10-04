package main

// import (
// 	"log"

// 	"github.com/not-raghu/go-app/models"
// )

// var users = []models.User{
// 	{
// 		Name:     "Alice",
// 		Email:    "alice@example.com",
// 		Password: "password123",
// 	},
// 	{
// 		Name:     "Bob",
// 		Email:    "bob@example.com",
// 		Password: "password456",
// 	},
// }

// func Seed() {

// 	for i := range users {
// 		result := Db.FirstOrCreate(&users[i], models.User{Email: users[i].Email})
// 		if result.Error != nil {
// 			log.Fatalf("Failed to seed user %s: %v", users[i].Name, result.Error)
// 		}
// 	}

// 	// for i := range blogs {
// 	// 	result := Db.FirstOrCreate(&blogs[i], models.Blog{Title: blogs[i].Title})
// 	// 	if result.Error != nil {
// 	// 		log.Fatalf("Failed to seed blog '%s': %v", blogs[i].Title, result.Error)
// 	// 	}
// 	// }
// }
