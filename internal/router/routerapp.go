package router

import (
	"fmt"
	"log"
	satelliteHandler "meliQuasar/internal/handlers"
	"net/http"

	"github.com/gorilla/mux"
)

func ServicesInit() {
	fmt.Println("Iniciando en localhost:850")
	mux := mux.NewRouter()
	mux.HandleFunc("/api/saludar", satelliteHandler.Hola).Methods("GET")
	log.Fatal(http.ListenAndServe(":850", mux))
}
