package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindPrices(drinkID int) []model.DrinkPriceName {
	// drinkprice := model.DrinkPrice{}

	drinkpricename := []model.DrinkPriceName{}

	db := dbmod.SqlConnect()
	// select
	// db.Order("price ASC").Where("drink_id = ?", drinkID).Find(&drinkprice)
	db.Order("price ASC").Table("drink_price").Where("drink_id = ?", drinkID).Select("drink_price.id, drink_price.shop_id, shop.name, drink_price.price").Joins("left join shop on shop.id = drink_price.shop_id").Scan(&drinkpricename)
	fmt.Println(drinkpricename)
	defer db.Close()

	return drinkpricename
}

func FetchPrices(c *gin.Context) {
	var id int
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	result := FindPrices(id)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}

func InsertDrinkPrice(c *gin.Context, registerDrinkPrice *model.DrinkPrice) {
	db := dbmod.SqlConnect()
	// insert

	result := db.Create(&registerDrinkPrice)
	err := result.Error
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
	defer db.Close()
}

func AddDrinkPrice(c *gin.Context) {
	drink_id_str := c.PostForm("drink_id")
	shop_id_str := c.PostForm("shop_id")
	price_str := c.PostForm("price")
	drink_id64, _ := strconv.ParseUint(drink_id_str, 10, 64)
	shop_id64, _ := strconv.ParseUint(shop_id_str, 10, 64)
	price64, _ := strconv.ParseUint(price_str, 10, 64)
	drink_id := uint(drink_id64)
	shop_id := uint(shop_id64)
	price := uint(price64)

	var drink_price = model.DrinkPrice{
		Price:   price,
		DrinkID: drink_id,
		ShopID:  shop_id,
	}

	InsertDrinkPrice(c, &drink_price)
}

func DeletePrice(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	drink_price := model.DrinkPrice{}
	db.Where("id = ?", id).Delete(&drink_price)
	defer db.Close()
}

func UpdatePricePrice(c *gin.Context) {
	var id int
	db := dbmod.SqlConnect()
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	pricestr := c.PostForm("price")
	price, _ := strconv.Atoi(pricestr)
	drink_price := model.DrinkPrice{}
	fmt.Print(price)
	fmt.Print("price")
	result := db.Model(&drink_price).Where("id = ?", id).Update("price", price)
	err := result.Error
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
	defer db.Close()
}
