package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"kuber/controllers"
	"log"
)

var (
	flagVar     string
	AnotherFlag bool
	getFlag     bool
	putFlag     bool
	hash        string
	checkFlag   bool
)

//Creating and parsing flags
func init() {
	flag.StringVar(&flagVar, "d", "", "print text")
	flag.BoolVar(&AnotherFlag, "y", false, "another func")
	flag.BoolVar(&getFlag, "g", false, "get from db")
	flag.BoolVar(&putFlag, "p", false, "put data in db")
	flag.BoolVar(&checkFlag, "c", false, "check data in db")
	flag.StringVar(&hash, "h", "", "resolving hash of file")
	flag.Parse()
}

func main() {

	var err error
	err = godotenv.Load()

	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	} else {
		fmt.Println("We are getting the env values")
	}
	controllers.CallFunc(flagVar, AnotherFlag, getFlag, putFlag, hash, checkFlag)
	//_, err = db.Exec("Insert into shasum(file, file_path, checksum, algorithm) VALUES ('sda','dsa1','fdas','md5')")

}
