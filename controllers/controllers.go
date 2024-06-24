package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MiguelGFerreira/go-api-rest/database"
	"github.com/MiguelGFerreira/go-api-rest/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalitie
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func GetPersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var p models.Personalitie
	database.DB.First(&p, id)
	json.NewEncoder(w).Encode(p)
}

func CreatePersonalitie(w http.ResponseWriter, r *http.Request) {
	var newPersonalitie models.Personalitie
	json.NewDecoder(r.Body).Decode(&newPersonalitie)
	database.DB.Create(&newPersonalitie)
	json.NewEncoder(w).Encode(newPersonalitie)
}

func DeletePersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalitie models.Personalitie
	database.DB.Delete(&personalitie, id)
	json.NewEncoder(w).Encode(personalitie)
}

func PutPersonalitie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalitie models.Personalitie
	database.DB.First(&personalitie, id)
	json.NewDecoder(r.Body).Decode(&personalitie)
	database.DB.Save(&personalitie)
	json.NewEncoder(w).Encode(personalitie)
}
