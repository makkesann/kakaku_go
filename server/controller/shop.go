package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindAllShops() []model.Shop {
	db := dbmod.SqlConnect()
	shops := []model.Shop{}

	// select文
	db.Order("ID asc").Find(&shops)
	defer db.Close()
	return shops
}

func FetchAllShops(c *gin.Context) {
	resultShops := FindAllShops()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultShops)
}

func DeleteShop(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	shop := model.Shop{}
	db.Where("id = ?", id).Delete(&shop)
	defer db.Close()
}

func UpdateShopName(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	shop_name := c.PostForm("name")
	shop := model.Shop{}
	fmt.Print(shop_name)
	fmt.Print("shop_name")
	result := db.Model(&shop).Where("id = ?", id).Update("name", shop_name)
	err := result.Error
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
	defer db.Close()
}

func InsertShop(c *gin.Context, registerDrink *model.Shop) {
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

func AddShop(c *gin.Context) {
	shop_name := c.PostForm("shop_name")

	var shop = model.Shop{
		Name: shop_name,
	}

	InsertShop(c, &shop)
}
