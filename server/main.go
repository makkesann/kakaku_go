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
	route := router.GetRouter()
	route.Run(":8081")
}
