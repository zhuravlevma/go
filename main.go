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

func getParents() {
	const size = 2 // с var компилятор будет ругаться
	var parents [size] int

	fmt.Println(parents)
}

func getChildren()  {
	children := [...]int{1, 2, 3}

	fmt.Println(children)
}

func getFriends() {
	var friends [] int
	fmt.Println(friends)

	friendsNew := make([] int, 0)

	friendsNew = append(friendsNew, 2, 3)
	fmt.Println("Friends new:", friendsNew)

	copyFriends := make([] int, len(friendsNew), len(friendsNew))
	fmt.Println("copy friends init:", copyFriends)
	copy(copyFriends, friendsNew)
	fmt.Println("copy friends:", copyFriends)
}

func childHeightInfo() {
	var person map[string]int = map[string]int{
		"Vladimir": 170,
	}
	person["Kate"] = 153
	person["Maksim"] = 145

	_, personExists := person["Goga"]

	fmt.Println("Goga exists", personExists)

	if _, kateExists := person["Kate"]; kateExists {
		fmt.Println("Kate exists in person map")
	}
}

func main() {
	getPersonalInfo()
	getParents()
	getChildren()
	getFriends()
	childHeightInfo()
}


