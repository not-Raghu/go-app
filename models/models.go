package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string //by default some random name will be assigned like reddit , later on user can do something about it (what do i make other than a blog app :<)
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
