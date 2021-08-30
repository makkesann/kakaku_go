package main

import (
	"server/router"
)

func main() {
	route := router.GetRouter()
	route.Run(":8082")
}
