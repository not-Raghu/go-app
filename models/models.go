package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name        string `gorm:"not null"`
	Email       string `gorm:"unique;not null"`
	Password    string `gorm:"not null"`
	Is_Verified bool   `gorm:"default:false;not null"`
	AvatarUrl   string `gorm:"default:https://tse4.mm.bing.net/th/id/OIP.U9mrdYXcN3yMCZXFUcMMeAHaHa?cb=12&w=980&h=980&rs=1&pid=ImgDetMain&o=7&rm=3;not null"`
}

type Blog struct {
	gorm.Model
	Title     string `gorm:"not null"`
	Thumbnail string
	Content   string `gorm:"type:text;not null"`

	AuthorID uint
	Author   User `gorm:"foreignKey:AuthorID"`
}
