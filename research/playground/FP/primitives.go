package main

type age uint
type phoneNumber string

type Person struct {
	name        string
	age         age
	phonenumber phoneNumber
}

func (a age) valid() bool {
	return a < 120
}

func isValidPerson(p Person) bool {
	return p.age.valid() && p.name != ""
}
