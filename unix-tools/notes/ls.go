package main

import (
	"fmt"
	"os"
)

func main() {
	dir := "."
	if len(os.Args) >= 2 {
		dir = os.Args[1]
	}

	file, err := os.Open(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	files, err := file.Readdir(0)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < len(files); i++ {
		fName := files[i].Name()
		fmt.Println(fName)
	}
}
