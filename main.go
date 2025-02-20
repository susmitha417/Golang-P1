package main

import (
	"MyProject/array"
	"MyProject/calci"
	"MyProject/loops"
	"MyProject/structures"
	"fmt"
)

func main() {
	// Calculator functionality
	//loops
	a, b := 2, 4

	sum, sub, mul, div := calci.Asmd(a, b)

	fmt.Println("Calculator Results:")
	fmt.Println("The sum of a, b is", sum)
	fmt.Println("The sub of a, b is", sub)
	fmt.Println("The mul of a, b is", mul)
	fmt.Println("The div of a, b is", div)
	loops.StudentMarks()

	//structure
	// Created an instance of the Name struct
	name := structures.Name{
		Fname: "Ajay",
		Lname: "Doez",
	}

	// Created an instance of the Address struct
	address := structures.Address{
		Flat:      101,
		Residence: "creekwood",
		Name:      name, // Nested struct
	}

	// Access struct fields
	fmt.Println("Accessing Struct Fields:")
	fmt.Println("First Name:", address.Name.Fname)
	fmt.Println("Last Name:", address.Name.Lname)
	fmt.Println("Residence:", address.Residence)

	// Called the DisplayAddress function from the structre package
	structures.DisplayAddress(address)

	//array
	array.Score()

	fmt.Println("HI")

}
