package model

import (
	"github.com/jinzhu/gorm"
)

type DrinkPrice struct {
	gorm.Model
	Price   uint `json:"price"`
	DrinkID uint
	ShopID  uint
}

type Shop struct {
	gorm.Model
	Google string
}

type Product struct {
	gorm.Model
	Name string `json:"name"`
}

type Drink struct {
	gorm.Model
	Name         string `json:"name"`
	DrinkGenreID uint
}

type DrinkGenre struct {
	gorm.Model
	Name   string `json:"name"`
	Drinks []Drink
}

type User struct {
	gorm.Model
	Name string `json:"name"`
	Pass string
}
