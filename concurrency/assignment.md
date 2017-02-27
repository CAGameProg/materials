# Concurrency and Networking

## 1. repeat.go

Write a function called `Repeat` which can be used to repeatedly
call another function every `n` milliseconds in a separate goroutine.

For example:

```go
func Repeat(f func(), delay time.Duration) {
    // Your implementation here
}

func SayHello() {
    fmt.Println("Hello")
}

func SayGoodbye() {
    fmt.Println("Goodbye")
}

func main() {
    Repeat(SayHello, 2 * time.Second)
    Repeat(SayGoodbye, 1 * time.Second)

    time.Sleep(10 * time.Second)
}
```

This program will print `Hello` every two seconds and `Goodbye` every second for
ten seconds.

## 2. fetch.go

Here is an implementation of a `fetch` function:

```go
func fetch(url string, ch chan string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("Error: %v\n", err)
		return
	}
	resp.Body.Close()
	ch <- fmt.Sprintf("%.2fs %s\n", time.Since(start).Seconds(), url)
}
```

This function takes two arguments: a `url` which is a website's address,
and a string channel.

The function downloads the contents of the website and stores the url and amount of
time it took to download the website in the channel.

Use this function to write a program which takes multiple command line arguments
and downloads all of them simultaneously. It should print how long it took to
download each website (this info should be stored in the channel for you).

For example:

```
$ ./fetch https://google.com https://github.com http://concordacademy.org
0.07s http://concordacademy.org
0.17s https://github.com
0.33s https://google.com
```

You need to loop through all the arguments (remember `os.Args`) except `os.Args[0]`
and fetch each website concurrently.

## 3. Bonus: dailywalk.go

(Problem taken from [here](http://whipperstacker.com/2015/10/05/3-trivial-concurrency-exercises-for-the-confused-newbie-gopher/))

Every morning, Alice and Bob go for a walk, and being creatures of habit, they follow the same routine every day.

First, they both prepare, grabbing sunglasses, perhaps a belt, closing open windows, turning off ceiling fans, and pocketing their phones and keys.

Once they’re both ready, which typically takes each of them between 10 and 20 seconds, they arm the alarm, which has a 20 second delay.

While the alarm is counting down, they both put on their shoes, a process which tends to take each of them between 5 and 15 seconds.

Then they leave the house together and lock the door, before the alarm has finished its countdown.

Write a program to simulate Alice and Bob’s morning routine.

Here’s some sample output from running a solution to this problem.

```
Let's go for a walk!
Bob has started getting ready.
Alice has started getting ready.
Alice spent 14 seconds getting ready
Bob spent 17 seconds getting ready
Bob started putting on shoes
Arming alarm
Alice started putting on shoes
Alice spent 9 seconds putting on shoes
Bob spent 11 seconds putting on shoes
Exiting and locking the door
Alarm is armed
```

Remember you can use the following function to generate random numbers:

```go
func Random(min, max int) int {
    return rand.Intn(max-min) + min
}
```

And you also need to seed the random number generator:

```go
rand.Seed(time.Now().UTC().UnixNano())
```
