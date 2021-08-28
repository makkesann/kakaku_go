// package main

// import (
// 	"github.com/makkesann/kakaku_go/routes"
// )

// func main() {
// 	r := routes.GetRouter()
// 	r.Run(":8081")
// }

package main // import "server"

import (
	"github.com/makkesann/kakaku_go/server/router"
)

func main() {
	// r := gin.Default()
	// ここからCorsの設定
	// r.Use(cors.New(cors.Config{
	// 	// アクセスを許可したいアクセス元
	// 	AllowOrigins: []string{
	// 		"http://localhost:8083",
	// 	},
	// 	// アクセスを許可したいHTTPメソッド(以下の例だとPUTやDELETEはアクセスできません)
	// 	AllowMethods: []string{
	// 		"POST",
	// 		"GET",
	// 		"OPTIONS",
	// 	},
	// }))
	// r.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "congraturation!!!!!!!!!!!",
	// 	})
	// })
	route := router.GetRouter()
	route.Run(":8081")
}
