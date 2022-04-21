package bd

import (
	"context"
	"time"

	"github.com/tobiasArce/tweter/models"
	"go.mongodb.org/mongo-driver/bson"
)

/*ConsultoRelacion consulta la relacion entre dos usuarios*/
func ConsultoRelacion(t models.Relacion) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("tweter")
	col := db.Collection("relacion")

	condicion := bson.M{
		"usuarioid":         t.UsuarioID,
		"usuariorelacionid": t.UsuarioRelacionID,
	}

	var resultado models.Relacion

	err := col.FindOne(ctx, condicion).Decode(&resultado)
	if err != nil {
		return false, err
	}
	return true, nil
}
