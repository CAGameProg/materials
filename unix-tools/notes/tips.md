# Tips

### os.Args

Remember that to access the command line inputs, you must
use the `os.Args` array. This is an array of strings with
all the inputs. For example, if I run the following command:

```
$ ./gocat file.txt
```

The `os.Args` array will contain: `["./gocat", "file.txt"]`.

### go build

Since these programs take command line arguments, they cannot
be run easily with `go run` as we have done so far. Instead you
must compile the program first with `go build` and then execute
the executable file that is produced.

For example, if I have written `gocat.go`, I can compile it with

```
$ go build gocat.go
```

This should produce a file called `gocat` or `gocat.exe` (on Windows).
You can then run that file with

```
$ ./gocat input1
```

### errors

Many functions dealing with the file system return errors (for example
if the file you are trying to look at does not exist). If you do not
want to deal with the error, you can put an underscore (`_`) instead
of a variable name. This is bad practice but is OK for now.

For example, the `ioutil.ReadDir(directory)` function returns a list of
files in the directory as well as a possible error (if the directory does
not exist).

You can check it like so:

```go
files, err := ioutil.ReadDir(".")

if err != nil {
    fmt.Println(err)
    // try to salvage the program
    // Or just exit with os.Exit(0)
}
```

Or you can ignore it like so:

```go
files, _ := ioutil.ReadDir(".")
```
