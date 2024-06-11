// structs2
// Make me compile!

package main

import "fmt"

type ContactDetails struct {
	number string
}

type Person struct {
	name   string
	age    int
	phone ContactDetails
}

func main() {
	contactDetails := ContactDetails{number: "344788226"}
	person := Person{name: "John", age: 32, phone: contactDetails}

	fmt.Printf("%s is %d years old and his phone is %s\n", person.name, person.age, person.phone)
}
