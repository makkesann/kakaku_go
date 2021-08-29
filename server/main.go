package main

import (
	"github.com/makkesann/kakaku_go/server/router"
)

func main() {
	route := router.GetRouter()
	route.Run(":8081")
}
