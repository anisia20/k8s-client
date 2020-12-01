package main

import (
	"os"

	"k8s-client/internal/routes"
)

const defaultPort string = "3000"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	r := routes.SetupRoutes()

	r.Run(":" + port)
}
