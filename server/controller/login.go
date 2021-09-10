package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UserLogin(c *gin.Context) {
	username := c.PostForm("username")
	pass := c.PostForm("pass")
	fmt.Println(username)
	user := model.User{}
	db := dbmod.SqlConnect()
	err := db.Where("name = ? AND pass = ?", username, pass).First(&user).Error
	fmt.Print(err)
	// fmt.Print(user)
	if err != nil {
		c.Status(400)
	}
	defer db.Close()
	// c.Status(400)

}

//favoritedrink
func FindFavoriteDrink(userID int) []model.FavoriteDrink {
	favorite_drinks := []model.FavoriteDrink{}

	db := dbmod.SqlConnect()
	// select
	db.Where("user_id = ?", userID).Find(&favorite_drinks)
	defer db.Close()

	return favorite_drinks
}

func FetchFavoriteDrink(c *gin.Context) {
	var id int
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	result := FindFavoriteDrink(id)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}

//favoriteshop

func FindFavoriteShop(userID int) []model.FavoriteShop {
	favorite_shops := []model.FavoriteShop{}

	db := dbmod.SqlConnect()
	// select
	db.Where("user_id = ?", userID).Find(&favorite_shops)
	defer db.Close()

	return favorite_shops
}

func FetchFavoriteShop(c *gin.Context) {
	var id int
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	result := FindFavoriteShop(id)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}
