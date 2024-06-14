package routes

import (
	"log"
	"net/http"

	"github.com/MiguelGFerreira/go-api-rest/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalitie).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
