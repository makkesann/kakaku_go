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
	found := db.Where("name = ? AND pass = ?", username, pass).First(&user)
	err := found.Error
	fmt.Print(err)
	// fmt.Print(user)
	if err != nil {
		c.Status(400)
	}
	fmt.Println(found)
	c.JSON(200, found)
	defer db.Close()
	// c.Status(400)

}

//getfavoritedrink
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

//addfavoritedrink
func InsertFavoriteDrink(registerFavoriteDrink *model.FavoriteDrink) {
	db := dbmod.SqlConnect()
	// insert
	db.Create(&registerFavoriteDrink)
	defer db.Close()
}

func AddFavoriteDrink(c *gin.Context) {
	drink_id_str := c.PostForm("drink_id")
	user_id_str := c.PostForm("user_id")
	drink_id64, _ := strconv.ParseUint(drink_id_str, 10, 64)
	user_id64, _ := strconv.ParseUint(user_id_str, 10, 64)
	drink_id := uint(drink_id64)
	user_id := uint(user_id64)

	var favorite_drink = model.FavoriteDrink{
		UserID:  user_id,
		DrinkID: drink_id,
	}
	InsertFavoriteDrink(&favorite_drink)
}

//deletefavoritedrink
func DeleteFavoriteDrink(c *gin.Context) {
	db := dbmod.SqlConnect()
	favorite_drink := []model.FavoriteDrink{}
	drink_id_str := c.PostForm("drink_id")
	user_id_str := c.PostForm("user_id")
	drink_id64, _ := strconv.ParseUint(drink_id_str, 10, 64)
	user_id64, _ := strconv.ParseUint(user_id_str, 10, 64)
	drink_id := uint(drink_id64)
	user_id := uint(user_id64)

	db.Where("user_id = ? AND drink_id = ?", user_id, drink_id).Delete(&favorite_drink)
	defer db.Close()
}

//getfavoriteshop

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
