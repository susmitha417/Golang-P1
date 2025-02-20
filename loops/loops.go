package loops

import "fmt"

func StudentMarks() {

	var student1, student2, student3, student4 int

	// Taking input from the user
	fmt.Println("Enter marks for Student 1:")
	fmt.Scan(&student1)

	fmt.Println("Enter marks for Student 2:")
	fmt.Scan(&student2)

	fmt.Println("Enter marks for Student 3:")
	fmt.Scan(&student3)

	fmt.Println("Enter marks for Student 4:")
	fmt.Scan(&student4)

	// Function to check the grade and feedback
	checkGrade(student1, 1)
	checkGrade(student2, 2)
	checkGrade(student3, 3)
	checkGrade(student4, 4)
}

// Function to check the grade and give feedback
func checkGrade(marks int, studentNumber int) {
	var grade string

	// Assigning grades based on marks
	if marks >= 90 {
		grade = "A"
	} else if marks >= 80 {
		grade = "B"
	} else if marks >= 70 {
		grade = "C"
	} else if marks >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}

	// Printing student number, marks, and grade
	fmt.Printf("Student %d - Marks: %d, Grade: %s - ", studentNumber, marks, grade)

	// Giving feedback using switch
	switch grade {
	case "A", "B", "C", "D":
		fmt.Println("Good job! Keep it up.")
	default:
		fmt.Println("Needs improvement. Don't give up!")
	}

	// Printing memory address of the marks
	fmt.Printf("Memory address of Student %d marks: %p\n\n", studentNumber, &marks)
}
