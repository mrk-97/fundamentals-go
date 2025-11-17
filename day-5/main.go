package main

import "fmt"

type Owner struct {
	FirstName string
	LastName  string
	Age       int
}

type Car struct {
	Make  string
	Model string
	Year  int
	Owner Owner
}

func (c Car) PrintDetails() {
	fmt.Printf("Car Details:\n")
	fmt.Printf("Make: %s\n", c.Make)
	fmt.Printf("Model: %s\n", c.Model)
	fmt.Printf("Year: %d\n", c.Year)
	fmt.Printf("Owner: %s %s (Age: %d)\n", c.Owner.FirstName, c.Owner.LastName, c.Owner.Age)
}

func (c *Car) UpdateYear(newYear int) {
	c.Year = newYear
}

func main() {

	myCar := Car{
		Make:  "Toyota",
		Model: "Innova",
		Year:  2025,
		Owner: Owner{
			FirstName: "John",
			LastName:  "Doe",
			Age:       30,
		},
	}

	fmt.Println("=====================Before Update=============================")
	myCar.PrintDetails()

	myCar.UpdateYear(2026)
	fmt.Println("=====================After Update=============================")
	myCar.PrintDetails()

}
