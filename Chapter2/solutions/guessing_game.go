package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	const max = 100
	const min = 1

	rand.Seed(time.Now().UTC().UnixNano())

	secretNum := rand.Intn((max+1)-min) + min
	fmt.Println("I'm thinking of a number between 1 and 100")

	for {
		fmt.Print("What's your guess? ")

		var guess int
		fmt.Scanf("%d", &guess)

		if guess == secretNum {
			fmt.Println("You got it!")
			break
		} else if guess < secretNum {
			fmt.Println("Too low!")
		} else if guess > secretNum {
			fmt.Println("Too high")
		}
	}
}
