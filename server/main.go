package main

import (
	"server/others"
)

func main() {
	route := others.GetRouter()
	route.Run(":8081")
}
