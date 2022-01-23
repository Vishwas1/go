package main

func main() {
	router := InitRoutes()
	router.Run("localhost:8080")
}
