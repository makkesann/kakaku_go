package controller

import (
	"server/dbmod"
	"server/model"

	"github.com/gin-gonic/gin"
)

// MVCにおけるmodel部分

func FindAllDrinks() []model.Drink {
	db := dbmod.SqlConnect()
	drinks := []model.Drink{}

	// select文
	db.Order("ID asc").Find(&drinks)
	defer db.Close()
	return drinks
}

func FindAllDrinkGenres() []model.DrinkGenre {
	db := dbmod.SqlConnect()
	drink_genres := []model.DrinkGenre{}

	// select文
	db.Order("ID asc").Find(&drink_genres)
	defer db.Close()
	return drink_genres
}

func InsertDrink(registerProduct *model.Drink) {
	db := dbmod.SqlConnect()
	// insert
	db.Create(&registerProduct)
	defer db.Close()
}

//MVCにおけるcontroller部分

func FetchAllDrinkGenres(c *gin.Context) {
	resultDrinkGenres := FindAllDrinkGenres()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultDrinkGenres)
}

func FetchAllDrinks(c *gin.Context) {
	resultDrinks := FindAllDrinks()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultDrinks)
}

func AddDrink(c *gin.Context) {
	drink_name := c.PostForm("drink_name")

	var drink = model.Drink{
		Name: drink_name,
	}

	InsertDrink(&drink)
}
