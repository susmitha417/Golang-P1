

# Go Project Overview

This project demonstrates various Go programming concepts such as functions, loops, arrays, structs, maps, and methods. 
It uses different packages to perform basic arithmetic operations, handle student marks, work with arrays and maps, and define and manipulate structs.

## Project Structure

```
/MyProject
  ├── main.go
  ├── array/
  │    └── array.go
  ├── calci/
  │    └── calci.go
  ├── loops/
  │    └── loops.go
  ├── maps/
  │    └── maps.go
  ├── methods/
  │    └── methods.go
  └── structures/
       └── structures.go
```

## Files and Functionalities

### 1. `main.go`

The `main.go` file is the entry point of the project, where various functionalities are demonstrated by calling functions from different packages.

#### Key Functionalities:
- **Arithmetic Operations**: Performs basic arithmetic operations (addition, subtraction, multiplication, and division) using the `calci` package.
- **Loops and Input**: Takes student marks as input and evaluates their grades using the `loops` package.
- **Structs**: Demonstrates the creation and usage of nested structs (`Name` and `Address`) from the `structures` package.
- **Array**: Demonstrates taking input for student scores and printing them using the `array` package.
- **Methods**: Defines a `Rectangle` struct and calculates its area using methods from the `methods` package.
- **Maps**: Displays a map of countries and their codes using the `maps` package.

### 2. `calci.go`

This file contains the `Asmd` function, which performs basic arithmetic operations on two integers.

#### Code:
```go
package calci

// Asmd performs four arithmetic operations (Add, Subtract, Multiply, Divide) on two integers
func Asmd(a, b int) (int, int, int, int) { 
	return a + b, a - b, a * b, a / b
}
```

#### Key Concept:
- **Functions**: The `Asmd` function takes two integers and returns their sum, difference, product, and quotient.

### 3. `loops.go`

This file defines functions to take student marks as input, check their grades, and provide feedback based on their marks.

#### Code:
```go
package loops

import "fmt"

// StudentMarks takes input for student marks and checks grades
func StudentMarks() {
	var student1, student2, student3, student4 int

	// Taking input from the user for 4 students' marks
	fmt.Println("Enter marks for Student 1:")
	fmt.Scan(&student1)
	// Repeat for other students...
}
```

#### Key Concept:
- **Loops and Input**: Using a loop to take marks as input and `if-else` conditions to assign grades based on the marks.
- **Memory Address**: Prints the memory address of each student's marks.

### 4. `array.go`

This file demonstrates how to work with arrays by taking student marks as input and printing them.

#### Code:
```go
package array

import "fmt"

// Score function takes marks for 5 students and prints them
func Score() {
	var studentScore int
	arr := [5]int{} // Declare an array with 5 integer elements, initialized to 0

	// Taking input for the array
	for i := 0; i < 5; i++ {
		fmt.Printf("Enter marks for student %d: ", i+1)
		fmt.Scan(&studentScore)
		arr[i] = studentScore // Store the input in the array
	}

	// Printing student scores
	fmt.Println("\nStudent Scores:")
	for i := 0; i < 5; i++ {
		fmt.Printf("Student %d: %d\n", i+1, arr[i])
	}
}
```

#### Key Concept:
- **Arrays**: This file demonstrates how to take input for an array of fixed size (5 in this case) and print the values.

### 5. `maps.go`

This file demonstrates how to use a map to store country codes and names, and then print them.

#### Code:
```go
package maps

import "fmt"

// PrintMap prints all key-value pairs in a given map
func PrintMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, "->", value)
	}
}
```

#### Key Concept:
- **Maps**: This file demonstrates the use of key-value pairs (map) in Go and iterating through them.

### 6. `methods.go`

This file demonstrates how to define methods for structs, specifically for calculating the area of a rectangle.

#### Code:
```go
package methods

import "fmt"

// Rectangle struct with Width and Height properties
type Rectangle struct {
	Width, Height float64
}

// Method to calculate the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Method to display rectangle details
func (r Rectangle) Display() {
	fmt.Printf("Rectangle [Width: %.2f, Height: %.2f, Area: %.2f]\n", r.Width, r.Height, r.Area())
}
```

#### Key Concept:
- **Methods**: This file demonstrates the concept of methods associated with a struct (i.e., `Rectangle`) and how they can be used to perform actions like calculating the area and displaying the details.

### 7. `structures.go`

This file demonstrates how to use nested structs in Go. It defines a `Name` struct and an `Address` struct that contains a `Name` struct as a field.

#### Code:
```go
package structures

import "fmt"

// Name struct with First and Last name fields
type Name struct {
	Fname string
	Lname string
}

// Address struct with Flat, Residence, and a Name field (nested struct)
type Address struct {
	Flat      int
	Residence string
	Name      Name
}

// DisplayAddress function to print address details
func DisplayAddress(address Address) {
	fmt.Println("Address Details:")
	fmt.Printf("Plot: %d, Residence: %s\n", address.Flat, address.Residence)
	fmt.Printf("Name: %s %s\n", address.Name.Fname, address.Name.Lname)
}
```

#### Key Concept:
- **Nested Structs**: The `Address` struct contains a `Name` struct, which is an example of how to work with nested structs in Go.
- **Display Function**: A function to print the details of an `Address`.

---

## How to Use


1. Open a terminal and navigate to the project folder.
2. Run the program using the command:
   ```
   go run main.go
   ```




