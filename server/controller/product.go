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

func InsertProduct(registerProduct *model.Product) {
	db := dbmod.SqlConnect()
	// insert
	db.Create(&registerProduct)
	defer db.Close()
}

func ShowMassage(c *gin.Context) {
	genre := c.Param("genre")
	c.JSON(200, gin.H{
		"massage": genre,
	})
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

	InsertProduct(&product)
}
