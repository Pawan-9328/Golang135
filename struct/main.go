package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct {
	Email string
	Phone string
}

type Address struct {
	House int
	Area  string
	State string
}

type Employee struct {
	  person_details Person 
	  contact_details Contact
	  address_details Address



}

func main() {

	var keshav Person
	fmt.Println("Prince Person :", keshav) // " " " " 0 - print default value
	keshav.FirstName = "Keshav"
	keshav.LastName = "Gola"
	keshav.Age = 78
	//fmt.Println("Keshav Person : ", keshav)

	// person1 := Person{
	// 	FirstName: "Aakash",
	// 	LastName:  "Singh",
	// 	Age:       29,
	// }

	//fmt.Println("Person 1 ", person1)

	//new keyword

	var employee1 = Employee
	employee1.person_details = Person {
		  FirstName: "Pawan",
		  LastName: "Kumar",
		  Age:56,
	}

	employee1.person_contact.Email = "pawngola@gmail.com"
	employee1.person_contact.Phone = "9989000089"



}
