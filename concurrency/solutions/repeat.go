package main

import (
	"fmt"
	"time"
)

func Repeat(f func(), n time.Duration) {
	go func() {
		for {
			f()
			time.Sleep(n)
		}
	}()
}

func SayHello() {
	fmt.Println("Hello")
}

func SayGoodbye() {
	fmt.Println("Goodbye")
}

func main() {
	Repeat(SayHello, 2*time.Second)
	Repeat(SayGoodbye, 1*time.Second)

	time.Sleep(10 * time.Second)
}
