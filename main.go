package main

import (
	"calender/config"
	"calender/config/route"
	"fmt"
	"log"
	"net/http"
)

func main() {
	err := config.AutoLoad()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Environment loaded...")
	log.Fatal(http.ListenAndServe(":8080", route.Routes()))
}
