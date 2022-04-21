package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/tobiasArce/tweter/bd"
)

/* ObtenerAvatar envia el banner al HTTP */
func ObternerBanner(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	perfil, err := bd.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Perfil no encontrado", http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("uploads/banner/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Error a copiar la imagen", http.StatusBadRequest)
	}
}
