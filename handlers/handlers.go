package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/tobiasArce/tweter/middlew"
	"github.com/tobiasArce/tweter/routers"
)

/*Manejadores setea el puerto, el handler y pone a escuchar el servidor */
func Manejadores() {
	router := mux.NewRouter()
	router.HandleFunc("/registro", middlew.ChequeoBD(routers.Registro)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
