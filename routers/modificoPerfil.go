package routers

import (
	"encoding/json"
	"net/http"

	"github.com/tobiasArce/tweter/bd"
	"github.com/tobiasArce/tweter/models"
)

/* ModificoPerfil modifica el perfil del usuaruio */
func ModificoPerfil(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool

	status, err = bd.ModificoRegistro(t, IDUsuario)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Intente nuevamente ", 400)
		return
	}

	if status == false {
		http.Error(w, "No se logro modificar el registro del usuario ", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
