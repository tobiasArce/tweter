package models

import "time"

/*GraboTweet es el formato que tendra nuestro tweet en la BD*/
type GraboTweet struct {
	UserID  string    `bson:"userid" json:"userid,omitempty"`
	Mensaje string    `bson:"mensaje" json:"mensaje,omitempty"`
	Fecha   time.Time `bson:"fecha" json:"fecha,omitempty"`
}
