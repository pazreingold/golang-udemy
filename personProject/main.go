package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

func main() {
	//First Option
	//paz := person{"Paz", "Reingold"}

	//Second Option
	//paz := person{firstName: "Paz", lastName: "Reingold"}

	//Third Option
	var paz person
	paz.firstName = "Paz"
	paz.lastName = " Reingold"

	fmt.Print(paz)
	fmt.Printf("%+v", paz)
}
