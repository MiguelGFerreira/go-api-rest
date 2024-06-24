package routes

import (
	"log"
	"net/http"

	"github.com/MiguelGFerreira/go-api-rest/controllers"
	"github.com/MiguelGFerreira/go-api-rest/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/personalities", controllers.AllPersonalities).Methods("Get")
	r.HandleFunc("/personalities/{id}", controllers.GetPersonalitie).Methods("Get")
	r.HandleFunc("/personalities", controllers.CreatePersonalitie).Methods("Post")
	r.HandleFunc("/personalities/{id}", controllers.DeletePersonalitie).Methods("Delete")
	r.HandleFunc("/personalities/{id}", controllers.PutPersonalitie).Methods("Put")
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
