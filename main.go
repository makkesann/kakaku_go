package main

func main() {
	r := server.GetRouter()
	r.Run(":8080")
}
