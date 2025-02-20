package structures

import "fmt"

type Name struct {
	Fname string
	Lname string
}

type Address struct {
	Flat      int
	Residence string
	Name      Name
}

func DisplayAddress(address Address) {
	fmt.Println("Address Details:")
	fmt.Printf("Plot: %d, Residence: %s\n", address.Flat, address.Residence)
	fmt.Printf("Name: %s %s\n", address.Name.Fname, address.Name.Lname)
}
