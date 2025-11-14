package main

import "fmt"

func main() {
	ages := make(map[string]int)

	ages["Alice"] = 30
	ages["Bob"] = 25
	ages["Charlie"] = 40

	for k, v := range ages {
		fmt.Printf("Key: %s, Value: %d\n", k, v)
	}

}
