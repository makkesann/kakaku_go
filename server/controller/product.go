package controller

import (
	"server/dbmod"
	"server/model"

	"github.com/gin-gonic/gin"
)

// MVCにおけるmodel部分

func FindAllProducts() []model.Product {
	db := dbmod.SqlConnect()
	products := []model.Product{}

	// select文
	db.Order("ID asc").Find(&products)
	defer db.Close()
	return products
}

func InsertProduct(c *gin.Context, registerProduct *model.Product) {
	db := dbmod.SqlConnect()
	// insert
	result := db.Create(&registerProduct)
	err := result.Error
	if err != nil {
		c.JSON(400, err)
	} else {
		c.JSON(200, result)
	}
	defer db.Close()
}

//MVCにおけるcontroller部分

func FetchAllProducts(c *gin.Context) {
	resultProducts := FindAllProducts()

	// URLへのアクセスに対してJSONを返す
	c.JSON(200, resultProducts)
}

func AddProduct(c *gin.Context) {
	product_name := c.PostForm("product_name")

	var product = model.Product{
		Name: product_name,
	}

	InsertProduct(c, &product)
}
