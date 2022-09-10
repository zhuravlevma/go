package main

type Animal struct {
	Name string
}

func (a *Animal) GetFullName() string {
	return a.Name
}