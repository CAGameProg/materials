package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	words := []string{"programming", "word", "dog", "array", "game"}
	secretWord := words[random(0, len(words))]

	dashes := make([]rune, len(secretWord))
	for i := range secretWord {
		dashes[i] = '-'
	}

	for {
		for _, c := range dashes {
			fmt.Print(string(c))
		}
		fmt.Println()

		fmt.Print("What's your guess? ")
		var guess rune
		fmt.Scanf("%c\n", &guess)

		correct := false
		for i, c := range secretWord {
			if c == guess {
				dashes[i] = c
				correct = true
			}
		}
		if correct {
			fmt.Println("Correct!")
		} else {
			fmt.Println("Incorrect!")
		}

		win := true
		for _, c := range dashes {
			if c == '-' {
				win = false
			}
		}

		if win {
			fmt.Printf("You win! The word was %s!\n", secretWord)
			break
		}
	}
}
