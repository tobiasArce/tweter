package middlew

import (
	"net/http"

	"github.com/tobiasArce/tweter/bd"
)

/*ChequeoBD es el middlew que me permite conocer el estado de la BD*/
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if bd.ChequeoConection() == 0 {
			http.Error(w, "conexion con la base de datos perdida", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
