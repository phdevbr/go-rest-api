package routes

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/phdevbr/go-rest-api/controllers"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("GET")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("GET")
	log.Fatal(http.ListenAndServe(":8000", r))
}
