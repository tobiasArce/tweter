package bd

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoCN es el objeto de conexion a la BD"*/
var MongoCN = ConectarBD()
var clientOptions = options.Client().ApplyURI("mongodb+srv://lolito93:River2015@tweter.vqup5.mongodb.net/myFirstDatabase?retryWrites=true&w=majority")

/* ConectarBD es la funcion que permite conectarse a la base de datos*/
func ConectarBD() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("conexion con la BD exitosa")
	return client
}

/* ChequeoConection es el ping a la BD*/
func ChequeoConection() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
