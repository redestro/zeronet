package main

import (
	"net/http"

	"github.com/shubhamgupta2956/zeronet/server/pkg/database"
	"github.com/shubhamgupta2956/zeronet/server/pkg/handlers"
)

func main() {

	client, ctx := database.NewDB()
	defer client.Disconnect(ctx)

	router := handlers.InitRoutes(client)

	// Start serving the application
	http.ListenAndServe(":8000", router)
}
