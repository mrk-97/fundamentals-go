package main

import "fmt"

// Grade Book
func addGrade(book map[string][]int, name string, grade int) {

	// check if the student already exists in book
	_, exists := book[name]
	if !exists {
		fmt.Printf("No student found, Added: %s\n", name)
		grades := []int{grade}
		book[name] = grades
	} else {
		book[name] = append(book[name], grade)
	}
}

func calculateAverage(book map[string][]int, name string) float64 {
	_, exists := book[name]
	sumOfGrades := 0
	if exists {
		grades := book[name]
		for _, v := range grades {
			sumOfGrades += v
		}
		return float64(sumOfGrades) / float64(len(grades))
	} else {
		return 0
	}
}

func main() {
	gradebook := make(map[string][]int)

	addGrade(gradebook, "Alice", 90)
	addGrade(gradebook, "Bob", 80)
	addGrade(gradebook, "Alice", 85)

	// Calculate and print Alice's average
	aliceAvg := calculateAverage(gradebook, "Alice")
	fmt.Printf("Alice's average: %.2f\n", aliceAvg)

	// Print all students and their grades
	fmt.Println("--- Full Gradebook ---")
	for name, grades := range gradebook {
		fmt.Printf("%s: %v\n", name, grades)
	}
}
