package models

/*Relacion es el modelo para grabar la relacion de un usuario con otro*/
type Relacion struct {
	UsuarioID         string `bson:"usuarioid" json:"usuarioId"`
	UsuarioRelacionID string `bson:"usuariorelacionid" json:"usuarioRelacioId"`
}
