package models

/* Tweet captura del body, el mesnaje que nos llega*/
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje"`
}
