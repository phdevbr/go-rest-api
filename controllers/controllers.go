package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phdevbr/go-rest-api/database"
	"github.com/phdevbr/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalidade models.Personalidade
	database.DB.First(&personalidade, id)
	json.NewEncoder(w).Encode(personalidade)
}

func CriaUmaNovaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var novapersonalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&novapersonalidade)
	database.DB.Create(&novapersonalidade)
	json.NewEncoder(w).Encode(novapersonalidade)
}
