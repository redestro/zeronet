package main

import (
	"net/http"

	"github.com/shubhamgupta2956/zeronet/server/pkg/handlers"
	"github.com/shubhamgupta2956/zeronet/server/pkg/sessions"
)

func main() {
	db := sessions.InitDB()

	router := handlers.InitRoutes(db)

	// Start serving the application
	http.ListenAndServe(":8000", router)
}
