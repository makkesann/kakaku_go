package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func InsertDrinkGenre(c *gin.Context, registerDrink *model.DrinkGenre) {
	db := dbmod.SqlConnect()
	// insert
	result := db.Create(&registerDrink)
	err := result.Error
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}

	defer db.Close()
}

func AddDrinkGenre(c *gin.Context) {
	genre_name := c.PostForm("genre_name")

	var drink_genre = model.DrinkGenre{
		Name: genre_name,
	}

	InsertDrinkGenre(c, &drink_genre)
}

func DeleteDrinkGenre(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	drink_genre := model.DrinkGenre{}
	db.Where("id = ?", id).Delete(&drink_genre)
	defer db.Close()
}

func UpdateDrinkGenreName(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	genre_name := c.PostForm("name")
	drink_genre := model.DrinkGenre{}
	fmt.Print(genre_name)
	fmt.Print("genre_name")
	result := db.Model(&drink_genre).Where("id = ?", id).Update("name", genre_name)
	err := result.Error
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
	defer db.Close()
}
