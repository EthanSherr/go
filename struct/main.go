package main

import "fmt"

type person struct {
	firstName   string
	lastName    string
	contactInfo // don't have to specify field name, can just reuse the field type
}

type contactInfo struct {
	email string
	zip   int
}

func (p person) toString() string {
	return p.firstName + " " + p.lastName
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func main() {
	ethan := person{
		firstName: "Ethan",
		lastName:  "Sherr",
		contactInfo: contactInfo{
			email: "esherrthan@gmail.com",
			zip:   02446,
		},
	}

	// ethan.updateName("No")

	fmt.Println("ethan.firstName", ethan.firstName)
}
