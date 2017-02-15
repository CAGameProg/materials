package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func searchDir(dir, searchName string) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if searchName == file.Name() {
			fmt.Println(dir + "/" + file.Name())
		}

		if file.IsDir() {
			searchDir(dir+"/"+file.Name(), searchName)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments")
		return
	}

	searchDir(".", os.Args[1])
}
