package controllers

import (
	"fmt"
	"kuber/repo"
)

func Printer(flag string) {
	fmt.Println(flag)
}

func AnotherPrinter() {
	fmt.Println("great, another function is working!")
}

func CallFunc(flagVar string, anotherFlag bool, getFlag bool, putFlag bool, hash string, check bool) {
	switch {
	case len(flagVar) > 0:
		Printer(flagVar)
	case anotherFlag:
		AnotherPrinter()
	case getFlag:
		fmt.Println(repo.GetData())
	case putFlag:
		repo.PutTable()
	case len(hash) > 0:
		fmt.Println(repo.HashOfFile(hash))
	case check:
		repo.CheckData()

	default:
		fmt.Println("error with flag usage")
	}
}
