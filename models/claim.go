package models

import (
	jwt "github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*Claim es la estrucutura usada para procesar el JWT*/
type Claim struct {
	Email string             `bson:"_id" json:"_id,omitempty"`
	ID    primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	jwt.StandardClaims
}
