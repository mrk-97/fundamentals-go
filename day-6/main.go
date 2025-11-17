package main

import (
	"fmt"
	"math"
)

type Shape interface{
	area() float64
}

type Rectangle struct{
	Width float64
	Length float64
}

type Circle struct{
	Radius float64
}

func (r Rectangle) area() float64{
	return r.Width * r.Length
}

func (c Circle) area() float64{
	return math.Pi * c.Radius * c.Radius
}

func printArea(s Shape){
	fmt.Println("The area of this shape is:",s.area())
}

func main() {

	
	
	r := Rectangle{Length: 10, Width: 5}
	c := Circle{Radius: 3}

	shapes := []Shape{}

	shapes = append(shapes,r)
	shapes = append(shapes,c)

	r2 := Rectangle{Length:7,Width:7}
	
	shapes = append(shapes,r2)

	for _,shape := range shapes{
		fmt.Println(shape.area())
	}
}