package server

import (
	"github.com/gin-gonic/gin"
	"github.com/makkesann/kakaku_go/controller"
)

func GetRouter() *gin.Engine { // *gin.Engineの表記は返り値の型
	router := gin.Default()
	router.LoadHTMLGlob("view/*.html")

	router.GET("/", controller.IndexDisplayAction)
	router.GET("/user", controller.UserListDisplayAction)
	router.GET("/user/add", controller.UserAddDisplayAction)
	router.GET("/user/edit/:id", controller.UserEditDisplayAction)

	return router
}
