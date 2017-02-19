package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func searchFile(filename, searchStr string) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	txt := string(data)
	lines := strings.Split(txt, "\n")
	filenamePrinted := false

	for i, l := range lines {
		if strings.Contains(l, searchStr) {
			if !filenamePrinted {
				fmt.Println(filename)
				filenamePrinted = true
			}

			fmt.Print(i + 1)
			fmt.Print(":")
			fmt.Println(l)
		}
	}

	return nil
}

func searchDir(directory, searchStr string) error {
	files, err := ioutil.ReadDir(directory)

	for _, file := range files {
		if file.IsDir() {
			err = searchDir(directory+"/"+file.Name(), searchStr)
		} else {
			err = searchFile(directory+"/"+file.Name(), searchStr)
		}

		if err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments")
		return
	}

	searchStr := os.Args[1]

	err := searchDir(".", searchStr)
	if err != nil {
		fmt.Println(err)
	}
}
