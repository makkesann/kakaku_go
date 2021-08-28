package router

import (
	"github.com/gin-gonic/gin"
	"github.com/makkesann/kakaku_go/controller"
)

func GetRouter() *gin.Engine { // *gin.Engineの表記は返り値の型
	router := gin.Default()
	router.GET("/", controller.IndexDisplayAction)

	return router
}
