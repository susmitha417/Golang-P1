package array

import "fmt"

// Function to take student scores as input and print them
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
