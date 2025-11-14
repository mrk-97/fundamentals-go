// // Rough version
// package main

// import "fmt"

// func main() {
// 	scores := []int{10, 20, 30, 40, 50}

// 	sum := 0
// 	for _, score := range scores {
// 		sum = sum + score
// 	}

//		fmt.Println("sum:", sum)
//	}
//
// Go's version
package main

import "fmt"

func sumOfScores(scores []int) int {
	sum := 0
	for _, score := range scores {
		sum = sum + score
	}
	return sum
}
func deleteElement(s []string, i int) []string {
	s = append(s[:i], s[i+1:]...)
	return s
}

func main() {
	scores := []int{10, 20, 30, 40, 50}

	fmt.Println("Total Score:", sumOfScores(scores))

	letters := []string{"a", "b", "c", "d", "e"}

	fmt.Println(letters)
	letters = deleteElement(letters, 2)
	fmt.Println("after deleting index at 2:", letters)
	fmt.Println("lenth of letters slice:", len(letters))
	fmt.Println("capacity of letters slice:", cap(letters))

	original := []int{10, 20, 30}

	duplicate := make([]int, len(original))
	copy(duplicate, original)

	duplicate[0] = 999

	fmt.Println("original:", original)
	fmt.Println("duplicate:", duplicate)
}
