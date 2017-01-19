# Notes

This file will provide you with some code examples which will be helpful to complete
the homework.

## Hello world

Here is simply a "Hello world" program written in Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello world")
}
```

## Getting user input

Here is an example to read user input into a variable:

```go
package main

import "fmt"

func main() {
    fmt.Print("Enter a number: ")

    // We define the variable but it doesn't yet have a value
    var inputNum int
    // We read the variable from the user
    // Don't worry about what the '&' sign means or how this works
	fmt.Scanf("%d", &inputNum)

    fmt.Println("You entered the number:", inputNum)
}
```

## Raising a number to a power

For example, to calculate 2^12, you would write

```go
package main

import "math"
import "fmt"

func main() {
    x := math.Pow(2.0, 12.0)
    fmt.Println(x) // This will print 4096
}
```
