package main

import (
	"tesbackend/config"
	"tesbackend/routes"
)

func main() {
	config.ConnectDB()

	// Setup routes
	r := routes.SetupRouter()

	// Jalankan server
	r.Run(":8080")
}
