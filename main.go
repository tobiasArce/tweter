package main

import (
	"log"

	"github.com/tobiasArce/tweter/bd"
	"github.com/tobiasArce/tweter/handlers"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("sin conexion a la BD")
	}
	handlers.Manejadores()

}
