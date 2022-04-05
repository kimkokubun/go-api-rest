package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/kimkokubun/go-api-rest/src/database"
	"github.com/kimkokubun/go-api-rest/src/models"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")

}
func TodasPersonalidades(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalidade

	database.DB.Find(&p)

	json.NewEncoder(w).Encode(p)
}

func RetornaUmaPersonalidade(w http.ResponseWriter, r *http.Request) {
	var personalidade models.Personalidade
	database.DB.First(&personalidade, mux.Vars(r)["id"])
	json.NewEncoder(w).Encode(personalidade)

}
