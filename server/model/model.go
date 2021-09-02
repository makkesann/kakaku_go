package model

import (
	"github.com/jinzhu/gorm"
)

type DrinkPrice struct {
	gorm.Model
	Price   uint `gorm:"not null" json:"price"`
	DrinkID uint `gorm:"not null"`
	ShopID  uint `gorm:"not null"`
}

type Shop struct {
	gorm.Model
	Name string `gorm:"not null"`
}

type Product struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
}

type Drink struct {
	gorm.Model
	Name         string `json:"name" gorm:"not null"`
	DrinkGenreID uint
	DrinkPrices  []DrinkPrice
}

type DrinkGenre struct {
	gorm.Model
	Name   string `json:"name" gorm:"not null"`
	Drinks []Drink
}

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
	Pass string `gorm:"not null"`
}
