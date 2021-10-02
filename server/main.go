package main

import (
	"server/router"
)

func main() {
	route := router.GetRouter()
	route.RunTLS(":8082", "./Certificate/server.pem", "./Certificate/server.key")
}
