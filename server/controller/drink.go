package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllDrinks() []model.Drink {
	db := dbmod.SqlConnect()
	drinks := []model.Drink{}

	// select文
	db.Order("ID asc").Find(&drinks)
	defer db.Close()
	return drinks
}

func FetchAllDrinks(c *gin.Context) {
	resultDrinks := FindAllDrinks()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultDrinks)
	fmt.Print(resultDrinks)
}

func FindAllDrinkGenres() []model.DrinkGenre {
	db := dbmod.SqlConnect()
	drink_genres := []model.DrinkGenre{}

	// select文
	db.Order("ID asc").Find(&drink_genres)
	defer db.Close()
	return drink_genres
}

func FetchAllDrinkGenres(c *gin.Context) {
	resultDrinkGenres := FindAllDrinkGenres()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultDrinkGenres)
}

func InsertDrinkGenre(registerDrink *model.DrinkGenre) {
	db := dbmod.SqlConnect()
	// insert
	db.Create(&registerDrink)
	defer db.Close()
}

func AddDrinkGenre(c *gin.Context) {
	genre_name := c.PostForm("genre_name")

	var drink_genre = model.DrinkGenre{
		Name: genre_name,
	}

	InsertDrinkGenre(&drink_genre)
}

func InsertDrink(registerDrink *model.Drink) {
	db := dbmod.SqlConnect()
	// insert
	db.Create(&registerDrink)
	defer db.Close()
}

func AddDrink(c *gin.Context) {
	drink_name := c.PostForm("drink_name")
	genre_id_str := c.PostForm("genre_id")
	genre_id64, _ := strconv.ParseUint(genre_id_str, 10, 64)
	genre_id := uint(genre_id64)
	jan_str := c.PostForm("jan")
	jan64, _ := strconv.ParseUint(jan_str, 10, 64)
	image := c.PostForm("image")
	quantity_str := c.PostForm("quantity")
	quantity64, _ := strconv.ParseUint(quantity_str, 10, 64)
	quantity := uint(quantity64)

	var drink = model.Drink{
		Name:         drink_name,
		DrinkGenreID: genre_id,
		Jan:          jan64,
		Image:        image,
		Quantity:     quantity,
	}

	InsertDrink(&drink)
}

func FindDrink(drinkID int) []model.Drink {
	drink := []model.Drink{}

	db := dbmod.SqlConnect()
	// select
	db.First(&drink, drinkID)
	defer db.Close()

	return drink
}

func FetchDrink(c *gin.Context) {
	var id int
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	resultDrink := FindDrink(id)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultDrink)
}

func DeleteDrink(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	drink := []model.Drink{}
	db.Where("id = ?", id).Delete(&drink)
	defer db.Close()
}

func DeleteDrinkGenre(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	drink := []model.Drink{}
	db.Where("id = ?", id).Delete(&drink)
	defer db.Close()
}
