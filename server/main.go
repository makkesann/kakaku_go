package main

import (
	"github.com/makkesann/kakaku_go/server/others"
)

func main() {
	route := others.GetRouter()
	route.Run(":8081")
}
