package routes

import (
	"github.com/gorilla/mux"
	"github.com/kimkokubun/go-api-rest/src/controllers"
	"log"
	"net/http"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
