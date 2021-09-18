package controller

import (
	"server/dbmod"
	"server/model"

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
