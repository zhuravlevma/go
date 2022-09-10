package main

import "fmt"

type Character interface {
	GetFullName() string
}

func PrintNameOfCharacter(c Character) {
	fmt.Println(c.GetFullName())
}
