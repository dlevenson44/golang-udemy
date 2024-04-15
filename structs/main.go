package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	// can specify name with contactInfo type
	// contact   contactInfo
	// or can infer name from struct name
	contactInfo //which is equal to contactInfo contactInfo
}

func main() {
	// alternative: alex := {firstName: "Alex", lastName: "Anderson"}
	// other alt: var alex person, this generates an empty struct
	// can then do alex.firstName = "Alex" and alex.lastName = "Alex"
	// this is because we're saying alex is a person struct but we didn't define any values
	// can also do: alex := person{"Alex", "Anderson", contactInfo: { contactInfo{"test@email.com", 12345}}}
	alex := person{"Alex", "Anderson", contactInfo{"test@email.com", 12345}} // can also embed struct like
	// fmt.Println(alex)
	// below prints key-value pairs for structs
	// fmt.Printf("%+v", alex)
	alex.print()
	// &variable operator-- give me memory address of value this is pointing at
	// alexPointer := &alex
	// Go infers that type Person works with type of pointer Person
	alex.updateName("dan")
	alex.print()
}

// receiver to take in struct person
func (p person) print() {
	fmt.Printf("%+v", p)
}

// * pointer operator-- give me the value this memory address is pointing at
// receiver is a pointer that points at a person
func (pointerToPerson *person) updateName(newName string) {
	(*pointerToPerson).firstName = newName
}
