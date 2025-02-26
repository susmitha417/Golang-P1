

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

# Go Goroutines and Channels: 

## 1. Goroutines

Goroutines are lightweight threads managed by the Go runtime. They enable functions or methods to run concurrently, allowing the Go program to perform multiple tasks simultaneously.

### Creating a Goroutine
To create a goroutine, you use the `go` keyword before calling a function. This will start the function as a separate goroutine that runs concurrently with the main program.

### Syntax:
```go
go functionName()
```

### Example:
```go
package main

import "fmt"

func sayHello() {
    fmt.Println("Hello from Goroutine!")
}

func main() {
    go sayHello()  // Start sayHello function as a goroutine

    // Allow goroutine to complete before exiting the main function
    fmt.Scanln()  // Block the main function until Enter is pressed
}
```

### Explanation:
- `go sayHello()` launches the `sayHello` function as a goroutine, running concurrently with the rest of the program.
- To prevent the main program from exiting immediately (before the goroutine finishes), we use `fmt.Scanln()` to block the main goroutine, allowing `sayHello` to complete its execution.

## 2. Channels

Channels are a way for goroutines to communicate with each other and synchronize their execution. A channel allows you to send and receive values between goroutines. Channels are typed, meaning they can carry only one type of data (e.g., `chan int`, `chan string`).

### Syntax to Create a Channel:
```go
ch := make(chan int)  // Creates an unbuffered channel of type int
```

### Sending Data to a Channel:
```go
ch <- value  // Send data to channel ch
```

### Receiving Data from a Channel:
```go
value := <-ch  // Receive data from channel ch
```

### Example of Using a Channel:
```go
package main

import "fmt"

func add(a, b int, ch chan int) {
    result := a + b
    ch <- result  // Send result to the channel
}

func main() {
    ch := make(chan int)  // Create a channel

    go add(3, 4, ch)  // Launch add function as a goroutine

    result := <-ch  // Receive result from the channel
    fmt.Println("Result from goroutine:", result)
}
```

### Explanation:
- A channel `ch` is created to communicate between the main goroutine and the `add` goroutine.
- The `add` function calculates the sum and sends the result to the channel `ch`.
- The main function waits for the result from the channel and prints it.

---

## Types of Channels:

### 1. **Unbuffered Channels**

An **unbuffered channel** has no capacity to hold data. The sending goroutine is blocked until another goroutine receives the value. Similarly, if a goroutine tries to receive from an unbuffered channel but no value has been sent, it will block until a value is available.

#### Characteristics:
- **Blocking**: Both sending and receiving operations block until data is exchanged.
- **Use Case**: Used when strict synchronization is needed between the sender and receiver (i.e., sender waits for receiver and vice versa).

#### Syntax to Create an Unbuffered Channel:
```go
ch := make(chan int)  // Creates an unbuffered channel of type int
```

#### Example of Unbuffered Channel:
```go
package main

import "fmt"

func main() {
    ch := make(chan int)  // Unbuffered channel

    go func() {
        fmt.Println("Sending value 42")
        ch <- 42  // Send value to channel (this will block until received)
        fmt.Println("Sent value 42")
    }()

    fmt.Println("Receiving value from channel")
    val := <-ch  // Receive value (this will block until data is sent)
    fmt.Println("Received value:", val)
}
```

### Explanation:
- The sender goroutine is blocked on `ch <- 42` until the receiver receives the value using `<-ch`.
- Similarly, the receiver is blocked until the sender sends the value.

---

### 2. **Buffered Channels**

A **buffered channel** has a capacity to hold a certain number of values before blocking the sender. The sender will only be blocked when the channel reaches its capacity. Once there is space in the buffer (i.e., the receiver takes a value), the sender can continue sending values. The receiver can receive values without blocking until the buffer is empty.

#### Characteristics:
- **Non-blocking until full**: The sender can send data into the channel without blocking until the buffer is full.
- **Use Case**: Useful when you want to decouple the sender and receiver, allowing the sender to continue sending data even if the receiver is not immediately available (up to the buffer's capacity).

#### Syntax to Create a Buffered Channel:
```go
ch := make(chan int, 2)  // Creates a buffered channel with capacity of 2
```

#### Example of Buffered Channel:
```go
package main

import "fmt"

func main() {
    ch := make(chan int, 2)  // Buffered channel with capacity 2

    // Send values into the channel without blocking
    ch <- 10  // Sent into the buffer
    ch <- 20  // Sent into the buffer

    // At this point, the buffer is full. Any further sends would block.

    fmt.Println("Receiving values from channel:")
    fmt.Println(<-ch)  // Received 10
    fmt.Println(<-ch)  // Received 20
}
```

### Explanation:
- The buffered channel can hold 2 values before blocking the sender.
- In this example, the first two `ch <-` operations send values into the channel without blocking because the buffer has space for two values.
- The receiver can then receive values from the channel, and the sender can continue sending once space becomes available in the buffer.

---

## When to Use Which?

### Use an **Unbuffered Channel** when:
- You want strict synchronization between the sender and receiver.
- Communication between goroutines should happen immediately (i.e., sender waits for the receiver and vice versa).
- Typically used when you need to ensure that the sender and receiver are synchronized at the exact moment of communication.

### Use a **Buffered Channel** when:
- You want the sender to send data without waiting for the receiver immediately (up to the buffer's capacity).
- You need to decouple the sender and receiver, allowing the sender to continue working even if the receiver hasn't processed the previous messages yet.
- Often used in cases where the sender produces data at a different rate than the receiver can consume.

---

## Example with Both Buffered and Unbuffered Channels:

```go
package main

import "fmt"

func main() {
    // Unbuffered channel example
    unbufferedCh := make(chan int)
    go func() {
        unbufferedCh <- 1  // This will block until received
        fmt.Println("Sent 1 to unbuffered channel")
    }()
    fmt.Println(<-unbufferedCh)  // This will unblock the sender
    // Output: Sent 1 to unbuffered channel

    // Buffered channel example
    bufferedCh := make(chan int, 2)
    bufferedCh <- 1  // No blocking
    bufferedCh <- 2  // No blocking
    fmt.Println(<-bufferedCh)  // Receive 1
    fmt.Println(<-bufferedCh)  // Receive 2
}
```

### Summary:
- **Unbuffered channels** block both the sender and receiver until data is exchanged.
- **Buffered channels** allow the sender to send multiple values without blocking until the buffer is full.

## How to Use


1. Open a terminal and navigate to the project folder.
2. Run the program using the command:
   ```
   go run main.go
   ```




