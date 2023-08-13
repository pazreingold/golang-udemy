package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

//Option 1
//type person struct {
//	firstName string
//	lastName  string
//	contact   contactInfo
//}

// Option 2
type person struct {
	firstName string
	lastName  string
	contactInfo
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
	paz.contactInfo = contactInfo{
		email:   "paz@gmail.com",
		zipCode: 94000,
	}

	pazPointer := &paz
	pazPointer.updateName("Raz")
	paz.print()
}

func (pointerToPerson *person) updateName(newFirstName string) {
	(*pointerToPerson).firstName = newFirstName
}
func (p person) print() {
	fmt.Printf("%+v", p)

}
