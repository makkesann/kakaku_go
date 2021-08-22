package main

import (
	"github.com/makkesann/kakaku_go/routes"
)

func main() {
	r := routes.GetRouter()
	r.Run(":8080")
}
