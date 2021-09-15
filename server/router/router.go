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
			"http://localhost:80",
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
	r.GET("/drinks", controller.FetchAllDrinks)
	r.GET("/drink/genres", controller.FetchAllDrinkGenres)
	r.GET("/drinks/:id", controller.FetchDrink)
	r.POST("/drink/add", controller.AddDrink)
	r.GET("/drinks/:id/prices", controller.FetchPrices)
	r.POST("/login", controller.UserLogin)
	r.POST("/favorite_drink/add", controller.AddFavoriteDrink)
	r.POST("/favorite_drink/delete", controller.DeleteFavoriteDrink)
	r.GET("/user/:id/favorite_drink", controller.FetchFavoriteDrink)
	r.GET("/user/:id/favorite_shop", controller.FetchFavoriteShop)
	return r
}
