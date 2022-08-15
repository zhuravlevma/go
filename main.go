package main

import (
	"fmt"
	"unicode/utf8"
)

func getPersonalInfo() string {
	var fullName string

	var firstName = "Maksim"
	var secondName = "Zhuravlev"
	fullName = firstName + secondName
	fmt.Println(fullName)

	var length = utf8.RuneCountInString(fullName)
	fmt.Println(length)

	var slice = fullName[1:3]

	fmt.Println(slice)

	return fullName
}

func main() {
	var fullName string

	var firstName = "Maksim"
	var secondName = "Zhuravlev"
	fullName = firstName + " " + secondName
	fmt.Println("Full name:", fullName)

	var length = utf8.RuneCountInString(fullName)
	fmt.Println("Length:", length)

	var slice = fullName[1:3]

	fmt.Println("Slice:", slice)
}


