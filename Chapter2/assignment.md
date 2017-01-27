# Control Structures and Loops

## 1. congress.go

Create a Go program that takes as input the age of a candidate and years of U.S. citizenship and outputs whether the candidate is eligible for the House of Representatives or Senate. 

You must be at least 25 and have been a US citizen for at least 7 years to be eligible for the House of Representatives, and you must be at least 30 and have been a US citizen for at least 9 years to be eligible for the Senate.

Example 1:

```
Enter age of candidate: 37
Enter years of U.S. citizenship: 3
 
The candidate is NOT eligible for election to either the House of Representatives or the Senate.
```

Example 2: 
 
```
Enter age of candidate: 47
Enter years of U.S. citizenship: 8
 
The candidate is eligible for election to the House of Representatives but is NOT eligible for election to the Senate.
```

Example 3: 
 
```
Enter age of candidate: 35
Enter years of U.S. citizenship: 25
 
The candidate is eligible for election to both the House of Representatives and the Senate.
```

## 2. donor.go

A certain charity designates donors who give $1,000.00 or more as Benefactors; those who give $500.00 to $999.99 as Patrons; those who give $100.00 to $499.99 as Supporters; those who give $15.00 to $99.99 as Friends; and those who give less than $15 as Cheapskates.  
 
Write a program containing a nested if-else statement (or statements) that, given a keyboard input with the amount of a contribution, outputs the correct designation for that contributor. Here’s an example that illustrates what your program should look like in action: 

```
Enter the amount of a contribution: $5.20
You are a cheapskate!
```
 
The alternative printouts are: 

You are a patron!
You are a supporter!
You are a benefactor!
You are a friend!

## 3. fizzbuzz.go

Write a program that prints the numbers from 1 to 100. But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz". For numbers which are multiples of both three and five print "FizzBuzz".

Remember that the modulus operator (`%`) gives the remainder between the division of two numbers.
You can use this to see if one number is divisible by another. For example, `15 % 3` is `0` which means
that `15` is divisible by `3`.

For example here is the output from 1 to 17:

```
1
2
Fizz
4
Buzz
Fizz
7
8
Fizz
Buzz
11
Fizz
13
14
FizzBuzz
16
17
```

## 4. perfect_squares.go

Write a program that will calculate the number of perfect squares from 1 up to and including a certain number (1 is a perfect square).

A perfect square is an integer which has an integer square root. Here are a few perfect squares:

1, 4, 9, 16, 25, 36...

Here an example of the program (where the user enters 1000):

```
Enter an integer: 1000
There are 31 perfect squares between 1 and 1000
```

## 5. Bonus: guessing_game.go

Write a random number guessing game.

Example output:

```
I'm thinking of a number between 1 and 100

What is your guess? 50
Too low!
What is your guess? 75
Too high!
What is your guess? 63
Too high!
What is your guess? 60
Too low!
What is your guess? 62
You got it!
```

Here are some tips to help you out:

To generate a random number (integer) within a range from [min, max), use the following statement:

```go
rand.Intn(max - min) + min
```

And place the line `rand.Seed(time.Now().UTC().UnixNano())` at the start of the main function
(and don't forget to import `math/rand` and `time`).
