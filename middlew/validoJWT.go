package middlew

import (
	"net/http"

	"github.com/tobiasArce/tweter/routers"
)

/*ValidoJWT permite validar el JWT que nos llega en la peticion*/
func ValidoJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routers.ProcesoToken(r.Header.Get("Authoritation"))
		if err != nil {
			http.Error(w, "Error en el Token!"+err.Error(), http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	}
}
