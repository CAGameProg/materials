package main

import "fmt"
import "time"

var done chan bool

func foo(delay int) {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
		time.Sleep(time.Duration(delay) * time.Second)
	}

	done <- true
}

func main() {
	done = make(chan bool)

	go func() {
		time.Sleep(5 * time.Second)
		fmt.Println("Done waiting 5 seconds")
	}()

	go foo(3)
	go foo(1)
	<-done
}
