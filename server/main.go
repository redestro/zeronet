package main

import (
	"github.com/shubhamgupta2956/zeronet/server/pkg/handlers"
)

func main() {
	router := handlers.InitRoutes()
	// Start serving the application
	router.Run(":8000")
}
