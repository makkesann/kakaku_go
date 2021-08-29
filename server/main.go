package main

import (
	"github.com/makkesann/kakaku_go/server/controller"
)

func main() {
	route := controller.GetRouter()
	route.Run(":8081")
}
