package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	IsVerified bool   `gorm:"default:false;not null"`
	AvatarUrl  string `gorm:"default:'https://tse4.mm.bing.net/th/id/OIP.U9mrdYXcN3yMCZXFUcMMeAHaHa?cb=12&w=980&h=980&rs=1&pid=ImgDetMain&o=7&rm=3';not null"`
	Role       string `gorm:"not null;default:'user'"`
}

type Status string

const (
	Draft     Status = "pending"
	Published Status = "approved"
)

type Blog struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Thumbnail string
	Content   string `gorm:"not null"`
	Likes     int    `gorm:"default:0"`
	Dislikes  int    `gorm:"default:0"`
	Status    Status

	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type Image struct {
	Url        string
	Descriptin string
	Title      string
}

// type RefreshToken struct {
// 	Token string `gorm:""`
// }
