package main

import (
	"errors"
	"fmt"
)

func calculate(a, b int) (int, error) {
	result, err := devide(a, b)
	if err != nil {
		return 0, err
	}
	return result + 10, nil
}

func devide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Cannot divide by zero")
	}
	return a / b, nil
}

func main() {
	x, y := 10, 2
	c, d := 10, 0

	fmt.Printf("\n ---- Test Case 1 ----\n")

	result, err := calculate(x, y)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("result:", result)
	}

	fmt.Printf("\n ---- Test Case 2 ----\n")

	result1, err := calculate(c, d)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Print("result:", result1)
	}

}
