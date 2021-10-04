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
			"http://kakaku-real-store.tk:8083",
			"http://kakaku-real-store.tk:8080",
			"http://kakaku-real-store.tk",
			"http://kakaku-real-store.tk:81",
			"https://kakaku-real-store.tk:443",
			"https://kakaku-real-store.tk:444",
			"http://ec2-54-65-204-164.ap-northeast-1.compute.amazonaws.com",
			"https://ec2-54-65-204-164.ap-northeast-1.compute.amazonaws.com",
			"https://kakaku-real-store.tk",
			"http://kakaku-real-store.tk",
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
	r.POST("/shop/add", controller.AddShop)
	r.POST("/shop/:id/delete", controller.DeleteShop)
	r.POST("/shop/name/:id/change", controller.UpdateShopName)
	r.GET("/drink/genres", controller.FetchAllDrinkGenres)
	r.POST("/drink/genre/add", controller.AddDrinkGenre)
	r.POST("/drink/genre/:id/delete", controller.DeleteDrinkGenre)
	r.GET("/drinks/:id", controller.FetchDrink)
	r.POST("/drink/add", controller.AddDrink)
	r.POST("/drink/:id/delete", controller.DeleteDrink)
	r.POST("/drink/name/:id/change", controller.UpdateDrinkName)
	r.POST("/drink/genre/:id/change", controller.UpdateDrinkGenreID)
	r.POST("/drink/genre/name/:id/change", controller.UpdateDrinkGenreName)
	r.POST("/drink/jan/:id/change", controller.UpdateDrinkJan)
	r.POST("/drink/image/:id/change", controller.UpdateDrinkImage)
	r.POST("/drink/quantity/:id/change", controller.UpdateDrinkQuantity)
	r.POST("/price/add", controller.AddDrinkPrice)
	r.POST("/price/:id/delete", controller.DeletePrice)
	r.POST("/price/price/:id/change", controller.UpdatePricePrice)
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
