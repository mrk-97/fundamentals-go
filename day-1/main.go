package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)

	secretNumber := r.Intn(10) + 1

	guessCount := 1

	fmt.Println("Guess a number between 1-10:")
	for {
		var guess int
		fmt.Scanln(&guess)
		fmt.Printf("Remaining Guesses : %d\n", 3-guessCount)
		if guess == secretNumber {
			switch guessCount {
			case 1:
				fmt.Println("Wow! First try! Are you psychic?!")
			case 2:
				fmt.Println("Nice! You got it in two tries.")
			case 3:
				fmt.Println("Phew! Got it on your last try.")
			}
			if guessCount > 3 {
				fmt.Println("Retry limit exceeded. You lose")
				break
			}
			break
		} else if guess < secretNumber {
			fmt.Println("Your Guess was too low, try again")
		} else {
			fmt.Printf("Your Guess was too high, try again")
		}
		guessCount++
	}
}
