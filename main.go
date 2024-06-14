package main

import (
	"fmt"

	"github.com/MiguelGFerreira/go-api-rest/models"
	"github.com/MiguelGFerreira/go-api-rest/routes"
)

func main() {
	models.Personalities = []models.Personalitie{
		{Id: 1, Name: "Name 1", History: "History 1"},
		{Id: 2, Name: "Name 2", History: "History 2"},
	}
	fmt.Println("Initializing api server")
	routes.HandleRequest()
}
