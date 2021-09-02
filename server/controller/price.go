package controller

import (
	"server/dbmod"
	"server/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FindPrices(drinkID int) []model.DrinkPrice {
	drinkprice := []model.DrinkPrice{}

	db := dbmod.SqlConnect()
	// select
	db.Order("price").Where("drink_id = ?", drinkID).Find(&drinkprice)
	defer db.Close()

	return drinkprice
}

func FetchPrices(c *gin.Context) {
	var id int
	idstr := c.Param("id")
	id, _ = strconv.Atoi(idstr)
	result := FindPrices(id)

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, result)
}
