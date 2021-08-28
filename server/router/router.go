package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/makkesann/kakaku_go/server/controller/top"
	"github.com/makkesann/kakaku_go/server/controller/user"
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
	r.GET("/", top.IndexDisplayAction)
	r.GET("/:massage", user.ShowMassage)

	return r
}
