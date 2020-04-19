package db

import (
	"calender/models"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(_collection string) (*mongo.Collection, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("mongouri"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}
	// Check connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}
	collection := client.Database(os.Getenv("database")).Collection(_collection)

	return collection, nil
}

func GetError(err error, w http.ResponseWriter) {

	log.Println(err.Error())
	var response = models.ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}
