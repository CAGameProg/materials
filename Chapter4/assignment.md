# 1. person.go

Write a `Person` struct which will store a person's information. It should specifically store
the person's first name, last name, age, and email address.

Now write a `Describe` function which takes a pointer to a `Person` object and outputs a
description of them:

```go
func Describe(p *Person) {
    // ...
}
```

For example:

```go
john := &Person{"John", "Smith", 27, "jsmith@gmail.com"}
Describe(john)
```

The output of this program should be:

```
John Smith is 27 years old. For more information, you can contact John at jsmith@gmail.com
```

# 2. circle.go

a. Write code for a `Circle` struct. This struct should store the circle's
x and y coordinates (as `float64`s) as well as the circle's radius (also `float64`).

b. Write a function called `CircleArea` that takes as input a pointer to a Circle
and returns the circle's area (remember circle area can be calculated with: pi*r*r).
You can use `const Pi = 3.14` or `math.Pi` for a more exact constant:

```go
func CircleArea(c *Circle) float64 {
    // ...
}
```

c. Write a function called `MakeUnitCircle` which will simply return a pointer to a Circle with
radius 1 and position (0, 0):

```
func MakeUnitCircle() *Circle {
    // ...
}
```

# 3. cards.go

Write a struct for a `Card`. It should store the card's suit (hearts, diamonds, clubs, or spades)
as well as the card's value (1-13, where 1 is Ace, 11 is Jack, 12 is Queen, and 13 is King).

You may want to store the suit as an integer where, for example, 0 is hearts, 1 is diamonds, 2 is clubs, and 3 is spades.

Write a function called `DescribeCard` which will take a pointer to a card and return a string:

```go
func DescribeCard(c *Card) string{
    // ...
}
```

Here is example output:

```go
card := &Card{3, 1}
s := DescribeCard(card)
fmt.Println(s) // This should print "3 of diamonds"
```

Now write a function called `MakeDeck` which will return an array of 52 cards (no shuffling necessary):

```go
func MakeDeck() []Card {
    // ...
}
```

Write a function called `PickRandomCard` which should take a Deck (an array of Cards) and return a random card. That card should also be removed from the deck.

Hints: here's how to remove an element of an array at index `i`:
```go
a = append(a[:i], a[i+1:]...)
```

Note: arrays are pointers so you can modify them directly in a function just like you can a pointer.

Here's the code for generating a random number between 0 and n:
```go
rand.Intn(n)
```

You'll need to seed the random number generator at the top of the program. Just put this line at the start of the main function.

```go
rand.Seed(time.Now().UTC().UnixNano())
```
(also import "math/rand" and "time")

## Bonus: Shuffle

Another method of getting a random card is to shuffle the deck and then pick the top card.

Write a function called `Shuffle` which takes an array of cards and randomly shuffles them.
The function should return the new shuffled array:

```go
func Shuffle(deck []Card) []Card {
    // ...
}
```

Hint: The best way to shuffle an array is to use the `rand.Perm` function (make sure to import `math/rand`).
You use the function like so:

```go
rand.Seed(time.Now().UTC().UnixNano())

array := []int{1, 2, 3, 4, 5, 6}
newArray := make([]int, len(array))
perm := rand.Perm(len(array))

for i := 0; i < len(perm); i++ {
    newArray[i] = array[perm[i]]
}

fmt.Println(newArray) // This should print a shuffled version of the `array` variable
```

The way this works is that `rand.Perm` gives you a slice that contains random numbers. Then you
assign each position in the `newArray` a random position from the original `array` variable.
