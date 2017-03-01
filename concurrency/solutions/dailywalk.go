package main

import (
	"fmt"
	"math/rand"
	"time"
)

var done chan bool
var alarm chan bool

func Random(min, max int) int {
	return rand.Intn(max-min) + min
}

func GetReady(name string) {
	fmt.Println(name, "has started getting ready.")
	n := Random(10, 20)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(name, "spent", n, "seconds getting ready")

	done <- true
}

func ArmAlarm() {
	fmt.Println("Arming alarm")
	time.Sleep(20 * time.Second)
	fmt.Println("Alarm is armed")
	alarm <- true
}

func PutShoes(name string) {
	fmt.Println(name, "started putting on shoes")
	n := Random(5, 15)
	time.Sleep(time.Duration(n) * time.Second)
	fmt.Println(name, "spent", n, "seconds putting on shoes")

	done <- true
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	alarm = make(chan bool)
	done = make(chan bool, 2)

	fmt.Println("Let's go for a walk!")

	go GetReady("Alice")
	go GetReady("Bob")

	<-done
	<-done

	go ArmAlarm()

	go PutShoes("Alice")
	go PutShoes("Bob")

	<-done
	<-done

	fmt.Println("Exiting and locking the door")

	<-alarm
}
