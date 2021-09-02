package router

import (
	"server/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine { // *gin.Engineの表記は返り値の型
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// アクセスを許可したいアクセス元
		AllowOrigins: []string{
			"http://localhost:8083",
			"http://localhost:8080",
		},
		// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
		},
	}))
	r.GET("/:genre", controller.ShowMassage)
	r.POST("/user/new", controller.NewUser)
	r.GET("/user/show", controller.ShowUser)
	r.GET("/products", controller.FetchAllProducts)
	r.POST("/product/add", controller.AddProduct)
	r.GET("/drinks", controller.FetchAllDrinks)
	r.GET("/drink/genres", controller.FetchAllDrinkGenres)
	r.POST("/drink/add", controller.AddDrink)

	return r
}
