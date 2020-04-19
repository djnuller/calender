package main

import (
	"calender/config"
	"fmt"
	"log"
	"os"
)

func main() {
	err := config.AutoLoad()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Config loaded")
	fmt.Println(os.Getenv("mongouri"))
}
