package model

import (
	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name string `json:"name"`
}

type User struct {
	gorm.Model
	Name string `json:"name"`
	Mail string `json:"mail"`
}
