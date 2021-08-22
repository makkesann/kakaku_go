package main

import (
	"github.com/makkesann/kakaku_go/routes/server"
)

func main() {
	r := server.GetRouter()
	r.Run(":8080")
}
