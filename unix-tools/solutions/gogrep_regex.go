package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	"github.com/fatih/color"
)

func isDirectory(path string) (bool, error) {
	fileInfo, err := os.Stat(path)
	return fileInfo.IsDir(), err
}

func searchFile(filename string, regex *regexp.Regexp) error {
	data, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	txt := string(data)
	lines := strings.Split(txt, "\n")
	filenamePrinted := false

	for i, l := range lines {
		matches := regex.FindAllString(l, -1)
		if matches != nil {
			if !filenamePrinted {
				color.Set(color.FgGreen, color.Bold)
				fmt.Println(filename)
				color.Unset()
				filenamePrinted = true
			}

			color.Set(color.FgYellow)
			fmt.Print(i + 1)
			color.Unset()
			fmt.Print(": ")

			for _, match := range matches {
				yellow := color.New(color.FgBlack, color.BgYellow).SprintFunc()
				l = strings.Replace(l, match, yellow(match), -1)
			}

			fmt.Println(l)
		}
	}

	return nil
}

func searchDir(directory string, searchRegex *regexp.Regexp) error {
	files, err := ioutil.ReadDir(directory)

	for _, file := range files {
		if file.IsDir() {
			err = searchDir(directory+"/"+file.Name(), searchRegex)
		} else {
			err = searchFile(directory+"/"+file.Name(), searchRegex)
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

	regex, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	if len(os.Args) < 3 {
		err := searchDir(".", regex)
		if err != nil {
			fmt.Println(err)
		}
	} else {
		for _, filename := range os.Args[2:] {
			isDir, err := isDirectory(filename)
			if err != nil {
				fmt.Println(err)
				continue
			}

			if isDir {
				err = searchDir(filename, regex)
			} else {
				err = searchFile(filename, regex)
			}

			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
