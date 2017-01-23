# Data Types and Variables

## 1. output.txt

a. Please write precisely what the following statements will print:

* `fmt.Println(5 + 3)`
* `fmt.Println("5 + 3")`

b. What does `fmt.Println(5 + '3')` print, and why does that compile but
`fmt.Println(5 + "3")` doesn't?

## 2. fahrenheit.go

The Fahrenheit to Celsius conversion is given by the following equation:
`C = (F - 32) * 5/9`.

Write a program that asks the user for a temperature in Fahrenheit and
prints out the temperature converted to Celsius.

Make sure that the user input is a float not an int because conversions
from Fahrenheit to Celsius often involve decimals (float64).

## 3.

a. An unsigned byte is 8 bits therefore the largest (unsigned) integer it can store (in binary) is `11111111` which is 255 in decimal.
What is the largest number in decimal that a 10 bit integer type can store?

b. Create a program that asks the user for a number of bits and prints out
what the maximum decimal number an unsigned integer with that size could store.

Hint: to raise a variable to a power in Go (i.e. `a^b`), you must use the `math.Pow(a, b)` function. Make sure to convert `a` and `b` to float64s first and to import the `math` package. See the `notes.md` file for an example.

c. Why is the maximum value of a signed 32 bit integer 2,147,483,647 and
not 4,294,967,295 (which would be the max for an unsigned 32 bit integer as shown calculated by the program above).
