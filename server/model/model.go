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

type DrinkPriceName struct {
	ShopID uint
	Name   string // name
	Price  uint
}

type Shop struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type Product struct {
	gorm.Model
	Name string `json:"name" gorm:"not null"`
}

type Drink struct {
	gorm.Model
	Name         string `json:"name" gorm:"unique;not null"`
	DrinkGenreID uint
	DrinkPrices  []DrinkPrice
	Jan          uint64
	Image        string
}

type DrinkGenre struct {
	gorm.Model
	Name   string `json:"name" gorm:"unique;not null"`
	Drinks []Drink
}

type User struct {
	gorm.Model
	Name string `json:"name" gorm:"unique;not null"`
	Pass string `gorm:"not null"`
}

type FavoriteDrink struct {
	UserID  uint `gorm:"not null;primary_key"`
	DrinkID uint `gorm:"not null;primary_key"`
}

type FavoriteShop struct {
	UserID uint `gorm:"not null;primary_key"`
	ShopID uint `gorm:"not null;primary_key"`
}
