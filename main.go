package main

import (
	"fmt"

	"github.com/MiguelGFerreira/go-api-rest/database"
	"github.com/MiguelGFerreira/go-api-rest/routes"
)

func main() {
	database.ConnectDatabase()
	fmt.Println("Initializing api server")
	routes.HandleRequest()
}
