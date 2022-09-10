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
	children := [...]int{1, 2, 3} // slice

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

	height := person["Maksim"]
	switch height {
		case 145:
			fmt.Println("Height is 145")
			fallthrough
		case 180:
			fmt.Println("It's wrong. Height is 145")
		case 190:
			fmt.Println("Height is 190")

	}
}

func loop() {
	var array = []int{2, 3, 5, 7, 11, 13}
	array = append(array, 12)
	for idx := range array {
		fmt.Println("Elem is", array[idx])
	}
}

func getFirstName() string {
	return "Maksim"
}

func getSecondName() string {
	return "Zhuravlev"
}

func getFullName() (fullName string) {
	fullName = getFirstName() + " " + getSecondName()
	return
}

func checkError(number int) (int, error) {
	if number > 100 {
		return 0, fmt.Errorf("NumberInvalid")
	}
	return number, nil
}

func main() {
	personalFunc := func(description string) {
		fmt.Println("Description:", description)
		getPersonalInfo()
	}

	relativeFunc := func(description string) {
		fmt.Println("Description:", description)
		getParents()
		getChildren()
		getFriends()
		childHeightInfo()
	}

	type personalFuncType func(string)
	type relativeFuncType func(string)

	printAllInfo := func (personalF personalFuncType, relativeF relativeFuncType) {
		personalF("Personal Info")
		relativeF("Relative info")
	}

	printAllInfo(personalFunc, relativeFunc)

	loop()

	fmt.Println(getFullName())
	fmt.Println(checkError(102))
}


