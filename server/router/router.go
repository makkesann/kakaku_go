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
			"http://54.65.204.164:8083",
			"http://54.65.204.164:8080",
			"http://54.65.204.164:80",
		},
		AllowMethods: []string{
			"POST",
			"GET",
			"OPTIONS",
			"PUT",
			"DELETE",
		},
	}))
	r.POST("/user/new", controller.AddUser)
	r.GET("/drinks", controller.FetchAllDrinks)
	r.GET("/shops", controller.FetchAllShops)
	r.GET("/drink/genres", controller.FetchAllDrinkGenres)
	r.POST("/drink/genre/add", controller.AddDrinkGenre)
	r.GET("/drinks/:id", controller.FetchDrink)
	r.POST("/drink/add", controller.AddDrink)
	r.POST("/drink/:id/delete", controller.DeleteDrink)
	r.POST("/price/add", controller.AddDrinkPrice)
	r.GET("/drinks/:id/prices", controller.FetchPrices)
	r.POST("/login", controller.UserLogin)
	r.POST("/favorite_drink/add", controller.AddFavoriteDrink)
	r.POST("/favorite_drink/delete", controller.DeleteFavoriteDrink)
	r.POST("/favorite_shop/add", controller.AddFavoriteShop)
	r.POST("/favorite_shop/delete", controller.DeleteFavoriteShop)
	r.GET("/user/:id/favorite_drink", controller.FetchFavoriteDrink)
	r.GET("/user/:id/favorite_shop", controller.FetchFavoriteShop)
	return r
}
