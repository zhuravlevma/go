package main

import "fmt"

type Character interface {
	GetFullName() string
}

func PrintNameOfCharacter(c Character) {
	switch c.(type) {
	case *Animal:
		fmt.Println("It's an animal", c.GetFullName())
	case *Person:
		fmt.Println("It's a person", c.GetFullName())
	}
}
