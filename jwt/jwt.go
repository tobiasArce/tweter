package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tobiasArce/tweter/models"
)

/* GeneroJWT genera el encriptado con JWT*/
func GeneroJWT(t models.Usuario) (string, error) {
	miClave := []byte("PracticandoGolang")
	payload := jwt.MapClaims{
		"email":            t.Email,
		"nombre":           t.Nombre,
		"apellido":         t.Apellido,
		"fecha_Nacimiento": t.FechaNacimiento,
		"biografia":        t.Biografia,
		"ubicacion":        t.Ubicacion,
		"sitioWeb":         t.SitioWeb,
		"_id":              t.ID.Hex(),
		"exp":              time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenSrt, err := token.SignedString(miClave)
	if err != nil {
		return tokenSrt, err
	}
	return tokenSrt, nil
}
