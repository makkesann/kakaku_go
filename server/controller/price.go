package controller

import (
	"fmt"
	"server/dbmod"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindPrices(drinkID int) []model.DrinkPriceName {
	// drinkprice := []model.DrinkPrice{}

	drinkpricename := []model.DrinkPriceName{}

	db := dbmod.SqlConnect()
	// select
	// db.Order("price ASC").Where("drink_id = ?", drinkID).Find(&drinkprice)
	db.Order("price ASC").Table("drink_price").Where("drink_id = ?", drinkID).Select("shop_id, shop.name, drink_price.price").Joins("left join shop on shop.id = drink_price.drink_id").Scan(&drinkpricename)
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
