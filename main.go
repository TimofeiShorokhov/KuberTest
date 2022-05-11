package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"kuber/controllers"
	"log"
	"net/http"
)

func main() {
	var err error
	err = godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}

	router := controllers.InitRoutes()
	log.Fatal(http.ListenAndServe(":8080", router))
	//_, err = db.Exec("Insert into shasum(file, file_path, checksum, algorithm) VALUES ('sda','dsa1','fdas','md5')")

}