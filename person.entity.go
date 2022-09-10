package main

type Person struct {
	FirstName string
	SecondName string
}

func (p *Person) SetFirstName(firstName string) {
	p.FirstName = firstName
}

func (p *Person) SetSecondName(secondName string) {
	p.SecondName = secondName
}

func (p *Person) getFullName() string {
	return p.FirstName + " " + p.SecondName
}
