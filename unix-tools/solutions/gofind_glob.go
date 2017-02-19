package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/zyedidia/glob"
)

func searchDir(dir string, findRegex *glob.Glob) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, file := range files {
		if findRegex.Match([]byte(file.Name())) {
			fmt.Println(dir + "/" + file.Name())
		}
		if file.IsDir() {
			searchDir(dir+"/"+file.Name(), findRegex)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Too few arguments")
		return
	}

	findRegex, err := glob.Compile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	dir := "."
	if len(os.Args) > 2 {
		dir = os.Args[2]
	}

	searchDir(dir, findRegex)
}
