You might want to take a look at [tips.md](notes/tips.md).

# 1. gocat.go

Write a simple clone of the `cat` unix tool. The program should take as input
the name of a file, and then should print the contents of the file.

Example:

```
$ gocat hello.go
package main

func main() {
    fmt.Println("Hello world")
}
```

Part of the challenge of this problem is figuring out which functions and packages
to use from the standard library.

Hint: you should look in the [ioutil](https://golang.org/pkg/io/ioutil/) package.

If you wanted to call, for example, `ReadDir` from the `ioutil` package, you
would do the following:

```go
files, _ := ioutil.ReadDir(directory)

// Do something with the files ...
```

Also recall that an array of bytes (`[]byte`) is the same as a string, and you
can cast to convert between the two.

Don't overthink this program. It can be done using Go's ioutil package in about
25 lines of code or less. You don't need to print the file line by line, you
can just read the entire contents and print that.

# 2. gofind.go

Write a command line tool that will search starting in the current directory 
for a file and print its path if it finds it.

The program should search recursively through all the directories in the current
directory.

For example if my current file structure looks like this:

```
main.go
file1.txt
file2.txt
gopath/
    src/
        program1/
            main.go
        program2/
            main.go
        program3/
            program3.go
    pkg/
    bin/
```

And I search for `main.go` by running:

```
$ gofind main.go
```

Then the program should print:

```
./main.go
./gopath/src/program1/main.go
./gopath/src/program2/main.go
```

This program is similar to the `ls` program we made in class (located [here](notes/ls.go)), but now you have
to check the name of the file before printing it, and you must also search
any sub-folders.

Also you'll want to use the `IsDir()` function from the `FileInfo` interface.

Hint:

Since you have to search sub-folders, and also their sub-folders, and on
and on, this is a recursive problem.

I recommend making a `searchDir` function to search a given directory:

```go
// dir is the directory we are searching and
// searchName is the name of the file we are searching for
func searchDir(dir string, searchName string) {
    // get all the files in the directory

    for _, f := range files {
        // Check if the file is a directory
        // If the file is a directory, then search it
        // with searchDir(dir + NameOfSubDirectory, searchName) // The searchName is still the same

        // If the file is not a directory, then check if
        // its name matches the name we are searching for (the searchName)
        // If it matches, print out dir + the name of the file
    }
}
```

Remember that `f.Name()` only gives the name of the file, not its entire path. To get the path,
in the searchDir function, you can append `dir` to `f.Name()` with `dir + f.Name()`.

Then in your main function, you can just call `searchDir(".")` to search the current
directory and all sub-directories as well.

# 3. Bonus: gogrep.go

`grep` is a unix tool used to search across many files for a certain word or phrase.

`gogrep` should be very similar to `gofind` except instead of searching for a certain
filename, it should search inside files for a certain string. The program should
check each line of every file, and if it contains the string, it should print filepath of the file, 
the line number and contents of the line the string was found on.

Hint: I recommend you use the `strings.Contains` function ([here](https://golang.org/pkg/strings/#Contains)) to check if a line contains a substring.

You may also want to use the `strings.Split` function ([here](https://golang.org/pkg/strings/#Split)) to split the file's contents into an array of strings (split on the `\n` character).

# 4. Bonus: Extending gogrep.go

There are a number of ways of extending the gogrep program. A good challenge is
adding support for regular expressions. Regexes provide a way to match a number
of different strings with one expression.

For example, in the regular expression `a.c`, the `.` is a special character
which matches any character. So `a.c` will match the string `abc`, or `aac` or
`a#c`.

There are many special characters in regex and you can learn more about them
(and specifically the specification Go uses) [here](https://github.com/google/re2/wiki/Syntax).

That page may seem dense so I recommend checking out [this website](https://regex-golang.appspot.com/assets/html/index.html).

Luckily, go has a regular expression engine built into the standard library.
You can see it at [golang.org/pkg/regexp/](https://golang.org/pkg/regexp/).

To compile a regex, run:

```go
r, err := regexp.Compile("a.c")
// r is type *regexp.Regexp
```

Then you can check if it matches strings with

```go
str1 := "abc"
str2 := "bac"
str3 := "aac"

fmt.Println(r.Match([]byte(str1))) // Prints true
fmt.Println(r.Match([]byte(str2))) // Prints false
fmt.Println(r.Match([]byte(str3))) // Prints true
```

Try extending your gogrep program so that you can search with regular expressions
and not just normal strings.
