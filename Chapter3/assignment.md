# Arrays and Functions

## 1. average.go

Write a function called `average` which takes a slice of numbers and returns
their average. The function signature would be `average([]float64) float64`

## 2. maxmin.go

Write two functions `max` and `min` which take a variadic amount of integers,
and return the largest and smallest integer respectively.

`min(nums ...int) int`: returns the smallest number in the given parameters.
`max(nums ...int) int`: returns the largest number in the given parameters.

Note: Here are definitions for the largest and smallest possible values of an integer:

```
const MaxInt = 2147483647
const MinInt = -2147483647
```

## 3. fibonacci.go

The Fibonacci sequence is defined as: `fib(0) = 0`, `fib(1) = 1`, `fib(n) = fib(n-1) + fib(n-2)`.

Here is the beginning of the sequence: 0, 1, 1, 2, 3, 5, 8, 13, 21

a. Write a recursive function `fibR(n int) int` which can find `fib(n)`.

b. Write an iterative (using for loops not recursion) function `fibIter(n int) int` which can find `fib(n)`.

You can try comparing the speed of these two functions for computing fib(40) for example. You should see that the iterative version is much faster.

## 4. Bonus: hangman.go

Write a hangman program. Here is some example output:

```
-----------
What's your guess? a
Correct!

-----a-----
What's your guess? e
Incorrect!

-----a-----
What's your guess? m
Correct!

-----amm---
What's your guess? o
Correct!

--o--amm---
What's your guess? h
Incorrect!

--og-amm--g
What's your guess? g
Correct!

...

-rogramming
What's your guess? p
Correct!

You win! The word was programming!
```

Here are some tips to help you out:

To generate a random number within a range from [min, max), use the following function:

```go
func random(min, max int) int {
    return rand.Intn(max - min) + min
}
```

And place the line `rand.Seed(time.Now().UTC().UnixNano())` at the start of the program
(don't forget to import `math/rand` and `time`).

To find the word, I would recommend simply making a short array of words the program can choose from
and then picking randomly from the array.

Make sure to make a separate array that will hold the dashes/partly made word. You can tell if the player
has won the game if there are no more dashes in the array.

Remember that to get user input as a rune and put it in the variable `guess` use `fmt.Scanf("%c\n", &guess)`.
