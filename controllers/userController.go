package controllers

import (
	"calender/config/db"
	"calender/models"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

const (
	CONTENT_TYPE = "application/json"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", CONTENT_TYPE)

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)
	// get table from environment
	collection, err := db.ConnectDB(os.Getenv("userTable"))
	if err != nil {
		db.GetError(err, w)
		return
	}
	var res models.User
	err = collection.FindOne(context.TODO(), bson.D{{"userName", user.UserName}}).Decode(&res)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 5)
			if err != nil {
				err = errors.New("Something went wrong hashing your password.")
				db.GetError(err, w)
				return
			}
			user.Password = string(hash)

			result, err := collection.InsertOne(context.TODO(), user)
			if err != nil {
				err = errors.New("Something went wrong creating your user.")
				db.GetError(err, w)
				return
			}
			json.NewEncoder(w).Encode(result)
			return
		}
		db.GetError(err, w)
		return
	}
	err = errors.New("User already exists")
	db.GetError(err, w)
	return
}

func SignInHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", CONTENT_TYPE)

	var user models.User

	_ = json.NewDecoder(r.Body).Decode(&user)

	collection, err := db.ConnectDB(os.Getenv("userTable"))
	if err != nil {
		db.GetError(err, w)
		return
	}

	var res models.User
	err = collection.FindOne(context.TODO(), bson.D{{"userName", user.UserName}}).Decode(&res)

	if err != nil {
		err = errors.New("Invalid username.")
		db.GetError(err, w)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(user.Password))

	if err != nil {
		err = errors.New("Invalid password.")
		db.GetError(err, w)
		return
	}
	tokenString, err := createToken(res.ID.String())

	if err != nil {
		//err = errors.New("Error while generating token")
		db.GetError(err, w)
		return
	}
	res.Token = tokenString
	res.Password = ""

	json.NewEncoder(w).Encode(res)
}

func createToken(userId string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = userId
	claims["expiration"] = time.Now().Add(time.Minute * 60).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	var key interface{}
	key = []byte(os.Getenv("access_secret"))
	tokenString, err := token.SignedString(key.([]byte))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
