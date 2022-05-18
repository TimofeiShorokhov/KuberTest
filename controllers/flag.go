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

func CallFunc(flagVar string, anotherFlag bool, getFlag bool, putFlag bool) {
	switch {
	case len(flagVar) > 0:
		Printer(flagVar)
	case anotherFlag:
		AnotherPrinter()
	case getFlag:
		repo.GetData()
	case putFlag:
		repo.PutTable()

	default:
		fmt.Println("error with flag usage")
	}
}
