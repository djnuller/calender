package route

import (
	"calender/controllers"

	"github.com/gorilla/mux"
)

func Routes() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/users", controllers.RegisterHandler).Methods("POST")
	r.HandleFunc("/users/signin", controllers.SignInHandler).Methods("POST")

	return r
}
