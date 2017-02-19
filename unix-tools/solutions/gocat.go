package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments")
		return
	}

	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print(string(data))
}
