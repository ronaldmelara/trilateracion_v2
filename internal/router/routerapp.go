package router

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	satelliteHandler "meliQuasar/internal/handlers"
)

func ServicesInit(){
	mux := mux.NewRouter()
	mux.HandleFunc("/api/saludar", satelliteHandler.Hola).Methods("GET")
	log.Fatal(http.ListenAndServe(":850", mux))
}